# build
FROM golang:1.14-alpine AS build

WORKDIR $GOPATH/src/github.com/gobench-io/gobench

COPY . .

RUN apk add build-base git

RUN make build

# deployment
FROM golang:1.14-alpine

RUN apk add build-base gcc

COPY --from=build gobench .

EXPOSE 8080

ENTRYPOINT [ "./gobench" ]