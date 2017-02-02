package main

import (
	"os"
	"testing"

	"io/ioutil"

	"bitbucket.org/gorouter/datastructs"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/dbtest"
)

var ()

// var resource1 = Resource{
// 	"Id": "584f89b595f3b0c9cab0d38a", "Description": "test", "Protected": "false", "Action": "forward", "Address": "www.yahoo.com"}
var resource1 = datastructs.Resource{Uuid: "034c1dd2-d454-11e6-a110-b34e9f2c654a", Description: "yahoo", Protected: "false", Action: "forward", Address: "https://www.yahoo.com"}
var resource2 = datastructs.Resource{Uuid: "059edd7c-d454-11e6-92b9-374c2fc3d623", Description: "yahoo", Protected: "false", Action: "forward", Address: "https://www.google.com"}

//{"_id" : ObjectId("584f89b595f3b0c9cab0d38a"), "description" : "test", "protected" : "false", "action" : "redirect", "address" : "https://www.yahoo.com" }
// Server holds the dbtest DBServer
var Server dbtest.DBServer

//var server dbtest.DBServer.Session()

func insertFixtures() {

	if err := session.DB(MongoDBDatabase).C("res").Insert(resource1); err != nil {

	}

	if err := session.DB(MongoDBDatabase).C("res").Insert(resource2); err != nil {

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

func TestGetResource(t *testing.T) {
	//session = Server.session()
	c := session.DB(MongoDBDatabase).C("res")
	result := datastructs.Resource{}
	err = c.Find(bson.M{"uuid": "059edd7c-d454-11e6-92b9-374c2fc3d623"}).One(&result)

	t.Log("test address ", result.Address)
	if err != nil {
		t.Log(err)
	}
	if testResource(result, resource2) == false {
		t.Log("v")
	}

	// Try to find a resource that DOES NOT exist ... and ensure it isn't found
	result2 := datastructs.Resource{}
	err = c.Find(bson.M{"uuid": "does-not-exist"}).One(&result2)
	if result2.Address == "" {
		t.Log(result)
	} else {
		t.Fatal("mongo resource not equal")
	}
	//return result
}

func testResource(r1 datastructs.Resource, r2 datastructs.Resource) bool {
	var resflag bool
	if r1.Action == r2.Action &&
		r1.Address == r2.Address &&
		r1.Description == r2.Description &&
		r1.Protected == r2.Protected &&
		r1.Uuid == r2.Uuid {
		resflag = true
	} else {
		resflag = false
	}

	return resflag

}
