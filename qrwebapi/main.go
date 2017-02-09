package main

import (
	"log"
	"net/http"
)

//Mongo database name
var MongoDBDatabase = "resources"

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8001", router))
}
