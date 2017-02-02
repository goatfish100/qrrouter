package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server *httptest.Server
)

func init() {

}

func TestGorilla(t *testing.T) {

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
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

func TestUUIDRoute(t *testing.T) {

	req, err := http.NewRequest("GET", "/uuid/059edd7c-d454-11e6-92b9-374c2fc3d623", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UUIDHandler)

	handler.ServeHTTP(rr, req)
	//
	if rr.Code == 404 {
		t.Log("asdf")
		//Log("asdf")
	}
	t.Log("testUUIDRoute code ", rr.Code)
	t.Log("testUUIDRoute code ", rr.Code)
	//fmt.Println(rr.Result())
	t.Log(string(rr.Body.Bytes()))

}
