# Format code 
fmt:
	go fmt ./...

# View possible issues in codebase
vet:
	go vet ./...

# Add any missing libraries and remove unsed ones
tidy: fmt
	go mod tidy

# Build the executable binary for the application
build:
	@go build -o bin/

# Run the root command 
run: build
	@./bin/ccmc

# Clean project files and remove current binary in ./bin
clean:
	go clean
	rm ./bin/ccmc

# Run the tests for command serialization
test serializer:
	go test -v ./internal/pkg/serialization

# View the makefile commads
view:
	@cat Makefile
