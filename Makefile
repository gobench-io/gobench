GO ?= go

PACKAGES := $(shell go list ./...)
EXAMPLES := $(shell go list ./... | grep "examples")
UI_PATH := ./web/ui/gobench-ui
githash := `git rev-parse HEAD`
gittag=`git describe --tags $(git rev-list --tags --max-count=1)`
LDFLAGS="-X github.com/gobench-io/gobench/master.gitCommit=$(githash) -X github.com/gobench-io.gobench/master.gitTag=$(gittag)"

.PHONY: lint build examples tools ent statik pb

pb:
	protoc -I pb pb/executor.proto --go_out=plugins=grpc:./pb
	protoc -I pb pb/agent.proto --go_out=plugins=grpc:./pb

# run supported packages
lint-pkgs:
	GO111MODULE=off go get -u honnef.co/go/tools/cmd/staticcheck
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell

build-pkgs:
	GO111MODULE=off go get -u github.com/rakyll/statik
	GO111MODULE=off go get -u github.com/facebook/ent/cmd/entc

lint:
	$(exit $(go fmt ./... | wc -l))
	go vet ./...
	# ignore statik.go file
	find . -type f -name "*.go" ! -name "statik.go" | xargs misspell -error -locale US
	staticcheck $(go list ./... | grep -v ent/privacy)

build:
	go build -ldflags $(LDFLAGS) -o gobench ./

test:
	./scripts/cov.sh

examples:
	$(foreach var, $(EXAMPLES), go build -buildmode=plugin -o ./.bin/${var}.so $(var);)

# generate ent models
ent:
	entc generate ./ent/schema

# generate statik file for web ui
statik:
	statik -src=$(UI_PATH)/build -dest=./web -f

build-web-ui:
	cd $(UI_PATH) && yarn run build-noauth

update-statik: build-web-ui statik

run:
	go run -ldflags $(LDFLAGS) .
