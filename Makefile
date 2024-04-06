BIN="./bin"
SRC=$(shell find . -name "*.go")

.PHONY: fmt lint test install_deps clean set_version

default: all

all: fmt test

fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

test: install_deps
	$(info ******************** running tests ********************)
	go test -v ./...

richtest: install_deps
	$(info ******************** running tests with kyoh86/richgo ********************)
	richgo test -v ./...

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...

clean:
	rm -rf $(BIN)

bump_version:
	bash scripts/bump_version.sh
