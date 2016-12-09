package main

import "gopkg.in/mgo.v2/bson"

type Org struct {
	//Id         string `json:"id" bson:"_id,omitempty"`
	Id         bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Orgname    string
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Postalcode string `json:"postalcode"`
	Users      Users
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Orgs []Org

type Users []User
