package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	//	"liangzi.local/qx"
)

func Show(gtx layout.Context, th *material.Theme) layout.Dimensions {
	inset := layout.UniformInset(unit.Dp(36))
	inset.Left = unit.Dp(480)
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Body1(th, "show " /* zhiQin() */).Layout(gtx)
			}),
		)
	})
}

// func zhiQin() string {
// 	t := T()
// 	wd := int(t.Weekday())
// 	if wd == 0 {
// 		wd = 7
// 	}
// 	year := t.Year()
// 	m := int(t.Month())
// 	d := t.Day()
// 	h := t.Hour()
// 	dgz, n := qx.DayGanZhi(year, m, d)
// 	hgz := qx.HourGanZhi(n, h)
// 	nqname := qx.GetNianQinName(year)
// 	yqname := qx.GetYueQinName(year, m)
// 	rqname := qx.GetRiQinName(wd, dgz)
// 	hqname := qx.GetShiQinName(wd, hgz)
// 	s := fmt.Sprintf("日幹支:%s 時幹支:%s\n年禽:%s 月禽:%s 日禽:%s 時禽:%s\n",
// 		dgz, hgz, nqname, yqname, rqname, hqname)
// 	return s
// }
