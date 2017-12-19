package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/pickUpKeys", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		//printKeys(w, r)
		printKeys(w, r)
	})

	http.HandleFunc("/dropOffKeys", func(w http.ResponseWriter, r *http.Request) {
		//recordKeys(r, "currentKeys.json")
		recordKeys(w, r)

	})

	log.Fatal(http.ListenAndServe(":33000", nil))

}

func recordKeys(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t Passwords
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	rankingsJSON, _ := json.Marshal(t)
	err = ioutil.WriteFile("currentKeys.json", rankingsJSON, 0644)
}

func printKeysOld(w http.ResponseWriter, r *http.Request) {

	todaysPasswords, err := ioutil.ReadFile("currentKeys.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Fprintf(w, "%q", todaysPasswords)

}

func printKeys(w http.ResponseWriter, r *http.Request) {
	raw, err := ioutil.ReadFile("currentKeys.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//var currentPassword Passwords
	//json.Unmarshal(raw, &currentPassword)
	w.Write(raw)

}
