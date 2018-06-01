package main

//go:generate mockgen -source=src/github.com/minio/minio-go/api.go -package=main -destination=cache_mock.go
import (
	"fmt"
	"log"
	"net/http"
	"time"

	"io/ioutil"
	"net/url"

	"github.com/minio/minio-go"
	"github.com/vulcand/oxy/forward"
)

func s3connect() (*minio.Client, error) {

	s3Client, error := minio.New(AwsURL, AwsKey, AwsPassPhrase, true)

	if err != nil {
		log.Println("Error s3Connect", error)
	}

	return s3Client, error

}
func s3getObjectBytes(resource string) ([]byte, error) {

	s3Client, err := vars3connect()
	if err != nil {
		log.Println("Error s3getObject ", err)
	}

	reader, err := s3Client.GetObject(AwsBucket, resource)
	defer reader.Close()
	if err != nil {
		log.Println("Error s3getObject ", err)
	}
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println("Error s3getObject ", err)
	}

	return b, err
}

func s3PreSignedURL(resource string) (*url.URL, error) {

	s3Client, err := vars3connect()
	if err != nil {
		log.Println("Error s3getObject ", err)
	}

	// Set request parameters
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+resource+"\"")

	presignedURL, err := s3Client.PresignedGetObject(AwsBucket, resource, time.Duration(1000)*time.Second, reqParams)
	if err != nil {
		log.Println("Error s3getObject ", err)
	}

	return presignedURL, err
}

var vars3connect = s3connect
var vars3getObjectBytes = s3getObjectBytes
var vars3PreSignedURL = s3PreSignedURL

//AmazonS3Handler proxy request home handler
func AmazonS3Handler(w http.ResponseWriter, r *http.Request, resource string, filename string) {
	fmt.Println("----AmazonS3Handler")

	b, err := vars3getObjectBytes(resource)

	w.Header().Set("Content-Disposition: inline; filename=", filename)
	w.Header().Set("Content-Type", "pdf")

	//b, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Println("AmazonS3Handler readall err")
		panic(err)
	}
	w.Write(b)

	fwd, _ := forward.New()

	fwd.ServeHTTP(w, r)
}

var varAmazonS3Handler = AmazonS3Handler

//AmazonS3URIHandler getnerate downlink link
func AmazonS3URIHandler(w http.ResponseWriter, r *http.Request, resource string, filename string) {

	// //s3Client, err := minio.New(AwsURL, AwsKey, AwsPassPhrase, true)
	// s3Client, err := vars3connect()
	if err != nil {
		log.Println("AmazonS3URIHandler readall err")
		panic(err)
	}

	// Set request parameters
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+filename+"\"")

	// Gernerate presigned get object url.
	signedURL, err := vars3PreSignedURL(resource)
	//presignedURL, err := s3Client.PresignedGetObject(AwsBucket, resource, time.Duration(1000)*time.Second, reqParams)
	if err != nil {
		log.Println("AmazonS3URIHandler Error getting pre signed url")
		panic(err)
	}
	log.Println("pre signed url", signedURL.String())
	http.Redirect(w, r, signedURL.String(), http.StatusFound)

}

var varAmazonS3URIHandler = AmazonS3URIHandler
