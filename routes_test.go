package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/vulcand/oxy/forward"
)

var (
	server *httptest.Server
)

func init() {

}

func TestGorilla(t *testing.T) {

	req, errg := http.NewRequest("GET", "/getorgs", nil)
	if errg != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TestHandler)

	handler.ServeHTTP(rr, req)
	//
	if rr.Code == 404 {
		t.Log("404 code found")
	}
	t.Log("rr code ", rr.Code)
	//fmt.Println(rr.Result())
	t.Log(string(rr.Body.Bytes()))

}
func InvokeHandler(handler http.Handler, routePath string,
	w http.ResponseWriter, r *http.Request) {

	// Add a new sub-path for each invocation since
	// we cannot (easily) remove old handler
	invokeCount++
	router := mux.NewRouter()
	http.Handle(fmt.Sprintf("/%d", invokeCount), router)

	router.Path(routePath).Handler(handler)

	// Modify the request to add "/%d" to the request-URL
	r.URL.RawPath = fmt.Sprintf("/%d%s", invokeCount, r.URL.RawPath)
	router.ServeHTTP(w, r)
}
func TestGetUUID1(t *testing.T) {
	t.Parallel()
	t.Log("TestGetUUID Test")

	varRedirect = func(w http.ResponseWriter, r *http.Request, resource string, status int) {
		t.Log("TestGetUUID1 redirect handler")
		w.Write([]byte("Gorilla!\n"))
	}

	path := "/uuid/444edd7c-d454-11e6-92b9-374c2fc3d623"
	t.Log("TestGetResource the path is ", path)
	r, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()

	InvokeHandler(http.HandlerFunc(UUIDHandler), "/uuid/{key}", w, r)
	t.Log("TestGetUUID1 code", w.Code)
	assert.Equal(t, http.StatusOK, w.Code)

	t.Log("TestGetUUID1 the return string is bbxx", string(w.Body.Bytes()))

}

func TestGetUUID2(t *testing.T) {
	t.Parallel()
	t.Log("TestGetResource Test")

	varPROXYHandler = func(w http.ResponseWriter, r *http.Request, resource string) {
		t.Log("TestGetUUID2 varProxyHandler")
		w.Write([]byte("Gorilla!\n"))
	}

	path := "/uuid/444edd7c-d454-11e6-92b9-374c2fc3d624"
	t.Log("TestGetResource the path is ", path)
	r, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()

	InvokeHandler(http.HandlerFunc(UUIDHandler), "/uuid/{key}", w, r)
	t.Log("TestGetUUID2 code", w.Code)
	assert.Equal(t, http.StatusOK, w.Code)
	//assert.Equal(t, http.StatusFound, w.Code)

	t.Log("the return string is bbxx", string(w.Body.Bytes()))

}

//TestProxyHandler - add more on this test later
func TestPROXYHandler(t *testing.T) {
	t.Parallel()
	t.Log("TestPROXYHandler Test")

	path := "http://www.google.com"
	r, _ := http.NewRequest("GET", path, nil)

	w := httptest.NewRecorder()
	PROXYHandler(w, r, path)
	session, err := CookieStore.Get(r, os.Getenv("COOKIE_SECRET"))

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, session.Values["redirection_url"], path)

}

