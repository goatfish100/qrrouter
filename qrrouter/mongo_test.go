package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"io/ioutil"

	"bitbucket.org/gorouter/datastructs"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/dbtest"
)

var ()

var resource1 = datastructs.Resource{Uuid: "034c1dd2-d454-11e6-a110-b34e9f2c654a", Description: "yahoo", Protected: "false", Action: "forward", Address: "https://www.yahoo.com"}
var resource2 = datastructs.Resource{Uuid: "059edd7c-d454-11e6-92b9-374c2fc3d623", Description: "yahoo", Protected: "false", Action: "forward", Address: "https://www.google.com"}

var Server dbtest.DBServer

func insertFixtures() {

	if err := session.DB(MongoDBDatabase).C("res").Insert(resource1); err != nil {
		log.Fatal("Unable to insert test record resource1")
	}

	if err := session.DB(MongoDBDatabase).C("res").Insert(resource2); err != nil {
		log.Fatal("Unable to insert test record resource2")
	}
}

func TestMain(m *testing.M) {
	//os.TempDir()
	fmt.Println("TestMain - routine")
	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	// My main session var is now set to the temporary MongoDB instance
	session = Server.Session()
	insertFixtures()
	// Run the test suite

	fmt.Println("m.Run - start")
	retCode := m.Run()
	fmt.Println("m.Run - End")
	// Make sure we DropDatabase so we make absolutely sure nothing is left or locked while wiping the data and
	// close session

	session.DB(MongoDBDatabase).DropDatabase()
	session.Close()

	Server.Stop()

	fmt.Println("exiting test main")
	// call with result of m.Run()
	os.Exit(retCode)

}

func TestGetResource(t *testing.T) {
	t.Log("TestGetResource Test")
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

func TestFetchResource(t *testing.T) {
	t.Log("TestFetchResource Test")

	testResult := FetchResource("059edd7c-d454-11e6-92b9-374c2fc3d623")
	t.Log("testresult is ", testResult)

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
