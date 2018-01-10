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

func s3connect(AwsBucket string, resource string) (*minio.Object, error) {
	s3Client, err := minio.New(AwsURL, AwsKey, AwsPassPhrase, true)

	//if err != nil {
	//	log.Fatal(err)
	//}

	reader, err := s3Client.GetObject(AwsBucket, resource)

	defer reader.Close()
	return reader, err

}

var vars3connect = s3connect
//AmazonS3Handler proxy request home handler
func AmazonS3Handler(w http.ResponseWriter, r *http.Request, resource string, filename string) {
	fmt.Println("----AmazonS3Handler")

	//s3Client, err := minio.New(AwsURL, AwsKey, AwsPassPhrase, true)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//reader, err := s3Client.GetObject(AwsBucket, resource)
	//
	//if err != nil {
	//	log.Println("AmazonS3Handler Error openning S3 connection")
	//	panic(err)
	//}
	//defer reader.Close()
	//
	//if err != nil {
	//	log.Println("AmazonS3Handler Error closing S3 connection")
	//
	//	panic(err)
	//}
	//
	reader, err := vars3connect(AwsBucket, resource)

	w.Header().Set("Content-Disposition: inline; filename=", filename)
	w.Header().Set("Content-Type", "pdf")

	b, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Println("AmazonS3Handler readall err")
		panic(err)
	}
	w.Write(b)

	fwd, _ := forward.New()

	fwd.ServeHTTP(w, r)
}



//AmazonS3URIHandler getnerate downlink link
func AmazonS3URIHandler(w http.ResponseWriter, r *http.Request, resource string, filename string) {
	s3Client, err := minio.New(AwsURL, AwsKey, AwsPassPhrase, true)
	if err != nil {
		log.Println("AmazonS3URIHandler readall err")
		panic(err)
	}

	// Set request parameters
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+filename+"\"")

	// Gernerate presigned get object url.
	presignedURL, err := s3Client.PresignedGetObject(AwsBucket, resource, time.Duration(1000)*time.Second, reqParams)
	if err != nil {
		log.Println("AmazonS3URIHandler Error getting pre signed url")
		panic(err)
	}
	log.Println("pre signed url", presignedURL)
	http.Redirect(w, r, presignedURL.String(), http.StatusFound)

}
var varAmazonS3URIHandler = AmazonS3URIHandler
