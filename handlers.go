package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
)

//NewRouter - new Gorilla Router
func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/uuid/{key}", UUIDHandler)
	r.HandleFunc("/test", TestHandler)
	http.Handle("/", r)
	r.NotFoundHandler = http.HandlerFunc(HomeHandler)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	return r
}

var varRedirect = http.Redirect

//TestHandler - a test handler - hello Gorilla
func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside test handler")
	w.Write([]byte("Gorilla!\n"))
}

//HomeHandler proxy request home handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// for relative urls - we need to get the resource
	// so we can proxy/server it -- aka
	// if a request comes in as /images/image1.jpg
	// we need to forward to http://resource.com/images/image1.jpg
	fmt.Println("HomeHandler")
	session, err := CookieStore.Get(r, os.Getenv("COOKIE_SECRET"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve our struct and type-assert it
	val := session.Values["redirection_url"].(string)
	fmt.Println("proxy session url", val)
	fmt.Println("proxy url", r.URL.String())

	//itemaddress = val + r.URL.String()
	fmt.Println(val + r.URL.String())
	r.URL = testutils.ParseURI(val + r.URL.String())

	fwd, _ := forward.New()
	fwd.ServeHTTP(w, r)

}

//PROXYHandler is a proxying/routine for URL's to be served from QRRouter
func PROXYHandler(w http.ResponseWriter, r *http.Request, address string) {
	// Proxy the result through service
	session, err := CookieStore.Get(r, os.Getenv("COOKIE_SECRET"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["redirection_url"] = address
	session.Save(r, w)

	//forward.New() - is proxying connection
	log.Println(r.URL)
	fwd, _ := forward.New()
	fwd.ServeHTTP(w, r)
}

var varPROXYHandler = PROXYHandler

//UUIDHandler This handler is to handle _ send resource on thier way
//either by proxying/forward the request or redirect
//Since - some sites may have additional resources - with just relative path - as in
// /images/icon1.jpg ... /js/library1.js
// we create a session - and handle these resources
// in homehandler
func UUIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("inside UUIDHandler")

	vars := mux.Vars(r)
	log.Println("uuid keys", vars["key"])
	result := FetchResource(vars["key"])
	log.Println("Address is ", result.Address)
	if result.Address != "" {

		log.Println("the address is " + result.Address)
		var saddress = result.Address

		//redirect the url
		log.Println("before IF block")
		if result.Action == "redirect" {
			r.URL = testutils.ParseURI(result.Address)
			r.RequestURI = ""
			log.Println("...http redirect called", result.Address)

			varRedirect(w, r, result.Address, http.StatusTemporaryRedirect)
			log.Println("redirecting")
		} else if result.Action == "proxy" {
			r.URL = testutils.ParseURI(result.Address)
			r.RequestURI = ""
			log.Println("...Proxying")
			varPROXYHandler(w, r, saddress)
		} else if result.Action == "s3serve" {
			log.Println("...AmazonS3Handler")
			varAmazonS3Handler(w, r, result.Address, result.Name)
			//AmazonS3Handler(w, r, result.Address, result.Name)
		} else if result.Action == "s3redirect" {
			log.Println("...AmazonS3URIHandler")
			varAmazonS3URIHandler(w, r, result.Address, result.Name)
		} else {
			log.Println("no catch found for ", result.Action)
		}

	} else {
		//TODO - Forward/send to real not found resource
		log.Println("UUIDHandler - no resource found for ", vars["key"])
		// Serve the helper resource not found page
		http.ServeFile(w, r, "./static/qrnotfound.html")
	}

}
