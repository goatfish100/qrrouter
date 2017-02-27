package datastructs

import "gopkg.in/mgo.v2/bson"

type Org struct {
	//Id         string `json:"id" bson:"_id,omitempty"`
	ID         bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Orgname    string        `json:"orgname"`
	Address    string        `json:"address"`
	City       string        `json:"city"`
	State      string        `json:"state"`
	Postalcode string        `json:"postalcode"`
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
type Resource struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Uuid        string `json:"uuid"`
	OrgId       string `json:"orgid" bson:"orgid,omitempty"`
	Description string `json:"Description"`
	Protected   string `json:"Protected"`
	Action      string `json:"Action"`
	Address     string `json:"Address"`
}

type JSONSuccess struct {
	Success string `json:"success"`
}