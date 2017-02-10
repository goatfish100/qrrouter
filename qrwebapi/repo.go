package main

import (
	"fmt"
	"log"

	"bitbucket.org/gorouter/datastructs"
	"gopkg.in/mgo.v2/bson"
)

//RepoFindResource - find resources by id
func RepoFindResource(id string) datastructs.Resource {
	c := session.DB("resources").C("orgusers")
	result := datastructs.Resource{}
	err = c.Find(bson.M{"_id": id}).One(&result)

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
