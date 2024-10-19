.PHONY: build run clean test

# Build the application
build:
	go build -o myapp ./cmd/main.go

# Run the application
run: build
	./myapp

# Clean up the build artifacts
clean:
	rm -f myapp

# Run tests
test:
	go test ./...