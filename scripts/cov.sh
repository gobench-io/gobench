#!/bin/bash -e
# Run from directory above via ./scripts/cov.sh

GO111MODULE="off" go get github.com/mattn/goveralls
GO111MODULE="off" go get github.com/wadey/gocovmerge

rm -rf ./cov
mkdir cov
go test -v -failfast -covermode=atomic -coverprofile=./cov/agent.out ./agent
go test -v -failfast -covermode=atomic -coverprofile=./cov/executor.out ./executor
go test -v -failfast -covermode=atomic -coverprofile=./cov/logger.out ./logger
go test -v -failfast -covermode=atomic -coverprofile=./cov/master.out ./master
go test -v -failfast -covermode=atomic -coverprofile=./cov/smtp.out ./services/smtp
go test -v -failfast -covermode=atomic -coverprofile=./cov/web.out ./web

gocovmerge ./cov/*.out > acc.out
rm -rf ./cov

# If we have an arg, assume github action run and push to coveralls. Otherwise launch browser results
if [[ -n $1 ]]; then
    $HOME/gopath/bin/goveralls -coverprofile=acc.out -service travis-ci
    rm -rf ./acc.out
else
    go tool cover -html=acc.out
fi
