# Define the binary name
BINARY_NAME=./out/generate-npc
BINARY_DEMO_NAME=./demo/generate-npc

DEMO_ASSETS=demo/assets

# Define the main package
MAIN_PACKAGE=./cmd/generate-npc

# Go Bin
GO_BIN=$(shell which go)

# Ensure the out directory exists
out:
	mkdir -p out

# Install the required tools for go generators
install-tools:
	@echo "Parsing tools.go and installing dependencies..."
	@go list -e -f '{{join .Imports " "}}' tools.go | xargs -t -n 1 $(GO_BIN) install
	@echo "all tools installed"

# Build the Go application
build: out
	@go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Build binary to generate the demo.gif inside the vhs docker container
build-demo: out
	env GOOS=linux go build -o $(BINARY_DEMO_NAME) $(MAIN_PACKAGE)
	mkdir -p $(DEMO_ASSETS)
	cp assets/* $(DEMO_ASSETS)

# Run the Go application
run: build
	$(BINARY_NAME) $(filter-out $@,$(MAKECMDGOALS))

# Run tests
test:
	@go test -v ./...

# Run the linter
lint:
	@revive -config .revive.toml -formatter friendly ./...

# Run all checks
check: lint test

save-demo-gif: build-demo
	docker run --rm -v ${PWD}/demo:/vhs ghcr.io/charmbracelet/vhs demo.tape

# Clean up build artifacts
clean:
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_DEMO_NAME)
	rm -rf $(DEMO_ASSETS)

# Default target
all: build

# Hack to make run proxy the arguments to the binary
%:
	@true

.PHONY: out build run test lint check clean all
