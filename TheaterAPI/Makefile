go-get:
	go get -u github.com/gorilla/mux
	go get -u go.mongodb.org/mongo-driver/bson
	go get -u go.mongodb.org/mongo-driver/bson/primitive
	go get -u go.mongodb.org/mongo-driver/mongo
	go get -u go.mongodb.org/mongo-driver/mongo/options

set-path:
	export GOPATH=$(GOPATH):$(shell pwd)

go-build:
	go build ./src/theaterapi

go-run: go-build
	./theaterapi