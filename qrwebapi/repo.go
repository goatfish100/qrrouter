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
	o.Id = i
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
