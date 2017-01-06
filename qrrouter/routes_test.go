package main

import (
	"fmt"
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
		fmt.Println("asdf")
	}
	fmt.Println("rr code ", rr.Code)
	//fmt.Println(rr.Result())
	fmt.Println(string(rr.Body.Bytes()))

}
