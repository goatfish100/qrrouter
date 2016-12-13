package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

func FetchResource(resourceid string) Resource {
	c := session.DB("resources").C("res")
	result := Resource{}
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(resourceid)}).One(&result)

	if err != nil {
		log.Fatal(err)
	}
	return result
}
