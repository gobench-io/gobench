GO ?= go

PACKAGES := $(shell go list ./...)
EXAMPLES := $(shell go list ./... | grep "examples")

.PHONY: lint build examples tools ent statik

# run lint
lint-pkgs:
	GO111MODULE=off go get -u honnef.co/go/tools/cmd/staticcheck
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell

lint:
	$(exit $(go fmt ./... | wc -l))
	go vet ./...
	# ignore statik.go file
	find . -type f -name "*.go" ! -name "statik.go" | xargs misspell -error -locale US
	staticcheck $(go list ./... | grep -v ent/privacy)

build:
	go build -o gobench ./

install:
	go install

test: install
	go build -buildmode=plugin -o scenario/test/valid.so scenario/test/scripts/valid/valid.go
	go build -buildmode=plugin -o scenario/test/invalid.so scenario/test/scripts/invalid/invalid.go

	go test ./...

examples:
	$(foreach var, $(EXAMPLES), go build -buildmode=plugin -o ./.bin/${var}.so $(var);)

tools:
	go build -o ./.bin/github.com/gobench-io/gobench/tools/gobench-viewer github.com/gobench-io/gobench/tools/gobench-viewer

# generate ent models
ent:
	entc generate ./ent/schema

# generate statik file for web ui
statik:
	statik -src=./web/ui/react-app/build -dest=./web -f
build-web-ui:
	cd web/ui/react-app && yarn build
update-statik: build-web-ui
	statik
run:
	go run .
