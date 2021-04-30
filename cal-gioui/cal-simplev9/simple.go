package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
	"github.com/gonoto/notosans"
	"liangzi.local/cal/cal"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(380), unit.Dp(420)),
			app.Title("干支历"),
		)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	ticker := time.NewTicker(time.Second)
	ui := NewUI()
	var ops op.Ops
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				ui.Layout(gtx)
				e.Frame(gtx.Ops)
			}
		case <-ticker.C:
			w.Invalidate()
		}
	}
}
func (ui *UI) Layout(gtx layout.Context) layout.Dimensions {
	inset := layout.UniformInset(unit.Dp(12)) //26
	//inset.Top = unit.Dp(26) //错开Android顶部栏
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			//实时显示时间
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				l := material.H6(ui.Theme, time.Now().Local().Format("2006-01-02 15:04:05"))
				l.Font = text.Font{Typeface: "Note"}
				maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
				l.Color = maroon
				l.Alignment = text.Middle
				return l.Layout(gtx)
			}),
			///btn
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return ui.btns.Layout(gtx, ui.Theme)
			}),
			////显示纪年
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				l := material.H6(ui.Theme, ymds)
				l.Font = text.Font{Typeface: "Noto"}
				maroon := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
				l.Color = maroon
				l.Alignment = text.Middle
				return l.Layout(gtx)
			}),
			//月历
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return ui.btns.LayoutBtnShow(gtx, ui.Theme)
			}),
		)
	})
}

func (bt *Btns) LayoutBtnShow(gtx layout.Context, th *material.Theme) layout.Dimensions {
	//本月
	for range bt.btn4.Clicks() {
		listM = yueli(y, m, d)
	}
	for range bt.btn.Clicks() {
		y = y + 1
		if y > 3498 && y < 1600 {
			y = time.Now().Year()
		}
		listM = yueli(y, m, d)
		ymds = ymd(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local))
	}
	for range bt.btn1.Clicks() {
		y = y - 1
		if y > 3498 && y < 1600 {
			y = time.Now().Year()
		}
		listM = yueli(y, m, d)
		ymds = ymd(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local))
	}
	for range bt.btn2.Clicks() {
		m += 1
		if m > 12 {
			y += 1
			m = 1
		}
		listM = yueli(y, m, d)
		ymds = ymd(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local))
	}
	for range bt.btn3.Clicks() {
		m -= 1
		if m < 1 {
			y -= 1
			m = 12
		}
		listM = yueli(y, m, d)
		ymds = ymd(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local))
	}
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return hGrid.Layout(gtx, bt.num, func(gtx layout.Context, i int) layout.Dimensions {
				s := fmt.Sprintf("%s    \n", listM.s[i]+"\n"+listM.l[i]+"\n"+listM.g[i])
				body := material.Body1(th, s)
				body.Font = text.Font{Typeface: "Noto"}
				body.Alignment = text.Start
				body.TextSize = unit.Dp(18)
				return body.Layout(gtx)
			})
		}),
	)
}
func (ck *Btns) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(3)}.Layout)
	btn := material.Button(th, &ck.btn, "下一年")
	btn1 := material.Button(th, &ck.btn1, "上一年")
	btn2 := material.Button(th, &ck.btn2, "下一月")
	btn3 := material.Button(th, &ck.btn3, "上一月")
	btn4 := material.Button(th, &ck.btn4, "本月")

	btn.Font = text.Font{Typeface: "Noto"}
	btn1.Font = text.Font{Typeface: "Noto"}
	btn2.Font = text.Font{Typeface: "Noto"}
	btn3.Font = text.Font{Typeface: "Noto"}
	btn4.Font = text.Font{Typeface: "Noto"}

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

func NewUI() *UI {
	listM = yueli(2020, 2, 1)
	ui := &UI{}
	ui.Theme = utf8Font()
	ui.btns.num = len(listM.l)

	return ui
}
func utf8Font() *material.Theme {
	fonts := gofont.Collection()
	fonts = appendOTC(fonts, text.Font{Typeface: "Noto"}, notosans.OTC())
	th := material.NewTheme(fonts)
	return th
}

func appendOTC(collection []text.FontFace, fnt text.Font, otc []byte) []text.FontFace {
	face, err := opentype.ParseCollection(otc)
	if err != nil {
		panic(fmt.Errorf("failed to parse font collection: %v", err))
	}
	return append(collection, text.FontFace{Font: fnt, Face: face})
}

func ymd(t time.Time) string {
	lunar := cal.NewLunarMonth(t.Year(), int(t.Month()), t.Day())
	var lunars string
	if lunar.LeapSday == 0 {
		lunars = fmt.Sprintf("%d年%s月(%s)%s %s",
			lunar.LY, cal.AliasLunarMonth(lunar.LM), lunar.Ydx, lunar.LRmc, lunar.Week)
	} else {
		lunars = fmt.Sprintf("%d年闰%s月(%s)%s %s",
			lunar.LY, cal.AliasLunarMonth(lunar.LeapM), lunar.LeapYdx, lunar.LeapRmc, lunar.Week)
	}
	gz := cal.NewGanZhiInfo(t.Year(), int(t.Month()), t.Day(), t.Hour())
	gzs := gz.YearGZ + "年 " + gz.MonthGZ + "月 " + gz.DayGZ + "日 " + gz.HourGZ + "时"

	arrt := cal.YueJiangJQArrT(t.Year())
	jq := cal.NewYueJiangJQ(arrt)
	todayJQs := jq.TodayJQ(t.Year(), t)

	s1 := gzs + "\n" + lunars + "\n" + todayJQs
	s := s1

	return s
}

//月历
func yueli(y, m, d int) M {
	yl := cal.NewYueLi(y, m, d)
	lunar := yl.LunarD
	solar := yl.SolarD
	gz := yl.GzD

	return M{s: solar, l: lunar, g: gz}
}

type M struct {
	s []string
	l []string
	g []string
}

var (
	hGrid = outlay.Grid{
		Num:  6, //每行的最大数
		Axis: layout.Horizontal,
	}
	t     = time.Now().Local()
	ymds  = ymd(time.Now().Local())
	listM M
	y     = t.Year()
	m     = int(t.Month())
	d     = t.Day()
)

type UI struct {
	Theme *material.Theme
	list  layout.List
	ed    widget.Editor
	btns  Btns
}
type Btns struct {
	btn  widget.Clickable
	btn1 widget.Clickable
	btn2 widget.Clickable
	btn3 widget.Clickable
	btn4 widget.Clickable //本月
	num  int
	ed   widget.Editor
}
