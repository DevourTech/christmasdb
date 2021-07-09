.DEFAULT_GOAL := fmt

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...
