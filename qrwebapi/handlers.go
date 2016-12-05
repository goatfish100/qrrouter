package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func GetResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var resourceId int
	var err error
	if resourceId, err = strconv.Atoi(vars["resourceId"]); err != nil {
		panic(err)
	}
	fmt.Println("resource id ", resourceId)
	res := RepoFindResource(resourceId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if res.Id > 0 {
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
		return
	}
	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func GetOrgs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println(orgs)
	if err := json.NewEncoder(w).Encode(orgs); err != nil {
		panic(err)
	}
	return

}

func ResourceCreate(w http.ResponseWriter, r *http.Request) {
	var resource Resource
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	fmt.Println(body)
	if err := json.Unmarshal(body, &resource); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Println(resource.Address)
}
