.PHONY: test testv

test:
	go test -timeout 30s ./...

testv:
	go test -timeout 30s -v ./...
