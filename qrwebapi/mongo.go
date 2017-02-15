package main

import (
	"log"

	"bitbucket.org/gorouter/datastructs"
)

// Insert resource into database
func InsertResource(resource datastructs.Resource) {

	if err := mgosession.DB(MongoDBDatabase).C("res").Insert(resource); err != nil {
		log.Println("Unable to insert resource", resource)
		//return false
	}

	//return true
}
