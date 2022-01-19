include .env
export

test:
	@go get -d github.com/mfridman/tparse

	go test ./... -v -race -count=1 -covermode=atomic  -json | ~/go/bin/tparse -all -dump


fmt:
	golangci-lint run -v
	go fmt ./...


release:
	npx release-it	


precommit: fmt
	go build .