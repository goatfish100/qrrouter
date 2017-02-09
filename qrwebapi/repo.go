package main

import (
	"fmt"
	"log"

	"bitbucket.org/gorouter/datastructs"
	"gopkg.in/mgo.v2/bson"
)

//var qrresource QRResource
var orgs datastructs.Orgs
var users datastructs.Users

//var jsonsucess JsonSucess

// Give us some seed data
func init() {

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

//RepoFindResource - find resources by id
func RepoFindResource(id string) datastructs.Resource {
	c := session.DB("resources").C("orgusers")
	result := datastructs.Resource{}
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)

	if err != nil {
		log.Fatal(err)
	}
	return result
}

//OrgCreate create a new organization
func OrgCreate(o datastructs.Org) datastructs.Org {
	// Insert Datas
	c := session.DB("resources").C("orgusers")
	i := bson.NewObjectId()
	fmt.Println("The id is ", i)
	o.ID = i
	err = c.Insert(o)

	if err != nil {
		panic(err)
	}
	orgs = append(orgs, o)
	return o
}

//UserCreate ... create a new user
func UserCreate(user datastructs.User) datastructs.User {
	// Insert Datas
	c := session.DB("resources").C("orgusers")
	err = c.Insert(user)

	if err != nil {
		panic(err)
	}
	users = append(users, user)
	return user
}

// Insert resource into database
func ResourceInsert(resource datastructs.Resource) {

	c := session.DB("resources").C("res")
	err = c.Insert(resource)

	if err != nil {
		log.Println("Unable to insert resource", resource)
		//return false
	}

	//return true
}
