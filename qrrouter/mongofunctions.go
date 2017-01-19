package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

func FetchResource(resourceid string) Resource {
	c := session.DB(MongoDBDatabase).C("res")
	result := Resource{}
	err = c.Find(bson.M{"uuid": resourceid}).One(&result)

	if err != nil {
		log.Print(err)
	}
	log.Println("leaving FetchResource")
	return result
}
