.PHONY: test lint vet

test:
	go test ./...

lint:
	golint ./...

vet:
	go vet ./...
