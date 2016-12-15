package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader //Ignore this for now
	usersUrl string
)

func init() {
	server = httptest.NewServer(Handlers()) //Creating new server with the user handlers

	//usersUrl = fmt.Sprintf("%s/users", server.URL) //Grab the address for the API endpoint
}
func TestGorilla(t *testing.T) {
	req, err := http.NewRequest("GET", "/test/test", nil)
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
	fmt.Println(rr.Result())
	fmt.Println(rr.Body.Bytes())

}
