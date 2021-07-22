# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

all: help
test :
	@echo " > run tests"
	@go test -v ./...
build :
	@echo " > build"
	@go build -o=bin/gotoozon
run :
	@echo " > run"
	@./bin/gotoozon
demo : build run

.PHONY: help
help:
	@echo "Choose a command to run:"
	@echo
	@echo "make test      - run all tests"
	@echo "make build     - build"
	@echo "make run       - run "
	@echo "make demo      - build and run demo"
