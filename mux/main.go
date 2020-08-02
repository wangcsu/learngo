package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	InitializeRouter(r)
	//http.Handle("/", r)
	http.ListenAndServe(":8888", r)
	//http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir("."))))
	//http.ListenAndServe(":8888", nil)
}
