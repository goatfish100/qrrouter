package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"gitlab.com/qrhelper/qrhelperdatastructs"
)

//FetchResource - fetch resource
func FetchResource(resourceid string) datastructs.Resource {
	println("inside fetchresource resourceid", resourceid)
	c := MgoSession.DB(MongoDBDatabase).C("res")
	result := datastructs.Resource{}
	err = c.Find(bson.M{"uuid": resourceid}).One(&result)
	println("fetched resource", result.Address)
	if err != nil {
		log.Print(err)
		//panic(err)
	}
	log.Println("leaving FetchResource")
	return result
}
