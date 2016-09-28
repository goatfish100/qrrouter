package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	s "strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/uuid/{key}", UUIDHandler)
	r.HandleFunc("/forward/{key}", ForwardHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/home/jamesl/temp2"))))
	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// This handler is to handle _ send resource on thier way
func UUIDHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// Forwards incoming requests to whatever location URL points to, adds proper forwarding headers
	var rediruri string = retURL(s.TrimPrefix(r.RequestURI, "/"))
	fmt.Println(rediruri)
	r.URL = testutils.ParseURI(rediruri)
	r.RequestURI = ""
	fwd, _ := forward.New()
	fwd.ServeHTTP(w, r)

}

// This handler is to handle _ send resource on thier way
func ForwardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside forward handler")
	r.URL = testutils.ParseURI("https://www.google.com")
	r.RequestURI = ""
	http.Redirect(w, r, "https://www.google.com", http.StatusFound)

}
func retURL(lookup string) string {
	elements := make(map[string]string)
	//missing resources
	elements["uuid/f793511c83c3"] = "https://www.google.com"
	//linked in works ...
	elements["uuid/071392a13c1b"] = "https://www.linkedin.com"
	//missing resources
	elements["uuid/071392a13c1d"] = "https://www.yahoo.com"

	return elements[lookup]
}
