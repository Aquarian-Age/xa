package main

import (
	"log"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/Aquarian-Age/xa/pkg/giox"
)

func main() {
	go func() {
		w := app.NewWindow(app.Title("中文字体"),
			app.Size(unit.Dp(640), unit.Dp(640)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}
func loop(w *app.Window) error {
	th := giox.FontNoto()
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			render(gtx, th)
			e.Frame(gtx.Ops)
		}
	}
}
func render(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				h2 := material.H2(th, "中文字体显示")
				h2.Font = text.Font{Weight: text.Bold}
				h2.Alignment = text.Middle
				return h2.Layout(gtx)
			}),
		)
	})
}
