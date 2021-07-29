package main

import (
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/Aquarian-Age/xa/pkg/giox"
)

func main() {
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
func loop(w *app.Window) error {
	fpath := "font/NotoSerifSC_Medium.otf"
	th := giox.FontX(fpath)
	var ops op.Ops
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				h3 := material.H3(th, time.Now().Format("2006-01-02 15:04:05"+"\n\n"+s))
				h3.Alignment = text.Middle
				h3.Layout(gtx)
				e.Frame(gtx.Ops)
			}
		case <-ticker.C:
			w.Invalidate()
		}
	}
}

var s = `长文本换行------------------->
			长文本中间
							长文本末尾
`
