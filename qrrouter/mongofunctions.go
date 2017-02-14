package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"bitbucket.org/gorouter/datastructs"
)

//FetchResource - fetch resource
func FetchResource(resourceid string) datastructs.Resource {
	c := mgosession.DB(MongoDBDatabase).C("res")
	result := datastructs.Resource{}
	err = c.Find(bson.M{"uuid": resourceid}).One(&result)

	if err != nil {
		log.Print(err)
	}
	log.Println("leaving FetchResource")
	return result
}
