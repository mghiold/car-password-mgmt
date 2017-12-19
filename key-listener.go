package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//FilePath - The file path to read and write too
const FilePath = "currentKeys.json"

//ListenPort - The port the the server will listen on
const ListenPort = ":33000"

func main() {

	http.Handle("/pickUpKeys", bodyCloser(http.HandlerFunc(determineRoute)))

	http.Handle("/dropOffKeys", bodyCloser(http.HandlerFunc(determineRoute)))

	log.Fatal(http.ListenAndServe(ListenPort, nil))

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
	err = ioutil.WriteFile(FilePath, rankingsJSON, 0644)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		rw.WriteHeader(http.StatusAccepted)
	}

}

func printKeys(w http.ResponseWriter, r *http.Request) {
	raw, err := ioutil.ReadFile(FilePath)

	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(raw)

}
