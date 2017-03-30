package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

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
func AmazonS3Handler(w http.ResponseWriter, r *http.Request, resource string) {
	fmt.Println("----AmazonS3Handler")

	s3Client, err := minio.New(AWSURL, AWSKEY, AWSPASSPHRASE, true)

	if err != nil {
		log.Fatal(err)
	}

	reader, err := s3Client.GetObject(AWSBUCKET, resource)

	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Disposition", resource)
	w.Header().Set("Content-Type", "pdf")

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)

	fwd, _ := forward.New()
	fwd.ServeHTTP(w, r)
}

//AmazonS3Handler getnerate downlink link
func AmazonS3URIHandler(w http.ResponseWriter, r *http.Request, resource string) {
	log.Println("...AmazonS3URIHandler Handler ")
	// Set request parameters for content-disposition.
	minioClient, err := minio.New(AWSURL, AWSKEY, AWSPASSPHRASE, true)

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\"test.txt\"")

	// Generates a presigned url which expires in a day.
	presignedURL, err := minioClient.PresignedGetObject("mybucket", "myobject", time.Duration(1000)*time.Second, reqParams)

	log.Println("...http redirect called")
	http.Redirect(w, r, presignedURL.String(), http.StatusFound)
	log.Println("redirecting")

	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Pre signed url is ", presignedURL)
	// 	fmt.Println("----AmazonS3Handler")
	//
	// 	s3Client, err := minio.New(AWSURL, AWSKEY, AWSPASSPHRASE, true)
	//
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	reader, err := s3Client.GetObject(AWSBUCKET, resource)
	//
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	defer reader.Close()
	//
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	//
	// 	w.Header().Set("Content-Disposition", resource)
	// 	w.Header().Set("Content-Type", "pdf")
	//
	// 	b, err := ioutil.ReadAll(reader)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	w.Write(b)
	//
	// 	fwd, _ := forward.New()
	// 	fwd.ServeHTTP(w, r)
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

		//redirect the url
		log.Println("before IF block")
		if result.Action == "redirect" {
			r.URL = testutils.ParseURI(result.Address)
			r.RequestURI = ""
			log.Println("...http redirect called")
			http.Redirect(w, r, result.Address, http.StatusFound)
			log.Println("redirecting")
		} else if result.Action == "proxy" {
			r.URL = testutils.ParseURI(result.Address)
			r.RequestURI = ""
			log.Println("...Proxying")
			PROXYHandler(w, r, saddress)
		} else if result.Action == "s3serve" {
			log.Println("...AmazonS3Handler")
			AmazonS3Handler(w, r, result.Address)
		} else if result.Action == "s3redirect" {
			log.Println("...AmazonS3Handler")
			AmazonS3URIHandler(w, r, result.Address)
		}
		log.Println("no catch found")

	} else {
		//TODO - Forward/send to real not found resource
		log.Println("UUIDHandler - no resource found for ", vars["key"])

		//http.Error("No QR Helper found!\n", "QR resource not found", http.StatusNotFound)
		w.Write([]byte("No QR Helper found!\n"))
	}
}
