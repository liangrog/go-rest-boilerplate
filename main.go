package main

import (
	"log"
	"net/http"
)

func main() {
	//load up mux router
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
