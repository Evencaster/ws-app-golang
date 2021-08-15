package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main()  {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./frontend"))

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	log.Fatal(http.ListenAndServe(":8111", r))
}
