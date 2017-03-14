# README #
QR HELPER Application - is a collection of Micro Services
QRWEBAPI is the Web Api portion
QRRouter is the router directing traffic to QR Resources


### This is the source code/info repository for the QR Helper project and associated items ###

* Quick summary
* Version

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

RUNNING:
qrrouter - can be run two ways - in Docker and orchestrated with docker-compose or with go with Mongo running on
locally or in Docker.

Running Locally
set two environmental variables - like this.  Note: please add a colon in front of port as in :8000
export MONGO_HOST=localhost
export QRROUTER_PORT=:8004

Dependencies/Building:
godeps is not used for package management and versioning -
In this folder there is GoDeps which has a json file with versions
and vendor file where the libraries are stored.

saving Dependencies
'godep save'

'godep go install' to build
