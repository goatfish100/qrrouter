package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))
var session, err = mgo.Dial("localhost")

// if err != nil {
// 	panic(err)
// }
// defer session.Close()

type Resource struct {
	Id          string
	Description string
	Protected   string
	Action      string
	Address     string
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc(`/{[a-zA-Z0-9=\-\/\//]+	}`, HomeHandler)

	r.HandleFunc("/uuid/{key}", UUIDHandler)
	r.HandleFunc("/forward/{key}", ForwardHandler)
	r.HandleFunc("/test/", TestHandler)
	http.Handle("/", r)
	r.NotFoundHandler = http.HandlerFunc(HomeHandler)
	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/home/jamesl/asdfdsfd"))))
	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
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

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// This handler is to handle _ send resource on thier way
func UUIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside UUIDHandler")

	vars := mux.Vars(r)
	result := FetchResource(vars["key"])
	fmt.Println(result.Address)
	var saddress string = result.Address
	r.URL = testutils.ParseURI(result.Address)
	r.RequestURI = ""

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

func FetchResource(resourceid string) Resource {
	c := session.DB("test4").C("res")
	result := Resource{}
	err = c.Find(bson.M{"id": resourceid}).One(&result)

	if err != nil {
		log.Fatal(err)
	}
	return result
}

// This handler is to handle _ send resource on thier way
func ForwardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside forward handler")
	r.URL = testutils.ParseURI("https://www.google.com")
	r.RequestURI = ""
	http.Redirect(w, r, "https://www.google.com", http.StatusFound)

}
