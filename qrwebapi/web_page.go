package main

import (
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"bitbucket.org/gorouter/datastructs"
)

//Mongo database name
var MongoDBDatabase = "resources"

var mgosession, err = mgo.Dial("localhost")

//var qrresource QRResource
var orgs datastructs.Orgs
var users datastructs.Users

//var jsonsucess JsonSucess

// Give us some seed data
func insertRecords() {

	OrgCreate(datastructs.Org{
		Orgname:    "Rest Holdings",
		Address:    "123 H street",
		City:       "Culver",
		State:      "CA",
		Postalcode: "84109",
		Users: []datastructs.User{{
			Username: "freegyg",
			Email:    "freddy@yahoo.com",
			Name:     "Freddy G Spot",
			Password: "lsls",
		}, {
			Username: "toyo",
			Email:    "lsl@yahoo.com",
			Name:     "asdf",
			Password: "asdf",
		},
		}})
	OrgCreate(datastructs.Org{
		Orgname:    "awake Holdings",
		Address:    "123 H street",
		City:       "Culver",
		State:      "CA",
		Postalcode: "84109",
		Users: []datastructs.User{{
			Username: "freegyg",
			Email:    "freddy@yahoo.com",
			Name:     "Freddy G Spot",
			Password: "lsls",
		}, {
			Username: "toyo",
			Email:    "lsl@yahoo.com",
			Name:     "asdf",
			Password: "asdf",
		},
		},
	})

	ResourceInsert(datastructs.Resource{
		ID:          "1",
		Uuid:        "1232-1232-12312-123",
		OrgId:       "asdf",
		Description: "test resource",
		Protected:   "false",
		Action:      "forward",
		Address:     "https://surfhawaii.com",
	})

}

func main() {

	//insertRecords()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8001", router))
}
