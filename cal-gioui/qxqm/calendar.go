package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"time"
)

//日
type cald struct {
	day   int
	Click widget.Clickable
}

//月
type calm struct {
	month int
	Click widget.Clickable
}

//时间24小时制
type calh struct {
	hour  int
	Click widget.Clickable
}

//年
type YearBtns struct {
	btn0 widget.Clickable
	btn1 widget.Clickable
	btn2 widget.Clickable
	btn3 widget.Clickable
	btn4 widget.Clickable
}

//年份Btn布局
func (ck *YearBtns) LayoutYear(gtx layout.Context, th *material.Theme) layout.Dimensions {
	for range ck.btn0.Clicks() {
		if y > 3498 {
			y = t.Year()
		}
		if y < 1600 {
			y = t.Year()
		}
		y += 1
	}
	for range ck.btn1.Clicks() {
		if y > 3498 {
			y = t.Year()
		}
		if y < 1600 {
			y = t.Year()
		}
		y -= 1
	}
	for range ck.btn2.Clicks() {
		if y > 3498 {
			y = t.Year()
		}
		if y < 1600 {
			y = t.Year()
		}
		y += 10
	}
	for range ck.btn3.Clicks() {
		if y > 3498 {
			y = t.Year()
		}
		if y < 1600 {
			y = t.Year()
		}
		y -= 10
	}
	if ck.btn4.Clicked() {
		y = t.Year()
	}
	swapY = y
	////
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(3)}.Layout)
	btn := material.Button(th, &ck.btn0, "y+")
	btn1 := material.Button(th, &ck.btn1, "y-")
	btn2 := material.Button(th, &ck.btn2, "y+10")
	btn3 := material.Button(th, &ck.btn3, "y-10")
	btn4 := material.Button(th, &ck.btn4, "now")

	inset := layout.UniformInset(unit.Dp(3))
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return btn.Layout(gtx)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return btn1.Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return btn2.Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return btn3.Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return btn4.Layout(gtx)
			})
		}),
	)
}

//T
func T() time.Time {
	if swapY == 0 {
		swapY = t.Year()
	}
	if swapM == 0 {
		swapM = int(t.Month())
	}
	if swapD == 0 {
		swapD = t.Day()
	}
	if swapH == 0 {
		swapH = t.Hour()
	}
	return time.Date(swapY, time.Month(swapM), swapD, swapH, 0, 0, 0, time.Local)
}
