#! /bin/sh

mkdir -p src/addr
protoc -I=. --go_out=src/addr addr.book.proto

# export GOPATH=$GOPATH:$PWD


