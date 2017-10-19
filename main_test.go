package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"io/ioutil"

	"gitlab.com/qrhelper/qrhelperdatastructs"

	"gopkg.in/mgo.v2/dbtest"
)

var ()

var resource1 = datastructs.Resource{Uuid: "333c1dd2-d454-11e6-a110-b34e9f2c654a", Description: "yahoo", Action: "forward", Address: "https://www.yahoo.com"}
var resource2 = datastructs.Resource{Uuid: "444edd7c-d454-11e6-92b9-374c2fc3d623", Description: "yahoo", Action: "forward", Address: "https://www.google.com"}
var resource3 = datastructs.Resource{Uuid: "444edd7c-d454-11e6-92b9-374c2fc3d624", Description: "redirect", Action: "redirect", Address: "/test"}
var resource4 = datastructs.Resource{Uuid: "444edd7c-d454-11e6-92b9-374c2fc3d625", Description: "forward", Action: "forward", Address: "/test"}
var resource5 = datastructs.Resource{Uuid: "444edd7c-d454-11e6-92b9-374c2fc3d626", Description: "S3 redirect", Action: "s3serve", Address: "/s3redirect"}
var resource6 = datastructs.Resource{Uuid: "444edd7c-d454-11e6-92b9-374c2fc3d627", Description: "S3 forward", Action: "s3redirect", Address: "/s3forward"}

var Server dbtest.DBServer

func insertFixtures() {
	if err := MgoSession.DB(MongoDBDatabase).C("res").Insert(resource1); err != nil {
		log.Fatal("Unable to insert test record resource1")
	}
	if err := MgoSession.DB(MongoDBDatabase).C("res").Insert(resource2); err != nil {
		log.Fatal("Unable to insert test record resource2")
	}
	if err := MgoSession.DB(MongoDBDatabase).C("res").Insert(resource3); err != nil {
		log.Fatal("Unable to insert test record resource3")
	}
	if err := MgoSession.DB(MongoDBDatabase).C("res").Insert(resource4); err != nil {
		log.Fatal("Unable to insert test record resource4")
	}
	if err := MgoSession.DB(MongoDBDatabase).C("res").Insert(resource5); err != nil {
		log.Fatal("Unable to insert test record resource4")
	}
	if err := MgoSession.DB(MongoDBDatabase).C("res").Insert(resource6); err != nil {
		log.Fatal("Unable to insert test record resource4")
	}
}

func TestMain(m *testing.M) {
	//os.TempDir()
	fmt.Println("TestMain - routine")
	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	// My main session var is now set to the temporary MongoDB instance
	MgoSession = Server.Session()
	insertFixtures()
	// Run the test suite

	fmt.Println("m.Run - start")
	retCode := m.Run()
	fmt.Println("m.Run - End")
	// Make sure we DropDatabase so we make absolutely sure nothing is left or locked while wiping the data and
	// close session

	MgoSession.DB(MongoDBDatabase).DropDatabase()
	MgoSession.Close()

	Server.Stop()

	fmt.Println("exiting test main")
	// call with result of m.Run()
	os.Exit(retCode)

}
