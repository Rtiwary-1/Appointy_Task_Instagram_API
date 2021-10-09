package main

import (
	"github.com/gorilla/mux"
	"net/http"
)




func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/posts", CreatePostEndpoint).Methods("POST")
	router.HandleFunc("/users/{id}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/posts/{id}", GetPostEndpoint).Methods("GET")
	router.HandleFunc("/posts/users/{id}", GetAllPostEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}