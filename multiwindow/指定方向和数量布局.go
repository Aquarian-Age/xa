package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////按每行６个显示
type wly func(gtx layout.Context) layout.Dimensions

func (view wly) Run(w *Window) error {
	var ops op.Ops

	applicationClose := w.App.Context.Done()
	for {
		select {
		case <-applicationClose:
			return nil
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				view(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}
}
func (v *Letters) LayoutN(gtx layout.Context) layout.Dimensions {
	hGrid := outlay.Grid{
		Num:  6, //每行的最大数
		Axis: layout.Horizontal,
	}
	th := v.win.App.Theme
	//space:= layout.Rigid(layout.Spacer{Width: unit.Dp(13)}.Layout)
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return hGrid.Layout(gtx, len(v.items), func(gtx layout.Context, index int) layout.Dimensions {
				item := v.items[index]
				////////////
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
				//////////
				return material.Button(th, &item.Click, item.Text).Layout(gtx)
			})
		}),
	)
}

//////////////////////////////////////////////////////////////////////////////////////
