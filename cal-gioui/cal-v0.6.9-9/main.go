package main

import (
	"fmt"
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
	"image/color"
	"liangzi.local/cal/cal"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(620), unit.Dp(600)),
			app.MaxSize(unit.Dp(620), unit.Dp(800)),
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
	inset := layout.UniformInset(unit.Dp(3)) //26
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
			///////
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return ui.edits.LayoutEdit(gtx, ui.Theme)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return ui.edits.LayoutShow(gtx, ui.Theme)
			}),
		)
	})
}

func (ed *Edit) LayoutShow(gtx layout.Context, th *material.Theme) layout.Dimensions {
	if ed.btn.Clicked() {
		y := convEdit(ed.ed1.Text())
		m := convEdit(ed.ed2.Text())
		d := convEdit(ed.ed3.Text())
		h := convEdit(ed.ed4.Text())
		t := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
		eds := ymd(t)
		ymds = eds
	} else if ed.btn2.Clicked() {
		y := convEdit(ed.ed1.Text())
		m := convEdit(ed.ed2.Text())
		d := convEdit(ed.ed3.Text())
		eds := yueli(y, m, d)
		listM = eds
	} else if ed.btn3.Clicked() {
		y := convEdit(ed.ed1.Text())
		eds := jq24(y)
		ymds = eds
	}
	if ed.btn4.Clicked() {
		ymds = about
	}
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(13)}.Layout) //小部件间距
	///
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		////显示纪年
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			l := material.H6(th, ymds)
			l.Font = text.Font{Typeface: "Noto"}
			maroon := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
			l.Color = maroon
			l.Alignment = text.Middle
			return l.Layout(gtx)
		}),
		space,
		////月历
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return hGrid.Layout(gtx, ed.num, func(gtx layout.Context, i int) layout.Dimensions {
				s := fmt.Sprintf("%s    \n", listM.s[i]+"\n"+listM.l[i]+"\n"+listM.g[i])
				body := material.Body1(th, s)
				body.Font = text.Font{Typeface: "Noto"}
				body.Alignment = text.Start
				body.TextSize = unit.Dp(18)
				///今日颜色
				if findDay() == i {
					maroon := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
					body.Color = maroon
				}
				return body.Layout(gtx)
			})
		}),
	)
}

//
func (ed *Edit) LayoutEdit(gtx layout.Context, th *material.Theme) layout.Dimensions {
	borderWidth := float32(0.5)
	borderColor := color.NRGBA{A: 107}
	switch {
	case ed.ed1.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 1
	case ed.ed2.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 1
	case ed.ed3.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 1
	case ed.ed4.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 1
	}
	wb := widget.Border{
		Color:        borderColor,
		CornerRadius: unit.Dp(2),
		Width:        unit.Dp(borderWidth),
	}
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(3)}.Layout) //小部件间距
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.ed1, "year").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.ed2, "month").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.ed3, "day").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.ed4, "hour").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &ed.btn, "纪年")
			btn.Font = text.Font{Typeface: "Noto"}
			return btn.Layout(gtx)
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn2 := material.Button(th, &ed.btn2, "月历")
			btn2.Font = text.Font{Typeface: "Noto"}
			return btn2.Layout(gtx)
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn3 := material.Button(th, &ed.btn3, "24节气")
			btn3.Font = text.Font{Typeface: "Noto"}
			return btn3.Layout(gtx)
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn4 := material.Button(th, &ed.btn4, "关于")
			btn4.Font = text.Font{Typeface: "Noto"}
			//maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			//btn4.Color = maroon
			return btn4.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H6(th, "github.com/Aquarian-Age/ccal").Layout(gtx)
		}),
	)
}
func convEdit(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("convEdit:", err)
	}
	return i
}

//
func NewUI() *UI {
	listM = yueli(y, m, d)
	ui := &UI{}
	ui.Theme = utf8Font()
	ui.edits.Inits()
	ui.edits.num = 30
	//ui.btns.num = len(listM.l)

	return ui
}
func (ed *Edit) Inits() {
	ed.ed1.SingleLine = true
	ed.ed2.SingleLine = true
	ed.ed3.SingleLine = true
	ed.ed4.SingleLine = true
}

////
func findDay() int {
	day := t.Day()
	month := int(t.Month())
	ds := fmt.Sprintf("%d/%d", month, +day)
	ds = strings.Trim(ds, "\n")
	var index int
	for i := 0; i < len(listM.s); i++ {
		cut := strings.Trim(listM.s[i], "\n")
		if strings.EqualFold(ds, cut) {
			index = i
			break
		}
	}
	return index
}

//////
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

	sw := cal.NewShuoWangTS(t.Year(), int(t.Month()), t.Day())
	shuo := "朔: " + sw.ShuoTS
	wang := "望: " + sw.WangTS
	shang := "上弦: " + sw.ShangXianTS
	xia := "下弦: " + sw.XiaXianTS
	s1 := gzs + "\n" + lunars + "\n" + todayJQs + "\n\n"
	s2 := shuo + "\n" + wang + "\n" + shang + "\n" + xia + "\n\n"

	dmjs := cal.Dmj(t.Year(), int(t.Month()), t.Day())
	s := s1 + s2 + dmjs
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

//24节气
func jq24(y int) string {
	jqt := cal.YueJiangJQArrT(y)
	yjq := cal.NewYueJiangJQ(jqt)
	return yjq.JQ24()
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
	about = `
cal 一个简单干支历
默认月历显示从阴历初一开始
可计算时间范围1601~3498
mail: liangzi@yandex.com
本人保留版权！
`
)

type UI struct {
	Theme *material.Theme
	list  layout.List
	edits Edit
	//num  int
	//ed   widget.Editor
	//btns Btns
}
type Edit struct {
	ed1, ed2, ed3, ed4 widget.Editor
	btn                widget.Clickable
	btn2               widget.Clickable
	btn3               widget.Clickable
	num                int
	btn4               widget.Clickable //关于
}

type Btns struct {
	btn  widget.Clickable
	btn1 widget.Clickable
	btn2 widget.Clickable
	btn3 widget.Clickable
	btn4 widget.Clickable //本月
	num int
	ed widget.Editor
}

func (ck *Btns) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(3)}.Layout)
	btn := material.Button(th, &ck.btn, "下一年")
	btn1 := material.Button(th, &ck.btn1, "上一年")
	btn2 := material.Button(th, &ck.btn2, "下一月")
	btn3 := material.Button(th, &ck.btn3, "上一月")
	btn4 := material.Button(th, &ck.btn4, "本月")
	//	btn5 := material.Button(th, &ck.btn5, "日")

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
