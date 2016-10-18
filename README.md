# README #
This is the main repo - and info wiki for the QR Helper

### This is the source code/info repository for the QR Helper project and associated items ###

* Quick summary
* Version
* [Learn Markdown](https://bitbucket.org/tutorials/markdowndemo)

### How do I get set up? ###


* Install Golang
* On Mac - done with brew (see http://brew.sh/ for install details)
* brew install golang
* create folder gowork under home folder
* add line export GOPATH=/Users/jamesl/gowork to .bash_profile
* make folder src/bitbucket.org in go path
* cd into src/bitbucket.org
* git clone the source into the folder git
* git clone git@bitbucket.org:goatfish100/gorouter.git
* git clone
*  go get github.com/gorilla/mux
*  go get github.com/gorilla/sessions
*  go get github.com/vulcand/oxy/forward
*  go get github.com/vulcand/oxy/testutils
*  go get gopkg.in/mgo.v2
*  go get gopkg.in/mgo.v2
*  go get gopkg.in/mgo.v2
*  go get gopkg.in/mgo.v2/bson


### Database - Mongo ###

* Mongodb is started with the docker-compose
* load initial data from gorillaweb/database/resources.json
* mongoimport --db test4 --collection res --file resources.json

### Gollum Wiki ###

* also started with docker-compose up from docker-compose repository
* jl - not about local storage

* Saving the instance #docker save gollum > gollum.tar
