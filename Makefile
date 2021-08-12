test:
	@go get github.com/mfridman/tparse

	go test ./... -v -race -covermode=atomic  -json | tparse -all -dump


fmt:
	golangci-lint run -v
	go fmt ./...


release:
	npx release-it	

