#!/bin/bash
# export CGO_ENABLED="0" 

CGO_ENABLED="0" go build -ldflags="-s -w" -tags trimpath -o run_wasm
CGO_ENABLED="0" GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -tags trimpath -o run_wasm.exe 
