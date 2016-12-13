package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))
var session, err = mgo.Dial("localhost")

// if err != nil {
// 	panic(err)
// }
// defer session.Close()

type QRResource struct {
	Data Resource `json:"data"`
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc(`/{[a-zA-Z0-9=\-\/\//]+	}`, HomeHandler)

	r.HandleFunc("/uuid/{key}", UUIDHandler)
	r.HandleFunc("/getResource/{key}", GetResourceHandler)

	r.HandleFunc("/forward/{key}", ForwardHandler)
	r.HandleFunc("/test", TestHandler)
	http.Handle("/", r)
	//r.NotFoundHandler = http.HandlerFunc(HomeHandler)
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	//r.PathPrefix("/jlsone/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/home/jamesl/gowork/src/bitbucket.org/gorillaweb/static"))))
	r.PathPrefix("/jlsone").Handler(http.FileServer(http.Dir(http.Dir("/home/jamesl/gowork/src/bitbucket.org/gorillaweb/static"))))
	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
