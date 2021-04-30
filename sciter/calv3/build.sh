#!/bin/bash

t=$(date +%Y-%m-%d)
go build -o ~/ccal/sciter/ccal-sciter-v2-$t -ldflags="-s -w" .
cp ./cal.html ~/ccal/sciter/cal.html
