## help - Display help about make targets for this Makefile
help:
	@cat Makefile | grep '^## ' --color=never | cut -c4- | sed -e "`printf 's/ - /\t- /;'`" | column -s "`printf '\t'`" -t

## install - Install globally from source
install:
	go build -o $(go env GOPATH)/bin/free

## clean - Clean the project
clean:
	rm free
	rm freedom
	rm $(go env GOPATH)/bin/free

## build - Build the project
build:
	go build -o free

## test - Test the project
test:
	go test -v

.PHONY: help install clean build test
