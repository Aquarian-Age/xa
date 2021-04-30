package main

import (
	"image/png"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		w := app.NewWindow(app.Title("img"))
		if err := loop(w); err != nil {
			log.Println(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
func loop(w *app.Window) error {
	var ops op.Ops
	input := "test.png"
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			drawIMG(input, gtx)
			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
}

///把图片显示到GUI界面
func drawIMG(input string, gtx layout.Context) {
	f, err := os.Open(input)
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
