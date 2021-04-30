#!/bin/bash
t=$(date +%Y-%m-%d)
go build -o /home/xuan/ccal/gioui/cal-v0.6.9-9-gioui-linux-$t -ldflags="-s -w" .
# GOOS=windows GOARCH=amd64 go build -o /home/xuan/ccal/gioui/cal-gioui-auth-windows-$t.exe -ldflags="-H windowsgui -s -w"
# go run gioui.org/cmd/gogio -o cal-v0.6.9-6-v3.apk -target android .
