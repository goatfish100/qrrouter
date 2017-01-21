package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"../datastructures"
)

func FetchResource(resourceid string) datastructures.Resource {
	c := session.DB(MongoDBDatabase).C("res")
	result := datastructures.Resource{}
	err = c.Find(bson.M{"uuid": resourceid}).One(&result)

	if err != nil {
		log.Print(err)
	}
	log.Println("leaving FetchResource")
	return result
}
