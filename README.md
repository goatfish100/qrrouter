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
* go get github.com/gorilla/mux
*  go get github.com/gorilla/sessions
*  go get github.com/vulcand/oxy/forward
*  go get github.com/vulcand/oxy/testutils
*  go gopkg.in/mgo.v2
*  go ge gopkg.in/mgo.v2
*  go get gopkg.in/mgo.v2
*  go get gopkg.in/mgo.v2/bson



* Configuration
* Dependencies
* Database configuration
* How to run tests
* Deployment instructions

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Gollum Wiki ###

* This instance pulled with docker instance from docker hub
* dependency - docker and local git? (may be in Docker instance)
1st step choose local/machine folder where files are stored and git initialize the folder
mkdir gollumwiki
cd gollumwiki
git init
now - you are ready to run the wiki
docker run -v `pwd`:/wiki -p 4567:80 gollum
... the docker image is now running on port 80 - internally - but serving on port 4567
http://localhost:4567 - is the addresss you can reach gollum/docker image at

* Saving the instance #docker save gollum > gollum.tar