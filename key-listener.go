package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.Handle("/pickUpKeys", bodyCloser(http.HandlerFunc(determineRoute)))

	http.Handle("/dropOffKeys", bodyCloser(http.HandlerFunc(determineRoute)))

	log.Fatal(http.ListenAndServe(":33000", nil))

}

func bodyCloser(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		handler.ServeHTTP(w, r)
	})

}

func determineRoute(rw http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "POST":
		recordKeys(rw, req)
	case "GET":
		printKeys(rw, req)
	}

}

func recordKeys(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t Passwords
	err := decoder.Decode(&t)

	if err != nil {
		fmt.Println(err.Error())
	}

	rankingsJSON, _ := json.Marshal(t)
	err = ioutil.WriteFile("currentKeys.json", rankingsJSON, 0644)
}

func printKeys(w http.ResponseWriter, r *http.Request) {
	raw, err := ioutil.ReadFile("currentKeys.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	w.Write(raw)

}
