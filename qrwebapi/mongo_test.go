package main

import (
	"fmt"
	"log"
	"testing"

	"io/ioutil"

	"bitbucket.org/gorouter/datastructs"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/dbtest"
)

var (
	//session   = mgo.Dial("localhost")
	resource1 = datastructs.Resource{Uuid: "034c1dd2-d454-11e6-a110-b34e9f2c654a", Description: "yahoo", Protected: "false", Action: "forward", Address: "https://www.yahoo.com"}
	resource2 = datastructs.Resource{Uuid: "059edd7c-d454-11e6-92b9-374c2fc3d623", Description: "yahoo", Protected: "false", Action: "forward", Address: "https://www.google.com"}

	org1 = datastructs.Org(datastructs.Org{
		Orgname:    "Rest Holdings",
		Address:    "123 H street",
		City:       "Culver",
		State:      "CA",
		Postalcode: "84109",
		Users: []datastructs.User{{
			Username: "freegyg",
			Email:    "freddy@yahoo.com",
			Name:     "Freddy G Spot",
			Password: "lsls",
		}, {
			Username: "toyo",
			Email:    "lsl@yahoo.com",
			Name:     "asdf",
			Password: "asdf",
		},
		}})
	org2 = datastructs.Org(datastructs.Org{
		Orgname:    "awake Holdings",
		Address:    "123 H street",
		City:       "Culver",
		State:      "CA",
		Postalcode: "84109",
		Users: []datastructs.User{{
			Username: "freegyg",
			Email:    "freddy@yahoo.com",
			Name:     "Freddy G Spot",
			Password: "lsls",
		}, {
			Username: "toyo",
			Email:    "lsl@yahoo.com",
			Name:     "asdf",
			Password: "asdf",
		},
		},
	})

	//Server := dbtest.DBServer
)

func insertFixtures() {

	if err := session.DB(MongoDBDatabase).C("res").Insert(resource1); err != nil {
		log.Fatal("Unable to insert test record resource1")
	}

	if err := session.DB(MongoDBDatabase).C("res").Insert(resource2); err != nil {
		log.Fatal("Unable to insert test record resource2")
	}
}

func TestMain(m *testing.M) {

	tempDir, _ := ioutil.TempDir("", "testing")

	var Server dbtest.DBServer
	//var Server = dbtest.DBServer
	Server.SetPath(tempDir)

	// My main session var is now set to the temporary MongoDB instance
	session := Server.Session()
	insertFixtures()
	// Run the test suite

	fmt.Println("m.Run - start")
	m.Run()
	fmt.Println("m.Run - End")
	// Make sure we DropDatabase so we make absolutely sure nothing is left or locked while wiping the data and
	// close session

	session.DB(MongoDBDatabase).DropDatabase()
	session.Close()
	//
	// Server.Stop()
	//
	// fmt.Println("exiting test main")
	// // call with result of m.Run()
	// os.Exit(retCode)

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

	testResult := RepoFindResource("059edd7c-d454-11e6-92b9-374c2fc3d623")
	if testResource(testResult, resource1) == false {
		t.Fail()
	}
	t.Log("testresult is ", testResult)

}

//
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

//
func testUser(rUser1 datastructs.User, rUser2 datastructs.User) bool {
	var resflag bool
	if rUser1.Email == rUser1.Email &&
		rUser1.Name == rUser2.Name &&
		rUser1.Password == rUser2.Password &&
		rUser1.Username == rUser1.Username {
		resflag = true
	} else {
		resflag = false
	}

	return resflag

}

//
//Test organization - to see they are the same ...
//TODO - need a way to check users
func testOrg(rOrg1 datastructs.Org, rOrg2 datastructs.Org) bool {
	var resflag bool
	if rOrg1.Address == rOrg2.Address &&
		rOrg1.City == rOrg2.City &&
		rOrg1.ID == rOrg2.ID &&
		rOrg1.Orgname == rOrg2.Orgname &&
		rOrg1.Postalcode == rOrg2.Postalcode &&
		rOrg1.State == rOrg2.State {
		resflag = true
	} else {
		resflag = false
	}

	return resflag

}
