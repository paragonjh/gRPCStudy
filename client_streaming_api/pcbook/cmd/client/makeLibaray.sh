#!/bin/sh

# Make libVNetClient Lib
go build -o libVNetClient.so -buildmode=c-shared libVNetClient.go

# build test app
gcc -o test test.cc ./libVNetClient.so
