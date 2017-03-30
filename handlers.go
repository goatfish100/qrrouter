package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"

	"github.com/minio/minio-go"
)

//NewRouter - new Gorilla Router
func NewRouter() *mux.Router {

	r := mux.NewRouter()
	//r.HandleFunc(`/{[a-zA-Z0-9=\-\/\//]+	}`, HomeHandler)

	r.HandleFunc("/uuid/{key}", UUIDHandler)
	r.HandleFunc("/test", TestHandler)
	r.HandleFunc("/amazon", AmazonS3Handler)
	http.Handle("/", r)
	r.NotFoundHandler = http.HandlerFunc(HomeHandler)

	//r.PathPrefix("/jlsone/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/home/jamesl/gowork/src/bitbucket.org/gorillaweb/static"))))
	r.PathPrefix("/jlsone/").Handler(http.FileServer(http.Dir(http.Dir("./static"))))

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

//AmazonS3Handler proxy request home handler
func AmazonS3Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AmazonS3Handler")

	//s3Client, err := minio.New(AWSURL, AWSKEY, AWSPASSPHRASE, true)
	s3Client, err := minio.New("s3.amazonaws.com", "AKIAJ7K7I7KUWLIR6CEA", "PBJn37kTAHt5Jbk0ELR6NqnQkHuxlmrCx/Rehf4h", true)

	if err != nil {
		log.Fatalln(err)
	}

	reader, err := s3Client.GetObject("goatfish100", "test_folder/vouncher.pdf")

	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	localFile, err := os.Create("my-testfile")
	if err != nil {
		log.Fatalln(err)
	}
	defer localFile.Close()

	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Disposition", localFile.Name())
	//w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Type", "pdf")

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	fwd, _ := forward.New()
	fwd.ServeHTTP(w, r)

}

func PROXYHandler(w http.ResponseWriter, r *http.Request, address string) {
	// Proxy the result through service
	session, err := store.Get(r, os.Getenv("COOKIE_SECRET"))
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
		} else {
			PROXYHandler(w, r, saddress)
		}

		log.Println("Proxying")

	} else {
		//TODO - Forward/send to real not found resource
		log.Println("UUIDHandler - no resource found for ", vars["key"])

		//http.Error("No QR Helper found!\n", "QR resource not found", http.StatusNotFound)
		w.Write([]byte("No QR Helper found!\n"))
	}
}
