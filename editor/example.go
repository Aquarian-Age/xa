package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
	"log"
	"os"
)

func main() {
	ui := NewUI()
	go func() {
		w := app.NewWindow(app.Title("example-edit"), app.Size(unit.Dp(600), unit.Dp(600)))
		if err := ui.Run(w); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()
	app.Main()
}

type UI struct {
	Theme *material.Theme
	edits Edit
	//	btn Btn
}
type Edit struct {
	ed1, ed2, ed3, ed4 widget.Editor
	btn                widget.Clickable
	checked            bool
}

//type Btn struct {
//	btn     widget.Clickable
//	checked bool
//}

func (ui *UI) Run(w *app.Window) error {
	var ops op.Ops
	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			//th:=ui.Theme
			ui.LayoutUI(gtx)
			e.Frame(gtx.Ops)
		case key.Event:
			switch e.Name {
			case key.NameEscape:
				return nil
			}
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}

func NewUI() *UI {
	ui := &UI{}
	ui.Theme = material.NewTheme(gofont.Collection())
	ui.edits.Inits()
	return ui
}
func (ui *UI) LayoutUI(gtx layout.Context) layout.Dimensions {
	inset := layout.UniformInset(unit.Dp(9))
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return ui.edits.LayoutEdit(gtx, ui.Theme)
			}),
		)
	})
}

///
//func (btn *Btn)LayoutBtn(gtx layout.Context,th *material.Theme)layout.Dimensions  {
//	if btn.btn.Clicked(){
//		fmt.Println(btn.btn.Clicks())
//	}
//
//}
///
func (ed *Edit) LayoutEdit(gtx layout.Context, th *material.Theme) layout.Dimensions {
	borderWidth := float32(0.5)
	borderColor := color.NRGBA{A: 107}
	switch {
	case ed.ed1.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 2
	case ed.ed2.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 2
	case ed.ed3.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 2
	case ed.ed4.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 2
	}
	wb := widget.Border{
		Color:        borderColor,
		CornerRadius: unit.Dp(4),
		Width:        unit.Dp(borderWidth),
	}
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(9)}.Layout) //小部件间距
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.ed1, "edit1").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.ed2, "edit2").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.ed3, "edit3").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.ed4, "edit4").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &ed.btn, "btn")
			if btn.Button.Clicked() {
				fmt.Println("btn checked")
				fmt.Printf("edit1:%s edit2:%s edit3:%s edit4:%s ",
					ed.ed1.Text(), ed.ed2.Text(), ed.ed3.Text(), ed.ed4.Text())
			}
			return btn.Layout(gtx)
		}),
	)
}
func (ed *Edit) Inits() {
	ed.ed1.SingleLine = true
	ed.ed2.SingleLine = true
	ed.ed3.SingleLine = true
	ed.ed4.SingleLine = true
}
