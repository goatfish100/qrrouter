package main

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
	"os"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))
var mongohost = os.Getenv("MONGO_HOST")
var mgosession, err = mgo.Dial(mongohost)

//MongoDBDatabase - mongo database name
var MongoDBDatabase = "resources"

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
