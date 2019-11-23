.PHONY: test lint vet

test:
	go test ./...

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm -f coverage.out
lint:
	golint ./...

vet:
	go vet ./...
