test:
	@go get github.com/mfridman/tparse

	go test ./... -v -race -covermode=atomic  -json | tparse -all -dump


fmt:
	golangci-lint run -v
	go fmt ./...


release:
	@go get github.com/caarlos0/svu
	@echo "Releasing $$(svu next)..."
	
	@git tag $$(svu next) && git push --tags
	@echo "Done"