//TestGetUUID3 - TODO - set up minio library
//to mock
func TestGetUUID3(t *testing.T) {
	t.Parallel()
	t.Log("TestGetResource Test")

	path := "/uuid/444edd7c-d454-11e6-92b9-374c2fc3d626"
	t.Log("TestGetResource the path is ", path)
	r, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()

	//varAmazonS3Handler
	varAmazonS3Handler = func(w http.ResponseWriter, r *http.Request, resource string, filename string) {
		fmt.Println("----varAmazonS3Handler local function")

		//b, err := vars3getObjectBytes(resource)
		b := []byte("AAAA")
		w.Header().Set("Content-Disposition: inline; filename=", filename)
		w.Header().Set("Content-Type", "pdf")

		//b, err := ioutil.ReadAll(reader)

		w.Write(b)

		fwd, _ := forward.New()

		fwd.ServeHTTP(w, r)
	}
	InvokeHandler(http.HandlerFunc(UUIDHandler), "/uuid/{key}", w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	t.Log("the return string is bbxx", string(w.Body.Bytes()))

}

//TestGetUUID4 - TODO - set up minio library
//to mock
func TestGetUUID4(t *testing.T) {
	t.Parallel()
	t.Log("TestGetUUID4 S3 redirect Test")

	path := "/uuid/444edd7c-d454-11e6-92b9-374c2fc3d627"
	t.Log("TestGetResource the path is ", path)
	r, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()

	varAmazonS3Handler = func(w http.ResponseWriter, r *http.Request, resource string, filename string) {
		// //s3Client, err := minio.New(AwsURL, AwsKey, AwsPassPhrase, true)
		// s3Client, err := vars3connect()
		if err != nil {
			log.Println("AmazonS3URIHandler readall err")
			panic(err)
		}

		// Set request parameters
		reqParams := make(url.Values)
		reqParams.Set("response-content-disposition", "attachment; filename=\""+filename+"\"")

		//presignedURL, err := s3Client.PresignedGetObject(AwsBucket, resource, time.Duration(1000)*time.Second, reqParams)
		if err != nil {
			log.Println("AmazonS3URIHandler Error getting pre signed url")
			panic(err)
		}
		w.Write([]byte("Gorilla!\n"))

	}

	InvokeHandler(http.HandlerFunc(UUIDHandler), "/uuid/{key}", w, r)
	//	assert.Equal(t, http.StatusFound, w.Code)

	t.Log("the return string is bbxx", string(w.Body.Bytes()))

}
func TestUUIDRoute5(t *testing.T) {

	varAmazonS3URIHandler = func(w http.ResponseWriter, r *http.Request, resource string, filename string) {
		t.Log("TestGetUUID5 inside test handler")
		w.Write([]byte("Gorilla!\n" + resource + filename))
	}

	path := "/uuid/444edd7c-d454-11e6-92b9-374c2fc3d627"
	t.Log("TestGetResource the path is ", path)
	r, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()

	InvokeHandler(http.HandlerFunc(UUIDHandler), "/uuid/{key}", w, r)
	//	assert.Equal(t, http.StatusFound, w.Code)

	t.Log("the return string is bbxx", string(w.Body.Bytes()))

}

// func TestUUIDRoute2(t *testing.T) {
// 	//r := mux.NewRouter()
//
// 	req, err2 := http.NewRequest("GET", "/uuid/444edd7c-d454-11e6-92b9-374c2fc3d623", nil)
// 	context.Set(req, "key", "444edd7c-d454-11e6-92b9-374c2fc3d623")
// 	context.Set(req, "uuid", "444edd7c-d454-11e6-92b9-374c2fc3d623")
//
// 	if err2 != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(UUIDHandler)
//
// 	handler.ServeHTTP(rr, req)
// 	//
// 	if rr.Code == 404 {
// 		t.Log("TestUUIDRoute2 Fail")
// 		t.Fail()
// 	}
// 	t.Log("testUUIDRoute code ", rr.Code)
// 	t.Log("testUUIDRoute code ", rr.Code)
// 	//fmt.Println(rr.Result())
// 	t.Log("TestUUIDRoute2", string(rr.Body.Bytes()))
//
// }

//
// func TestUUIDRoute3(t *testing.T) {
// 	//r := mux.NewRouter()
//
// 	req, err2 := http.NewRequest("GET", "/uuid/444edd7c-d454-11e6-92b9-374c2fc3d624", nil)
//
// 	context.Set(req, "uuid", "444edd7c-d454-11e6-92b9-374c2fc3d624")
//
// 	if err2 != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(UUIDHandler)
//
// 	handler.ServeHTTP(rr, req)
// 	//
// 	if rr.Code == 404 {
// 		t.Log("TestUUIDRoute2 Fail")
// 		t.Fail()
// 	}
//
// 	t.Log("The rr header is xox", rr.Header())
// 	t.Log("testUUIDRoute code ", rr.Code)
// 	//fmt.Println(rr.Result())
// 	t.Log("TestUUIDRoute3", string(rr.Body.Bytes()))
//
// }
//

//
// func TestUUIDRoute6(t *testing.T) {
// 	//r := mux.NewRouter()
//
// 	req, err2 := http.NewRequest("GET", "/uuid/444edd7c-d454-11e6-92b9-374c2fc3d627", nil)
//
// 	context.Set(req, "uuid", "444edd7c-d454-11e6-92b9-374c2fc3d627")
//
// 	if err2 != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(UUIDHandler)
//
// 	handler.ServeHTTP(rr, req)
// 	//
// 	if rr.Code == 404 {
// 		t.Log("TestUUIDRoute5 Fail")
// 		t.Fail()
// 	}
// 	t.Log("testUUIDRoute code ", rr.Code)
// 	t.Log("testUUIDRoute code ", rr.Code)
// 	//fmt.Println(rr.Result())
// 	t.Log("TestUUIDRoute5", string(rr.Body.Bytes()))
//
// }
//
func TestUUIDRoute6FailTest(t *testing.T) {
	//r := mux.NewRouter()

	//choose one item that does not exist
	req, err2 := http.NewRequest("GET", "/uuid/doesnotexist", nil)

	if err2 != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UUIDHandler)

	handler.ServeHTTP(rr, req)
	//
	if rr.Code == 404 {
		t.Log("TestUUIDRoute6FailTest Fail")
		t.Fail()
	}
	t.Log("TestUUIDRoute6FailTest code ", rr.Code)
	t.Log("TestUUIDRoute6FailTest code ", rr.Code)
	//fmt.Println(rr.Result())
	//result := string(rr.Body.Bytes())

	t.Log("TestUUIDRoute6FailTest", string(rr.Body.Bytes()))

}
