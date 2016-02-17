package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func todos(w http.ResponseWriter, r *http.Request) {
	//mockup data
	todos := Todos{
		Todo{Name: "feed dogs"},
		Todo{Name: "do shopping"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}
