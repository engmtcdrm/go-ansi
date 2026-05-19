#!/bin/zsh

# alias go-run="go run ./example/main.go"
alias go-test-cover="go test -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html"
alias go-test="go test -timeout 30s ./..."
alias go-testv="go test -timeout 30s -v ./..."
