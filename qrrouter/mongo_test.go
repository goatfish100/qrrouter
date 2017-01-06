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

func TestMongo(m *testing.M) {
	//os.TempDir()
	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	// My main session var is now set to the temporary MongoDB instance
	Session = Server.Session()
	Server.Stop()

	// Run the test suite
	retCode := m.Run()

	// call with result of m.Run()
	os.Exit(retCode)
}
