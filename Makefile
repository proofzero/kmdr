SRC=$(shell find . -name "*.go")

ifeq (, $(shell which golangci-lint))
$(warning "could not find golangci-lint in $(PATH), run: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh")
endif

ifeq (, $(shell which richgo))
$(warning "could not find richgo in $(PATH), run: go get github.com/kyoh86/richgo")
endif

.PHONY: fmt lint test install_deps clean build install_local cue

default: all

all: fmt test build

fmt:
	$(info ******************** checking formatting ********************)
	#@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

lint:
	$(info ******************** running lint tools ********************)
	golangci-lint run -v

test: install_deps lint
	$(info ******************** running tests ********************)
	RICHGO_LOCAL=1 richgo test -v ./... -coverprofile .coverageprofile

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1
	go get -u github.com/kyoh86/richgo
	go get -v ./...

clean:
	bazel clean

build:
	$(info ******************** building binary ********************)
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel build //... 

install_local:
	$(info ******************** installing locally ********************)
	go install .

cue:
	cue eval ./cue/ -c -e space -t name=foo -t version=v1alpha1 --out=yaml
