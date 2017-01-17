package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()
	//r.HandleFunc(`/{[a-zA-Z0-9=\-\/\//]+	}`, HomeHandler)

	r.HandleFunc("/uuid/{key}", UUIDHandler)
	r.HandleFunc("/getResource/{key}", GetResourceHandler)

	r.HandleFunc("/test", TestHandler)
	http.Handle("/", r)
	r.NotFoundHandler = http.HandlerFunc(HomeHandler)

	r.PathPrefix("/jlsone/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/home/jamesl/gowork/src/bitbucket.org/gorillaweb/static"))))
	r.PathPrefix("/jlsone").Handler(http.FileServer(http.Dir(http.Dir("/home/jamesl/gowork/src/bitbucket.org/gorillaweb/static"))))

	return r
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside test handler")
	w.Write([]byte("Gorilla!\n"))
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// for relative urls - we need to get the resource
	// so we can proxy/server it -- aka
	// if a request comes in as /images/image1.jpg
	// we need to forward to http://resource.com/images/image1.jpg
	fmt.Println("HomeHandler")
	session, err := store.Get(r, "gorillasession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve our struct and type-assert it
	val := session.Values["redirection_url"].(string)
	fmt.Println("redirection url", val)
	fmt.Println("requested url", r.URL.String())
	//describe(val)

	//itemaddress = val + r.URL.String()
	fmt.Println(val + r.URL.String())
	r.URL = testutils.ParseURI(val + r.URL.String())
	// // r.RequestURI = ""
	// //
	fwd, _ := forward.New()
	fwd.ServeHTTP(w, r)

}

func GetResourceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["key"])
	w.Header().Set("Content-Type", "application/vnd.api+json")
	result := FetchResource(vars["key"])
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// This handler is to handle _ send resource on thier way
func UUIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside UUIDHandler")

	vars := mux.Vars(r)
	fmt.Println(vars["key"])
	result := FetchResource(vars["key"])
	fmt.Println("the address is " + result.Address)
	var saddress string = result.Address
	r.URL = testutils.ParseURI(result.Address)
	r.RequestURI = ""

	if result.Action == "forward" {
		http.Redirect(w, r, result.Address, http.StatusFound)

	}

	session, err := store.Get(r, "gorillasession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["redirection_url"] = saddress
	session.Save(r, w)

	fwd, _ := forward.New()
	fwd.ServeHTTP(w, r)
}
