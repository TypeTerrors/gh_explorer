# Variables
BINARY_NAME=main
BUILD_DIR=./
SOURCE_DIR=./main.go

# ANSI color codes
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[0;33m
BLUE=\033[0;34m
MAGENTA=\033[0;35m
CYAN=\033[0;36m
NO_COLOR=\033[0m

# Phony targets for commands that don't represent files
.PHONY: build-local build-linux build-windows start publish create-new clean help confirm-publish

# HOWTO: Installation of Gum for interactive prompts
# Gum is used in this Makefile for interactive prompts when creating a new Azure Functions app.
# If you're not familiar with Gum, it's a tool for glamorous shell scripts that makes it easy
# to create interactive CLI applications. You'll need to have Gum installed to use the 'create-new' target.
# To install Gum, you can use the following command:
#
#   go install github.com/charmbracelet/gum@latest
#
# For more detailed installation instructions, refer to the official Gum repository:
# https://github.com/charmbracelet/gum
#
# Once Gum is installed, you can proceed with using 'make create-new' to interactively create
# a new Azure Functions app.

# Initialize project dependencies, including Gum
init:
	@echo ""
	@echo "Installing Gum for interactive prompts..."
	go get
	go mod tidy
	go install github.com/charmbracelet/gum@latest

# Build the application for local development
build-mac:
	@echo ""
	@echo GOOS=darwin GOARCH=amd64
	@gum spin --title "Building application for local development..." -- \
	go build -o $(BUILD_DIR)$(BINARY_NAME) $(SOURCE_DIR)

# Cross-compile the application for Linux
build-linux:
	@echo ""
	@GOOS=linux GOARCH=amd64
	@gum spin --title "Cross-compiling application for Linux..." -- \
	go build -o main

# Cross-compile the application for Windows
build-windows:
	@echo ""
	@GOOS=windows GOARCH=amd64
	@gum spin --title "Cross-compiling application for Windows..." -- \
	go build -o $(BINARY_NAME).exe


# Clean build artifacts
clean:
	@echo ""
	@echo "${RED}Cleaning up build artifacts...${NO_COLOR}"
	rm -f $(BUILD_DIR)$(BINARY_NAME) $(BUILD_DIR)$(BINARY_NAME).exe

# Help documentation for the Makefile
help:
	@echo ""
	@echo "${CYAN}Makefile Commands Description:${NO_COLOR}"
	@echo "  ${GREEN}make init${NO_COLOR} - Install project dependencies, including Gum for interactive prompts."
	@echo "    - Installs Go dependencies defined in go.mod."
	@echo "    - Prepares your environment for development and building."
	@echo ""
	@echo "  ${GREEN}make build-mac${NO_COLOR} - Compile the application for macOS."
	@echo "    - Sets GOOS and GOARCH for macOS and builds the binary."
	@echo ""
	@echo "  ${GREEN}make build-linux${NO_COLOR} - Cross-compile the application for Linux."
	@echo "    - Sets GOOS and GOARCH for Linux and builds the binary."
	@echo ""
	@echo "  ${GREEN}make build-windows${NO_COLOR} - Cross-compile the application for Windows."
	@echo "    - Sets GOOS and GOARCH for Windows and builds the binary with .exe extension."
	@echo ""
	@echo "  ${GREEN}make clean${NO_COLOR} - Remove build artifacts from your project directory."
	@echo "    - Cleans up any compiled binaries to ensure a fresh start for the next build."
	@echo ""
	@echo "${CYAN}Getting Started with the Project:${NO_COLOR}"
	@echo "  ${GREEN}Install Go${NO_COLOR}:"
	@echo "    - Download and install Go from ${BLUE}https://golang.org/dl/${NO_COLOR}."
	@echo "    - Follow the installation instructions for your OS."
	@echo "    - Verify installation by running ${BLUE}go version${NO_COLOR} in your terminal."
	@echo ""
	@echo "  ${GREEN}Starting the Project:${NO_COLOR}"
	@echo "    - Clone the repository and navigate into the directory."
	@echo "    - Run ${BLUE}make init${NO_COLOR} to install dependencies."
	@echo "    - Use ${BLUE}make build-local${NO_COLOR} (or relevant command for your target) to build the project."
	@echo ""
	@echo "${YELLOW}Additional Information:${NO_COLOR}"
	@echo "  This Makefile uses ${MAGENTA}Gum${NO_COLOR} for enhanced CLI interactivity. Install Gum with 'make init'."
	@echo "  More about Gum: ${BLUE}https://github.com/charmbracelet/gum${NO_COLOR}."
	@echo "  Ensure Go is installed and properly set up before proceeding with 'make' commands."
	@echo "  Project dependencies are managed through Go modules. Run 'make init' to install necessary dependencies."
	@echo ""