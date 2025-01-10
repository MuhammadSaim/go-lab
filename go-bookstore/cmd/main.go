package main

import (
	"log"
	"net/http"

	"github.com/MuhammadSaim/go-lab/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9019", r))
}
