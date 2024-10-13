.DEFAULT_GOAL := build

package := github.com/tmsmr/go-lambda-terraform/fn
go_version := 1.23
go_arch := arm64

test:
	docker run --platform linux/$(go_arch) --rm \
		-v $(shell pwd):/go/src/$(package) \
		-w /go/src/$(package) \
		golang:$(go_version) \
		go test $(package)/internal/pkg/echo

compile:
	docker run --platform linux/$(go_arch) --rm \
		-v $(shell pwd):/go/src/$(package) \
		-w /go/src/$(package) \
		-e GOOS=linux -e GOARCH=$(go_arch) -e CGO_ENABLED=0 \
		golang:$(go_version) \
		go build -tags lambda.norpc -ldflags '-s -w' -o /go/src/$(package)/bin/bootstrap $(package)/lambda

clean:
	rm -rf bin

build: clean test compile

.PHONY: test compile clean build
