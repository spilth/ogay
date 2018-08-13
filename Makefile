all: deps vet test format

deps:
	go get -v github.com/stretchr/testify

vet:
	go vet

test:
	go test -v ./...

format:
	go fmt ./...
