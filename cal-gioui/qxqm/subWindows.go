package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

//日的子窗口
func SubWD(day int, th *material.Theme, v *BFYQWindow, txt string) {
	for i := 1; i <= 31; i++ {
		if day == i {
			title := fmt.Sprintf("%d日", i)
			subWindowD(th, v, title, title+txt)
			break
		}
	}
}

//子窗口
func subWindowD(th *material.Theme, v *BFYQWindow, title, txt string) {
	ds := material.Body1(th, txt)
	size := ds.TextSize
	size.V *= 24
	v.win.App.NewWindow(title, bfyqView(func(gtx layout.Context) layout.Dimensions {
		gtx.Reset()
		return layout.Center.Layout(gtx, ds.Layout)
	}))
}

//时辰 子窗口
func subWindowH(th *material.Theme, v *BFYQWindow, title, txt string) {
	info := material.H6(th, txt)
	size := info.TextSize
	size.V *= 36 //窗口初始大小
	v.win.App.NewWindow(title,
		bfyqView(func(gtx layout.Context) layout.Dimensions {
			gtx.Reset() //清除主窗口内容
			return layout.Center.Layout(gtx, info.Layout)
		}),
		app.Size(size, size),
	)
}
