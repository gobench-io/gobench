# build
FROM golang:1.14-alpine AS build

WORKDIR $GOPATH/src/github.com/gobench-io/gobench

COPY . .

RUN apk add build-base

RUN go build -o /gobench ./

# deployment
FROM golang:1.14-alpine

RUN apk add build-base

COPY --from=build gobench .

EXPOSE 8080

ENTRYPOINT [ "./gobench" ]