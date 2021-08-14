# build
FROM golang:1.16-alpine AS build
WORKDIR $GOPATH/src/github.com/gobench-io/gobench

RUN apk add build-base git python2 nodejs npm

COPY . .

RUN npm i -g yarn
RUN make build-web-ui
RUN make build

# deployment
FROM golang:1.16-alpine

RUN apk add build-base gcc

COPY --from=build $GOPATH/src/github.com/gobench-io/gobench/gobench .

EXPOSE 8080

ENTRYPOINT [ "./gobench" ]
