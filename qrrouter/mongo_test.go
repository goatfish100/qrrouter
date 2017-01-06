package main

import "testing"
import "gopkg.in/mgo.v2/dbtest"
import "io/ioutil"
import "os"

var ()

// Server holds the dbtest DBServer
var Server dbtest.DBServer

func init() {

}

func TestMain(m *testing.M) {
	//os.TempDir()
	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	// My main session var is now set to the temporary MongoDB instance
	session = Server.Session()
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
