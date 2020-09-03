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

test:
	./scripts/cov.sh

examples:
	$(foreach var, $(EXAMPLES), go build -buildmode=plugin -o ./.bin/${var}.so $(var);)

# generate ent models
ent:
	entc generate ./ent/schema

# generate statik file for web ui
statik:
	statik -src=./web/ui/react-app/build -dest=./web -f
build-web-ui:
	cd web/ui/react-app && yarn build
update-statik: build-web-ui statik
run:
	go run .
