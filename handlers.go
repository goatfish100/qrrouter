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
	//r.HandleFunc(`/{[a-zA-Z0-9=\-\/\//]+	}`, HomeHandler)

	r.HandleFunc("/uuid/{key}", UUIDHandler)
	r.HandleFunc("/test", TestHandler)
	http.Handle("/", r)
	r.NotFoundHandler = http.HandlerFunc(HomeHandler)

	//r.PathPrefix("/jlsone/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/home/jamesl/gowork/src/bitbucket.org/gorillaweb/static"))))
	r.PathPrefix("/jlsone/").Handler(http.FileServer(http.Dir(http.Dir("/home/jamesl/gowork/src/bitbucket.org/gorillaweb/static"))))

	return r
}

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
	session, err := store.Get(r, os.Getenv("COOKIE_SECRET"))
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
	if result.Address != "" {

		log.Println("the address is " + result.Address)
		var saddress = result.Address
		r.URL = testutils.ParseURI(result.Address)
		r.RequestURI = ""

		//redirect the url
		if result.Action == "redirect" {
			log.Println("http redirect called")
			http.Redirect(w, r, result.Address, http.StatusFound)
			log.Println("redirecting")
		}

		log.Println("Proxying")

		// Proxy the result through service
		session, err := store.Get(r, os.Getenv("COOKIE_SECRET"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set some session values.
		session.Values["redirection_url"] = saddress
		session.Save(r, w)

		//forward.New() - is proxying connection
		log.Println(r.URL)
		fwd, _ := forward.New()
		fwd.ServeHTTP(w, r)
	} else {
		//TODO - Forward/send to real not found resource
		log.Println("UUIDHandler - no resource found for ", vars["key"])

		//http.Error("No QR Helper found!\n", "QR resource not found", http.StatusNotFound)
		w.Write([]byte("No QR Helper found!\n"))
	}
}
