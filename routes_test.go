package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
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

func TestUUIDRoute1(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/uuid/{key}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))

	ts := httptest.NewServer(r)
	//rr := httptest.NewRecorder()

	defer ts.Close()

	// Table driven test
	var suid = "444edd7c-d454-11e6-92b9-374c2fc3d623"

	url := ts.URL + "/uuid/" + suid

	resp, errget := http.Get(url)
	if errget != nil {
		t.Fatal(errget)
	}

	log.Println("status is ", resp.StatusCode)

}

func TestUUIDRoute2(t *testing.T) {
	//r := mux.NewRouter()

	req, err2 := http.NewRequest("GET", "/uuid/444edd7c-d454-11e6-92b9-374c2fc3d623", nil)

	context.Set(req, "key", "444edd7c-d454-11e6-92b9-374c2fc3d623")

	if err2 != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UUIDHandler)
	//
	handler.ServeHTTP(rr, req)
	//
	// if rr.Code == 404 {
	// 	t.Log("TestUUIDRoute2 Fail")
	// 	t.Fail()
	// }
	// t.Log("testUUIDRoute code ", rr.Code)
	// t.Log("testUUIDRoute code ", rr.Code)
	// //fmt.Println(rr.Result())
	// t.Log("TestUUIDRoute2", string(rr.Body.Bytes()))

}

//
// func TestUUIDRoute5(t *testing.T) {
// 	//r := mux.NewRouter()
//
// 	req, err2 := http.NewRequest("GET", "/uuid/444edd7c-d454-11e6-92b9-374c2fc3d626", nil)
//
// 	context.Set(req, "key", "444edd7c-d454-11e6-92b9-374c2fc3d626")
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
// func TestUUIDRoute6(t *testing.T) {
// 	//r := mux.NewRouter()
//
// 	req, err2 := http.NewRequest("GET", "/uuid/444edd7c-d454-11e6-92b9-374c2fc3d627", nil)
//
// 	context.Set(req, "key", "444edd7c-d454-11e6-92b9-374c2fc3d627")
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
// func TestUUIDRoute6FailTest(t *testing.T) {
// 	//r := mux.NewRouter()
//
// 	//choose one item that does not exist
// 	req, err2 := http.NewRequest("GET", "/uuid/44433333-d454-11e6-92b9-374c2fc3DDDD", nil)
//
// 	context.Set(req, "key", "444edd7c-d454-11e6-92b9-374c2fc3d627")
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
