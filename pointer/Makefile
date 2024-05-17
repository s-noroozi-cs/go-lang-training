.DEFAULT_GOAL := run

fmt: 
	go fmt ./...

vet: fmt
	go vet ./...

run: vet
	clear
	go run main.go
