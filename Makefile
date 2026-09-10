.PHONY: run test

run:
	go run cmd/kvstore/main.go

test:
	go test ./internal/...