#!/bin/sh
rm -f *.go
protoc --go_out=plugins=grpc:. *.proto
go clean github.com/go-trellis/log-farm
go install -v