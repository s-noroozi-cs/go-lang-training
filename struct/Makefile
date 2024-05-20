.DEFAULT_GOAL := run

clean: 
	go clean
fmt: clean
	go fmt ./...

vet: fmt
	go vet ./...

run: vet
	clear
	go run main.go
