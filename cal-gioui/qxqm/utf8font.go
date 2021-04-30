package main

import (
	"gioui.org/font/opentype"
	"gioui.org/text"
	"gioui.org/widget/material"
	"log"
	"os"
)

//utf8
func utf8Font() *material.Theme {
	f, err := os.Open("font/MaShanZheng_Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	ttf, err := opentype.ParseCollectionReaderAt(f)
	if err != nil {
		log.Println(err)
	}
	th := material.NewTheme([]text.FontFace{{Face: ttf}})
	return th
}
