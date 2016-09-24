package main

import (
	"fmt"
	"net/http"
	s "strings"

	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
)

func main() {
	// Forwards incoming requests to whatever location URL points to, adds proper forwarding headers
	fwd, _ := forward.New()

	redirect := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		if s.Contains(req.RequestURI, "redirrand") {
			// let us forward this request to another server
			var rediruri string = retURL(s.TrimPrefix(req.RequestURI, "/"))
			fmt.Println(rediruri)
			req.URL = testutils.ParseURI(rediruri)
			req.RequestURI = ""
			//fmt.Println(req.Referer(), req.URL, req.Referer())
			fwd.ServeHTTP(w, req)
		} else {
			fwd.ServeHTTP(w, req)
		}

	})

	// that's it! our reverse proxy is ready!
	s := &http.Server{
		Addr:    ":8080",
		Handler: redirect,
	}
	s.ListenAndServe()
}
func retURL(lookup string) string {
	elements := make(map[string]string)
	elements["redirrand/f793511c83c3"] = "https://www.google.com"
	elements["redirrand/071392a13c1b"] = "https://www.linkedin.com"
	elements["redirrand/071392a13c1d"] = "https://www.yahoo.com"

	return elements[lookup]
}
