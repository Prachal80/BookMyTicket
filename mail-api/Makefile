all: clean

clean: 
	find . -name 'user' -type f -exec rm -f {} \; 
	go clean

go-get:
	rm -rf src/github.com
	go get -v go.mongodb.org/mongo-driver/mongo
	go get -v go.mongodb.org/mongo-driver/mongo/options
	go get -v github.com/gorilla/mux
	go get -v go.mongodb.org/mongo-driver/bson
	go get -v go.mongodb.org/mongo-driver/bson/primitive
	go get -v github.com/subosito/twilio
	go get -v gopkg.in/mgo.v2
	go get -v github.com/jordan-wright/email
	
set-path:
	export GOPATH=$(GOPATH):$(shell pwd)
run:
	go run ./src/mail

install:
	go install ./src

build:
	go build ./src/mail