package main

import (
	"gioui.org/layout"
	"gioui.org/op/paint"
	"image/png"
	"log"
	"os"
	"path"
	"path/filepath"
)

///把图片显示到GUI界面
func showPNG(starpng string, gtx layout.Context) {
	dir, _ := filepath.Abs(filepath.Dir("."))
	p := path.Join(dir, "bfyq")
	np := path.Join(p, starpng) //文件绝对路径
	///
	f, err := os.Open(np)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	imgOP := paint.NewImageOp(img)
	imgOP.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
}
