# README #
QRRouter is a microservice application for taking URL's and
routing or proxying the request accordingly.  Written in GoLang,
it uses Gorilla/MUX for page routing and Mailgun's Vulcand to
reverse proxy traffic through it's URL - or redirect depending upon
how the URL resource is configured.

https://github.com/vulcand/vulcand


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
* make folder src/gitlab.com in go path
* cd into src/bitbucket.org
* git clone the source into the folder git
* git clone git@bitbucket.org:goatfish100/gorouter.git
* git clone
* The dependencies are now stored in godep folder


### Database - Mongo ###

* Mongodb is started with the docker-compose
* the database is currently manually loaded with
* load('mongojsload.js') command from the mongo shell
* mongojsload.js is currently in qrhelpermongo folder

### Gollum Wiki ###

* also started with docker-compose up from docker-compose repository
* jl - not about local storage

* Saving the instance #docker save gollum > gollum.tar

RUNNING:
qrrouter - can be run two ways - in Docker and orchestrated with docker-compose or with go with Mongo running on
locally or in Docker.

Running Locally - please include a .env file with the following
settings configured.  This gets picked up on start up of app
AWS_BUCKET=XXXX
AWS_PASSPHRASE=XXXX
AWS_KEY=xxxx
AWS_URL=s3.amazonaws.com
COOKIE_SECRET=whateveryouwant
MONGO_HOST=localhost
QRROUTER_PORT=8010




Dependencies/Building:
godeps is not used for package management and versioning -
In this folder there is GoDeps which has a json file with versions
and vendor file where the libraries are stored.

saving Dependencies
'godep save'

'godep go install' to build
