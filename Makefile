build:
	@go build -o bin/library main.go

run: build
	@./bin/library

