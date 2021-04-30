package main

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image/color"
	"time"
)

//实时时间
func realTimeLayout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	inset := layout.UniformInset(unit.Dp(1))
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				l := material.H6(th, time.Now().Local().Format("2006-01-02 15:04:05"))
				maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
				l.Color = maroon
				l.Alignment = text.Middle
				return l.Layout(gtx)
			}),
		)
	},
	)
}
