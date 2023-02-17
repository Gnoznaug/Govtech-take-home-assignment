package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gnoznaug/src/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterTeacherRoutes(router)
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe("localhost:8080",router))
}