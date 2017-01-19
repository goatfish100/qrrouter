package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

var qrresource QRResource
var orgs Orgs
var users Users

//var jsonsucess JsonSucess

// Give us some seed data
func init() {

	OrgCreate(Org{
		Orgname:    "Rest Holdings",
		Address:    "123 H street",
		City:       "Culver",
		State:      "CA",
		Postalcode: "84109",
		Users: []User{{
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
	OrgCreate(Org{
		Orgname:    "awake Holdings",
		Address:    "123 H street",
		City:       "Culver",
		State:      "CA",
		Postalcode: "84109",
		Users: []User{{
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

	UserCreate(User{Username: "bjorn balls",
		Email:    "test@yahoo.com",
		Name:     "bjorn balls",
		Password: "secret",
	})

}

func RepoFindResource(id string) Resource {
	c := session.DB("resources").C("orgusers")
	result := Resource{}
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)

	if err != nil {
		log.Fatal(err)
	}
	return result
}

func OrgCreate(o Org) Org {
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

func UserCreate(user User) User {
	// Insert Datas
	c := session.DB("resources").C("orgusers")
	err = c.Insert(user)

	if err != nil {
		panic(err)
	}
	users = append(users, user)
	return user
}
