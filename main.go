package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handleLoginRequest)
	http.HandleFunc("/logout", handleLogoutRequest)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}

func handleLoginRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from login endpoint. you visited %s\n", r.URL.Path)
}

func handleLogoutRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from logout endpoint. you visited %s\n", r.URL.Path)
}