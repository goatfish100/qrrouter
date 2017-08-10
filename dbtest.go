package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"gitlab.com/qrhelper/qrhelperdatastructs"

	"os"



	"gopkg.in/mgo.v2"
)
var MongoHost = os.Getenv("MONGO_HOST")

//CookieStore - the cookie store/encrypt phrase

//AwsURL - Amazon AWS url
var AwsURL = os.Getenv("AWS_URL")

//AwsKey - Amazon AWS key
var AwsKey = os.Getenv("AWS_KEY")

//AwsPassPhrase - Amazon AWS passphrase
var AwsPassPhrase = os.Getenv("AWS_PASSPHRASE")

//AwsBucket - Amazon AWS Bucket
var AwsBucket = os.Getenv("AWS_BUCKET")

//MgoSession - Amazon AWS key
var MgoSession, err = mgo.Dial(MongoHost)

//MongoDBDatabase - mongo database name
var MongoDBDatabase = "resources"
//FetchResource - fetch resource
func main() {

	// change := mgo.Change{
	//         Update: bson.M{"$inc": bson.M{"AccessCount": 1}},
	//         ReturnNew: true,
	// }
	session, err := mgo.Dial(MongoHost)

	defer session.Close()

	coll := session.DB("resources").C("res")

	err = coll.Insert(M{"n": 42})

	session.SetMode(mgo.Monotonic, true)

	result := M{}
	info, err := coll.Find(M{"n": 42}).Apply(mgo.Change{Update: M{"$inc": M{"n": 1}}}, result)
	log.Println(result["n"])
	// c.Assert(err, IsNil)
	// c.Assert(result["n"], Equals, 42)
	// c.Assert(info.Updated, Equals, 1)
	// c.Assert(info.Matched, Equals, 1)
	// c.Assert(info.Removed, Equals, 0)
	// c.Assert(info.UpsertedId, IsNil)

	// col := MgoSession.DB("resources").C("res")
	// result := datastructs.Resource{}
	// //result2 := datastructs.Resource{}
	// //info, err := col.Find(bson.M{"_id": bson.ObjectIdHex("598a352e2d950edcd0b0a85d")}).Apply(change, &result)
  // err := col.Find(bson.M{"uuid": "S4b90c38-d455-11e6-83ef-9ba086d72a0a"}).One(&result)
	//
  // //fmt.Println(result.AccessCount)
	//
	// log.Print(result.AccessCount)
	// //log.Print(info)
	//
	// if err != nil {
	// 	log.Print(err)
	// 	//panic(err)
	// }
	// log.Println("leaving FetchResource")

}
