.DEFAULT_GOAL := run

clean:
	go clean
fmt: clean
	go fmt ./...
vet: fmt
	go vet ./...
run: vet
	go run main.go
