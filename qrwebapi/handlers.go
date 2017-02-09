package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gorilla/sessions"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"bitbucket.org/gorouter/datastructs"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))
var session, err = mgo.Dial("localhost")

//Index - sample index/test page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

//GetResource - return information about a resource
func GetResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var resourceID = vars["resourceId"]

	fmt.Println("resource id ", resourceID)
	res := RepoFindResource(resourceID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println(res)
	if res.ID == "" {
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

//GetOrg return all orgs
func GetOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//log.Printf("%s\t%s\t%s\t%s")
	c := session.DB("resources").C("orgusers")
	org := datastructs.Org{}
	// number, _ := strconv.Atoi(resourceid)

	//organization := OrgFind(orgId)
	//err = c.Find(bson.M{"_id": orgId}).One(&org)
	idQueryier := bson.ObjectIdHex(vars["orgId"])
	log.Printf("Org identifier%s\t ", idQueryier)

	err = c.Find(bson.M{"_id": idQueryier}).One(&org)

	//bson.M{"_id": bson.ObjectIdHex(orgId)}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println("inside GetOrg ", org)
	// check if empty array ...
	if org.ID == "" {
		error := jsonErr{Code: 404, Text: "Not Found"}
		if err := json.NewEncoder(w).Encode(error); err != nil {
			panic(err)
		}
	} else if err := json.NewEncoder(w).Encode(org); err != nil {
		panic(err)
	}

}

//GetOrgs - get organizations refactor into one func ...
func GetOrgs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println("inside GetOrgs", orgs)
	if err := json.NewEncoder(w).Encode(orgs); err != nil {
		panic(err)
	}
	return
}

//PostCreateOrg - a post operation to create organization
func PostCreateOrg(w http.ResponseWriter, r *http.Request) {
	var organization datastructs.Org
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &organization); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	OrgCreate(organization)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var jsonsuccess = datastructs.JSONSuccess{Success: "true"}

	if err := json.NewEncoder(w).Encode(jsonsuccess); err != nil {
		panic(err)
	}
	fmt.Println("org is ", organization.ID)

}

//PostCreateUser - A Post Operation to create a user
func PostCreateUser(w http.ResponseWriter, r *http.Request) {
	var user datastructs.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	//OrgCreate(organization)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var jsonsuccess = datastructs.JSONSuccess{Success: "true"}

	if err := json.NewEncoder(w).Encode(jsonsuccess); err != nil {
		panic(err)
	}
	//fmt.Println("org is ", organization.Id)

}

//TestHandler - a test handler - hello Gorilla
func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside test handler")
	w.Write([]byte("Gorilla!\n"))
}

//ResourceCreate - create a resource
func ResourceCreate(w http.ResponseWriter, r *http.Request) {
	var resource datastructs.Resource
	fmt.Println("inside Resource Create")
	log.Println("inside Resource Create")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	log.Println(body)
	if err := json.Unmarshal(body, &resource); err != nil {
		log.Println("Resource create - resouce decoded")
		InsertResource(resource)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	//InsertResource(resource)

	fmt.Println(resource.Address)
}
