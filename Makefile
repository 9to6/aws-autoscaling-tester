PWD := $(shell pwd)
GOROOT := $(shell go env GOROOT)
GOPATH := $(shell go env GOPATH)
SRC := metric-generator

all: install

getdeps:
	@go get github.com/golang/lint/golint && echo "Installed golint:"
	@go get github.com/fzipp/gocyclo && echo "Installed gocyclo:"
	@go get github.com/client9/misspell/cmd/misspell && echo "Installed misspell:"
	@go get github.com/ahmetb/govvv && echo "Installed govvv"

verifiers: vet fmt lint cyclo spelling

vet:
	@echo "Running $@:"
	@go tool vet -all ./cmd
	#@go tool vet -shadow=true ./cmd

fmt:
	@echo "Running $@:"
	@gofmt -s -l cmd

lint:
	@echo "Running $@:"
	#@${GOPATH}/bin/golint ${SRC}/cmd...

cyclo:
	@echo "Running $@:"
	@${GOPATH}/bin/gocyclo -over 65 cmd

spelling:
	@echo "Running $@:"
	@${GOPATH}/bin/misspell -error cmd/**/*

build: getdeps verifiers

test: build
	@echo "Running all testing:"
	@go test $(GOFLAGS) ${SRC}/cmd...

install: build
	@go install

clean:
	@echo "Cleaning up all the generated files:"
	@find . -name '*.test' | xargs rm -fv
	@rm -rf build
	@rm -rf release

docker-build:
	@docker build -t dreg.be/tkwon/${SRC}:`cat VERSION` .

docker-push:
	@docker push dreg.be/tkwon/${SRC}:`cat VERSION`
