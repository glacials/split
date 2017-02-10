watch:
	reflex -r '\.go' make run

setup:
	go get github.com/cespare/reflex

run:
	go build -o $(GOPATH)/bin/splitbot
	splitbot

build:
	go build -o $(GOPATH)/bin/splitbot
