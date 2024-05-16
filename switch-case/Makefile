.DEFAULT_GOAL := run

clean:
	reset
	go clean
fmt: clean
	go fmt ./...

vet: fmt
	go vet ./...

run: vet
	go run main.go
