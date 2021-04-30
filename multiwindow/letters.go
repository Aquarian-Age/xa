// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"strings"
)

// Letters displays a clickable list of text items that open a new window.
type Letters struct {
	win   *Window
	items []*LetterListItem
	list  layout.List
}
type LetterListItem struct {
	Text  string
	Click widget.Clickable
}

// NewLetters creates a new letters view with the provided log.
func NewLetters() *Letters {
	view := &Letters{
		list: layout.List{Axis: layout.Vertical},
		//layout.Rigid(layout.Spacer{Width: unit.Dp(13)}.Layout), //小部件间距
	}

	for text := 'a'; text <= 'z'; text++ {
		view.items = append(view.items, &LetterListItem{Text: string(text)})
	}
	return view
}

// Run implements Window.Run method.
func (v *Letters) Run(w *Window) error {
	v.win = w
	return wly(v.LayoutN).Run(w)
	//return WidgetView(v.Layout).Run(w)
}

// Layout handles drawing the letters view.
func (v *Letters) Layout(gtx layout.Context) layout.Dimensions {
	th := v.win.App.Theme
	return v.list.Layout(gtx, len(v.items), func(gtx layout.Context, index int) layout.Dimensions {
		item := v.items[index]
		if item.Click.Clicked() {
			if strings.EqualFold("a", item.Text) { //如果点击了btn-a显示另外窗口　其他点击不显示
				newWindow() //新的窗口　可根据需求对窗口内容进行各种操作
			}
			if strings.EqualFold("b", item.Text) {
				bigText := material.H6(th, item.Text+" B") //次级窗口内容 纯文本显示
				size := bigText.TextSize
				size.V *= 8 //初次生成次级窗口的大小
				v.win.App.NewWindow(item.Text,
					WidgetView(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, bigText.Layout)
					}),
					app.Size(size, size),
				)
			}
			switch item.Text {
			case "c":
				bigText := material.H6(th, item.Text+" C") //次级窗口内容 纯文本显示
				size := bigText.TextSize
				size.V *= 18 //初次生成次级窗口的大小
				v.win.App.NewWindow(item.Text,
					WidgetView(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, bigText.Layout)
					}),
					app.Size(size, size),
				)
			case "d":
			case "e":
			case "f":
				bigText := material.H6(th, item.Text+" F") //次级窗口内容 纯文本显示
				size := bigText.TextSize
				size.V *= 32 //初次生成次级窗口的大小
				v.win.App.NewWindow(item.Text,
					WidgetView(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, bigText.Layout)
					}),
					app.Size(size, size),
				)
			}
		}
		//btn 竖列布局
		return layout.Flex{}.Layout(gtx,
			layout.Flexed(10, func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, &item.Click, item.Text).Layout(gtx)
			}),
			//layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			//	return material.Button(th, &item.Click, item.Text).Layout(gtx)
			//}),
		)
	})
}

///////////////////////////////////
var (
	wbtn widget.Clickable
)

func newWindow() {
	go func() {
		w := app.NewWindow()
		events := w.Events()
		th := material.NewTheme(gofont.Collection())
		var ops op.Ops
		var s string
		for {
			e := <-events
			switch e := e.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				if wbtn.Clicked() {
					s = "\n" + `wbtn checked -----` + "\n"
				}
				material.Body1(th, "New Window"+s).Layout(gtx)
				layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, &wbtn, "btnShow").Layout(gtx)
						}),
					)
				})
				e.Frame(gtx.Ops)
			}
		}
	}()
}
