package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/sessions"

	mgo "gopkg.in/mgo.v2"
)

//MongoHost - the mongo host name
var MongoHost = os.Getenv("MONGO_HOST")

//CookieStore - the cookie store/encrypt phrase
var CookieStore = sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SECRET")))

//AwsURL - Amazon AWS url
var AwsURL = os.Getenv("AWS_URL")

//AwsKey - Amazon AWS key
var AwsKey = os.Getenv("AWS_KEY")

//AwsPassPhrase - Amazon AWS passphrase
var AwsPassPhrase = os.Getenv("AWS_PASSPHRASE")

//AwsBucket - Amazon AWS Bucket
var AwsBucket = os.Getenv("AWS_BUCKET")

//MgoSession - Amazon AWS key
var MgoSession, err = mgo.Dial(MongoHost)

//MongoDBDatabase - mongo database name
var MongoDBDatabase = "resources"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {

	router := NewRouter()
	var portnumber = ":" + os.Getenv("QRROUTER_PORT")
	log.Println("Port number is ", portnumber)
	// Set up CORS
	corsObj := handlers.AllowedOrigins([]string{"*"})
	// and use combined logging handler as well
	log.Fatal(http.ListenAndServe(portnumber, handlers.CombinedLoggingHandler(os.Stdout, handlers.CORS(corsObj)(router))))
}
