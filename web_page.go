package main

import (
	"log"
	"net/http"

	"os"

	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SECRET")))
var mongohost = os.Getenv("MONGO_HOST")
var AWSURL = os.Getenv("AWS_URL")
var AWSKEY = os.Getenv("AWS_KEY")
var AWSPASSPHRASE = os.Getenv("AWS_PASSPHRASE")
var AWSBUCKET = os.Getenv("AWS_BUCKET")

var mgosession, err = mgo.Dial(mongohost)

//MongoDBDatabase - mongo database name
var MongoDBDatabase = "resources"

func main() {

	router := NewRouter()
	var portnumber = os.Getenv("QRROUTER_PORT")
	log.Fatal(http.ListenAndServe(portnumber, router))
}
