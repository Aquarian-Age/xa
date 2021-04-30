package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/color"
	"os"
)

func qrencode() error {
	/*	var err error
		var q *qrcode.QRCode
		q, err = qrcode.New("https://github.com/aquarian-Age/ccal", qrcode.Medium)
		checkError(err)
		err = q.WriteFile(128, "/tmp/ccal.png") //把文件输出到/tpm目录 名称为out.png 输出128x128*/
	err := qrcode.WriteColorFile("https://github.com/aquarian-Age/ccal", qrcode.Medium, 128, color.Black, color.White, "/tmp/ccal.png")
	checkError(err)
	return err
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
