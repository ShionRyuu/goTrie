.PHONY: compile fmt test clean

all: compile

compile:
	go build

fmt:
	go fmt

test: compile

clean:
	go clean
