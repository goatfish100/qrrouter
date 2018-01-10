package main

import (
	"testing"

	"gitlab.com/qrhelper/qrhelperdatastructs"
)

func TestFetchResource(t *testing.T) {
	t.Log("TestFetchResource Test")

	testResult := FetchResource("444edd7c-d454-11e6-92b9-374c2fc3d623")
	if testResource(testResult, resource2) == false {
		t.Log("feched resource is not same as inserted resource")
		t.Fail()
	}

}

func testResource(r1 datastructs.Resource, r2 datastructs.Resource) bool {
	var resflag bool
	if r1.Action == r2.Action &&
		r1.Address == r2.Address &&
		r1.Description == r2.Description &&
		r1.Uuid == r2.Uuid {
		resflag = true
	} else {
		resflag = false
	}

	return resflag

}
