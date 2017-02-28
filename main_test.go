package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"io/ioutil"

	"bitbucket.org/goatfish100/qrhelperdatastructs"

	"gopkg.in/mgo.v2/dbtest"
)

var ()

var resource1 = datastructs.Resource{Uuid: "333c1dd2-d454-11e6-a110-b34e9f2c654a", Description: "yahoo", Protected: "false", Action: "forward", Address: "https://www.yahoo.com"}
var resource2 = datastructs.Resource{Uuid: "444edd7c-d454-11e6-92b9-374c2fc3d623", Description: "yahoo", Protected: "false", Action: "forward", Address: "https://www.google.com"}

var Server dbtest.DBServer

func insertFixtures() {

	if err := mgosession.DB(MongoDBDatabase).C("res").Insert(resource1); err != nil {
		log.Fatal("Unable to insert test record resource1")
	}

	if err := mgosession.DB(MongoDBDatabase).C("res").Insert(resource2); err != nil {
		log.Fatal("Unable to insert test record resource2")
	}
}

func TestMain(m *testing.M) {
	//os.TempDir()
	fmt.Println("TestMain - routine")
	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	// My main session var is now set to the temporary MongoDB instance
	mgosession = Server.Session()
	insertFixtures()
	// Run the test suite

	fmt.Println("m.Run - start")
	retCode := m.Run()
	fmt.Println("m.Run - End")
	// Make sure we DropDatabase so we make absolutely sure nothing is left or locked while wiping the data and
	// close session

	mgosession.DB(MongoDBDatabase).DropDatabase()
	mgosession.Close()

	Server.Stop()

	fmt.Println("exiting test main")
	// call with result of m.Run()
	os.Exit(retCode)

}
