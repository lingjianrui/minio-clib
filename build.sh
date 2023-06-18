#!/bin/bash

if [ $1 == "win" ]; then
    GODEBUG=cgocheck=0 GOOS=windows CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ GOARCH="386" GOFLAGS=-tags=windows CGO_ENABLED=1 go build -buildmode=c-shared  -v -ldflags="-s -w" -trimpath -o bin/win/minioc.dll .
fi
if [ $1 == "mac" ]; then
    GODEBUG=cgocheck=0 GOOS=darwin GOARCH=amd64 GOFLAGS=-tags=darwin CGO_ENABLED=1 go build -buildmode=c-shared -v -ldflags="-s -w" -trimpath -o bin/mac/minioc.dylib .
fi
if [ $1 == "linux" ]; then
    CC=x86_64-unknown-linux-gnu-gcc CXX=x86_64-unknown-linux-gnu-g++ GODEBUG=cgocheck=0 GOOS=linux GOARCH=amd64 GOFLAGS=-tags=linux CGO_ENABLED=1 go build -buildmode=c-shared -o bin/linux/minioc.so .
fi

