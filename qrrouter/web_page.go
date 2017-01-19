package main

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))
var session, err = mgo.Dial("localhost")
var MongoDBDatabase = "resources"

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
