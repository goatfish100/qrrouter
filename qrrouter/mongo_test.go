package main

import (
	"log"
	"os"
	"testing"

	"gopkg.in/mgo.v2/dbtest"
)

import "io/ioutil"

var ()

// var resource1 = Resource{
// 	"Id": "584f89b595f3b0c9cab0d38a", "Description": "test", "Protected": "false", "Action": "forward", "Address": "www.yahoo.com"}
var resource1 = Resource{Uuid: "034c1dd2-d454-11e6-a110-b34e9f2c654a", Description: "yahoo", Protected: "false", Action: "forward", Address: "http://www.yahoo.com"}
var resource2 = Resource{Uuid: "059edd7c-d454-11e6-92b9-374c2fc3d623", Description: "yahoo", Protected: "false", Action: "forward", Address: "http://www.yahoo.com"}

//{"_id" : ObjectId("584f89b595f3b0c9cab0d38a"), "description" : "test", "protected" : "false", "action" : "redirect", "address" : "https://www.yahoo.com" }
// Server holds the dbtest DBServer
var Server dbtest.DBServer

func insertFixtures() {

	if err := session.DB(MongoDBDatabase).C("res").Insert(resource1); err != nil {
		log.Println(err)
	}

	if err := session.DB(MongoDBDatabase).C("res").Insert(resource2); err != nil {
		log.Println(err)
	}
}
func init() {

}

func TestMain(m *testing.M) {
	//os.TempDir()
	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	// My main session var is now set to the temporary MongoDB instance
	session = Server.Session()
	insertFixtures()
	// Run the test suite
	retCode := m.Run()
	// Make sure we DropDatabase so we make absolutely sure nothing is left or locked while wiping the data and
	// close session

	session.DB(MongoDBDatabase).DropDatabase()
	session.Close()

	Server.Stop()

	// call with result of m.Run()
	os.Exit(retCode)
}
