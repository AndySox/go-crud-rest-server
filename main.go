package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

var users []User

func createUser(w http.ResponseWriter, r *http.Request) {

}

func getUser(w http.ResponseWriter, r *http.Request) {

}

func listUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/create", createUser).Methods("POST")
	router.HandleFunc("/get/{:id}", getUser).Methods("GET")
	router.HandleFunc("/list", listUsers).Methods("GET")
	router.HandleFunc("/update/{:id}", updateUser).Methods("PUT")
	router.HandleFunc("/delete/{:id}", createUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":12345", router))
}
