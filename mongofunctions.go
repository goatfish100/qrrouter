package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"gitlab.com/qrhelper/qrhelperdatastructs"
)

//FetchResource - fetch resource
func FetchResource(resourceid string) datastructs.Resource {

	// change := mgo.Change{
	// 	Update:    bson.M{"$inc": bson.M{"n": 1}},
	// 	ReturnNew: true,
	// }

	c := MgoSession.DB(MongoDBDatabase).C("res")
	result := datastructs.Resource{}

	err = c.Find(bson.M{"uuid": resourceid}).One(&result)
	//ls err := result.Apply(change, &result)

	if err != nil {
		log.Print(err)
		//panic(err)
	}

	// // Update
	colQuerier := bson.M{"uuid": resourceid}
	change := bson.M{"$inc": bson.M{"accesscount": 1}}
	err = c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}
	log.Println("leaving FetchResource")
	return result
}
