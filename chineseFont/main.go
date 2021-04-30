package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/gonoto/notosans"
)

// go get github.com/gonoto/notosans
// go mod download golang.org/x/sys

func main() {
	w := app.NewWindow()
	go func() {
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := utf8Font()
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			l := material.H1(th, "Gio中文")
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			l.Color = maroon
			l.Alignment = text.Middle
			l.Font = text.Font{Typeface: "Noto"}
			l.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}

//加载字体
func utf8Font() *material.Theme {
	fonts := gofont.Collection()
	fonts = appendOTC(fonts, text.Font{Typeface: "Noto"}, notosans.OTC())
	th := material.NewTheme(fonts)
	return th
}

func appendOTC(collection []text.FontFace, fnt text.Font, otc []byte) []text.FontFace {
	face, err := opentype.ParseCollection(otc)
	if err != nil {
		panic(fmt.Errorf("failed to parse font collection: %v", err))
	}
	return append(collection, text.FontFace{Font: fnt, Face: face})
}

//加载本地的指定字体
func yaheiTTf() *material.Theme {
	f, err := os.Open("/path/yahei.ttf")
	if err != nil {
		log.Fatal(err)
	}
	ttc, err := opentype.ParseCollectionReaderAt(f)
	if err != nil {
		log.Fatal(err)
	}
	th := material.NewTheme([]text.FontFace{{Face: ttc}})
	return th
}

//加载本地指定字体
func NotoBlack() *material.Theme {
	f, err := os.Open("/path/NotoSerifCJKsc-Black.otf")
	if err != nil {
		log.Fatal(err)
	}
	ttc, err := opentype.ParseCollectionReaderAt(f)
	if err != nil {
		log.Fatal(err)
	}
	th := material.NewTheme([]text.FontFace{{Face: ttc}})
	return th
}
