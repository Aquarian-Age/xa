package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
	"github.com/gonoto/notosans"
	"liangzi.local/cal"
	"liangzi.local/qx"
	"liangzi.local/sky/taisui"
	"liangzi.local/xjbfs"
)

var (
	windowCount int32
	zrBtn       widget.Clickable //择日窗口
)
var (
	listday                = &layout.List{Axis: layout.Vertical}
	ed1, ed2, ed3, ed4     widget.Editor
	btn1, btn2, btn3, btn4 widget.Clickable
	h                      = t.Hour()
	ymdZRs                 = ymdZR(t)
	showZRs                = ShowZR()
)

type (
	D = layout.Dimensions
	C = layout.Context
)
type zrWindos struct {
	*app.Window
	zrBtnClose widget.Clickable //关闭窗口
}

func zeRiWindow() {
	atomic.AddInt32(&windowCount, +1)
	go func() {
		w := new(zrWindos)
		w.Window = app.NewWindow(app.Title("xjbfs"))
		if err := w.ZrRiLoop(w.Events()); err != nil {
			log.Fatal(err)
		}
		if c := atomic.AddInt32(&windowCount, -1); c == 0 {
			os.Exit(0)
		}
	}()
}

func (w *zrWindos) ZrRiLoop(events <-chan event.Event) error {
	ticker := time.NewTicker(time.Second)
	th := utf8Font()
	var ops op.Ops
	for {
		select {
		case e := <-events:
			switch e := e.(type) {
			case system.FrameEvent:
				if w.zrBtnClose.Clicked() {
					w.Close()
				}
				gtx := layout.NewContext(&ops, e)
				realTLayout(gtx, th) //实时时间
				//横向布局输入框 按钮
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return EditLayoutV(gtx, th)
					}),
				)
				//显示
				render(gtx, th)
				e.Frame(gtx.Ops)
			case system.DestroyEvent:
				return e.Err
			}
		case <-ticker.C:
			w.Invalidate()
		}
	}
}
func realTLayout(gtx layout.Context, th *material.Theme) layout.Dimensions {
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

//布局输入框+按钮
func EditLayoutV(gtx layout.Context, th *material.Theme) layout.Dimensions {
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(3)}.Layout) //小部件横向间隔空间
	ed1.SingleLine = true
	ed2.SingleLine = true
	ed3.SingleLine = true
	ed4.SingleLine = true
	borderWidth := float32(0.3)
	borderColor := color.NRGBA{A: 107}

	if ed1.Focused() || ed2.Focused() || ed3.Focused() || ed4.Focused() { //输入状态显示颜色变化
		borderColor = th.Palette.ContrastBg
		borderWidth = 1
	}
	wb := widget.Border{
		Color:        borderColor,
		CornerRadius: unit.Dp(2),
		Width:        unit.Dp(borderWidth),
	}

	ln := layout.UniformInset(unit.Dp(9)) //输入框高度
	inset := layout.UniformInset(unit.Dp(1))
	inset.Top = unit.Dp(24) //小部件放在实时时间下方
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return ln.Layout(gtx,
						material.Editor(th, &ed1, "y").Layout)
				})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return ln.Layout(gtx,
						material.Editor(th, &ed2, "m").Layout)
				})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return ln.Layout(gtx,
						material.Editor(th, &ed3, "d").Layout)
				})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return ln.Layout(gtx,
						material.Editor(th, &ed4, "h").Layout)
				})
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, &btn1, "show").Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, &btn2, "h-").Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, &btn3, "h+").Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, &btn4, "me").Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				h6 := material.H6(th, "GUI: https://github.com/Aquarian-Age/ccal")
				maroon := color.NRGBA{R: 0, G: 0, B: 255, A: 255}
				h6.Color = maroon
				return h6.Layout(gtx)
			}),
		)
	})
}

//显示
func render(gtx C, th *material.Theme) D {
	if btn1.Clicked() {
		y = convEdit(ed1.Text())
		m = convEdit(ed2.Text())
		d = convEdit(ed3.Text())
		h = convEdit(ed4.Text())
		t = time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
		ymdZRs = ymdZR(t)
		showZRs = ShowZR()
	}
	//上一小时
	if btn2.Clicked() {
		h -= 1
		if h < 1 {
			d -= 1
			h = 23
		}
		if ((y%4 == 0 && y%100 != 0) || y%400 == 0) && m == 3 {
			if d < 1 {
				m = 2
				d = 1
			}
		}
		if !((y%4 == 0 && y%100 != 0) || y%400 == 0) && m == 3 {
			if d < 1 {
				m = 2
				d = 1
			}
		}
		if m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
			if d < 1 {
				m -= 1
				if m < 1 {
					m += 12
				}
				d = 30
			}
		}
		if m == 1 || m == 2 || m == 4 || m == 6 || m == 9 || m == 11 {
			if d < 1 {
				m -= 1
				if m < 1 {
					m += 12
				}
				d = 31
			}
		}
		if m > 12 {
			y -= 1
		}
		t = time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
		ymdZRs = ymdZR(t)
		showZRs = ShowZR()
	}
	//下一小时
	if btn3.Clicked() {
		h += 1
		if h > 23 {
			d += 1
			h = 0
		}
		if ((y%4 == 0 && y%100 != 0) || y%400 == 0) && m == 2 {
			if d > 29 {
				m += 1
				d = 1
			}
		} else if !((y%4 == 0 && y%100 != 0) || y%400 == 0) && m == 2 {
			if d > 28 {
				m += 1
				d = 1
			}
		}
		if m == 4 || m == 6 || m == 9 || m == 11 {
			if d > 30 {
				m += 1
				d = 1
			}
		}
		if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
			if d > 31 {
				m += 1
				d = 1
			}
		}
		if m > 12 {
			y += 1
		}
		t = time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
		ymdZRs = ymdZR(t)
		showZRs = ShowZR()
	}
	if btn4.Clicked() {
		ymdZRs = ymdZR(t)
		showZRs = aboutZR
	}
	inset := layout.UniformInset(unit.Dp(3)) //限制宽度
	inset.Top = unit.Dp(43)                  //距离上面输入小部件的间隔
	return inset.Layout(gtx, func(gtx C) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return layout.Inset{Top: unit.Sp(14)}.Layout(gtx, func(gtx C) D {
					//lunar := ymdZR(t)
					l := material.H6(th, ymdZRs)
					l.TextSize = unit.Dp(14)
					l.Alignment = text.Middle
					c := color.NRGBA{R: 255, A: 200}
					l.Color = c
					l.Font = text.Font{Typeface: "Noto"}
					return l.Layout(gtx)
				})
			}),
			layout.Flexed(1, func(gtx C) D {
				return inset.Layout(gtx, func(gtx C) D {
					return listday.Layout(gtx, 1, func(gtx C, i int) D {
						l := material.Body1(th, showZRs)
						l.Font = text.Font{Typeface: "Noto"}
						return l.Layout(gtx)
					})
				})
			}),
		)
	})
}

func ymdZR(t time.Time) string {
	lunar := cal.NewLunarMonth(t.Year(), int(t.Month()), t.Day())
	var lunars string
	if lunar.LeapM != 0 && lunar.LeapRmc != "" {
		lunars = fmt.Sprintf("%d年闰%s月(%s)%s %s",
			lunar.LY, cal.AliasLunarMonth(lunar.LeapM), lunar.LeapYdx, lunar.LeapRmc, lunar.Week)
	} else if lunar.LeapRmc == "" {
		lunars = fmt.Sprintf("%d年%s月(%s)%s %s",
			lunar.LY, cal.AliasLunarMonth(lunar.LM), lunar.Ydx, lunar.LRmc, lunar.Week)
	}

	gz := cal.NewGanZhiInfo(t.Year(), int(t.Month()), t.Day(), t.Hour())
	gzs := gz.YearGZ + "年 " + gz.MonthGZ + "月 " + gz.DayGZ + "日 " + gz.HourGZ + "时"

	arrt := cal.YueJiangJQArrT(t.Year())
	jq := cal.NewYueJiangJQ(arrt)
	todayJQs := jq.TodayJQ(t.Year(), t)
	//月将
	yj := jq.YueJiang(t)
	yjg := yj.Star
	yjm := "月将: " + yj.Name + " " + yjg
	s1 := gzs + "\n" + lunars + "\n" + todayJQs + "\n" + yjm

	return s1
}

//協紀辯方擇日
func ShowZR() string {
	t = time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	gz := cal.NewGanZhiInfo(y, m, d, h)
	xjbf := xjbfs.NewXJBF(gz.YearGZ, gz.MonthGZ, gz.DayGZ, gz.HourGZ)
	jz60 := xjbfs.MakeJZ60()
	ji, taisuiwugui, sha, xiong := xjbf.GetNianBiaoS(jz60)
	//年表
	x1 := fmt.Sprintf("♡岁吉♡: %s\n", ji)
	x2 := fmt.Sprintf("♥太岁五鬼♥: %s\n", taisuiwugui)
	x3 := fmt.Sprintf("♥岁煞♥: %s\n", sha)
	x4 := fmt.Sprintf("♥岁凶♥: %s\n", xiong)
	//------------
	ts := taiSui(gz.YearGZ) //生肖太岁

	nian := x1 + x2 + x3 + x4 + ts
	//月表
	jcday := xjbf.JianChuDay()
	krb := xjbf.KaiRi(jcday)
	st := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	jqt := cal.YueJiangJQArrT(y)
	yzl := xjbf.Yzl()
	mjs, mxs := xjbf.GetYueBiaoS(krb, st, jqt)
	x5 := fmt.Sprintf("\n【月总论】:\n%s", yzl)
	//------------------
	sw := cal.NewShuoWangTS(y, m, d)
	shuo := "\n●朔●: " + sw.ShuoTS
	wang := "○望○: " + sw.WangTS
	shang := "☽上弦☽: " + sw.ShangXianTS
	xia := "☾下弦☾: " + sw.XiaXianTS
	s2 := shuo + "\n" + wang + "\n" + shang + "\n" + xia + "\n"

	yue := x5 + s2

	//黄黑
	hhd := xjbf.HuangHeiDay()
	hhh := xjbf.HuangHeiHour()
	RiQin, _ := xjbf.QinXingDay(t)
	riQinInfo := qx.QinXingMiaoYong(RiQin) //禽星妙用

	//日建除
	x8 := fmt.Sprintf("\n☺日建除☺: %s\n", jcday)
	x9 := fmt.Sprintf("♠日黄黑♠: %s\n", hhd)
	x6 := fmt.Sprintf("♤日吉♤: %s\n", mjs)
	x7 := fmt.Sprintf("♠日凶♠: %s\n", mxs)
	xrq := fmt.Sprintf("☆日禽☆:%s\n", RiQin)
	qxmy := fmt.Sprintf("%s\n", riQinInfo)

	//日表
	rjs, rxs := xjbf.GetRiBaoS(jz60)

	ri := x8 + x9 + x6 + x7 + xrq + qxmy
	//时辰
	x10 := fmt.Sprintf("\n♣时黄黑♣: %s\n", hhh)
	x11 := fmt.Sprintf("♧时吉♧: %s\n", rjs)
	x12 := fmt.Sprintf("♦时凶♦: %s\n", rxs)
	shi := x10 + x11 + x12
	//其他
	jf, jf2 := xjbf.JinFu()
	x13 := fmt.Sprintf("《金符经》: %s %s\n", jf, jf2)
	yl := cal.NewYueLi(y, m, d)
	days := yl.GzD
	qs := xjbf.JinFuMonthQS(days, gz.DayGZ)
	x14 := fmt.Sprintf("♨金符经-本月七煞日♨: %s\n", qs)
	//---------
	dmjs := cal.Dmj(y, m, d)
	other := x13 + x14
	all := nian + yue + ri + shi + other + "\n卐地母经卐:\n" + dmjs
	return all
}

var aboutZR = `
	协纪辩方书择日
	朔 望 上弦 下弦 
	月将 节气 十二星宫
	金符经七煞
	mail: liangzi@yandex.com
	本人保留版权!
`

/////////////////////////////

func main() {
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(600), unit.Dp(640)),
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
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			//实时显示时间
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				l := material.H6(ui.Theme, time.Now().Local().Format("2006-01-02 15:04:05"))
				//l.Font = text.Font{Typeface: "Noto"}
				maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
				l.Color = maroon
				l.Alignment = text.Middle
				return l.Layout(gtx)
			}),
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
		y = convEdit(ed.ed1.Text())
		m = convEdit(ed.ed2.Text())
		d = convEdit(ed.ed3.Text())
		h := convEdit(ed.ed4.Text())
		t = time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
		eds := ymd(t)
		ymds = eds
	} else if ed.btn2.Clicked() {
		y = convEdit(ed.ed1.Text())
		m = convEdit(ed.ed2.Text())
		d = convEdit(ed.ed3.Text())
		eds := yueli(y, m, d)
		listM = eds
	} else if ed.btn3.Clicked() {
		y = convEdit(ed.ed1.Text())
		eds := jq24(y)
		ymds = eds
	}
	if ed.btn4.Clicked() {
		ymds = about
	}
	if zrBtn.Clicked() {
		zeRiWindow()
	}
	//////////////
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(13)}.Layout) //小部件间距
	///
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		////显示纪年
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			l := material.H6(th, ymds)
			l.Font = text.Font{Typeface: "Noto"}
			l.TextSize = unit.Dp(18)
			maroon := color.NRGBA{R: 255, G: 69, B: 0, A: 255}
			l.Color = maroon
			l.Alignment = text.Middle
			return l.Layout(gtx)
		}),
		space,
		////月历
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return hGrid.Layout(gtx, len(listM.s), func(gtx layout.Context, i int) layout.Dimensions {
				s := fmt.Sprintf("    %s\n", listM.s[i]+"\n"+listM.g[i]+"\n"+listM.l[i]+"\n"+listM.qin[i])
				body := material.Body1(th, s)
				body.Font = text.Font{Typeface: "Noto"}
				body.Alignment = text.Middle
				body.TextSize = unit.Dp(16)
				///今日颜色
				if findDay() == i && listM.qsb[i] == true {
					maroon := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
					body.Color = maroon
				}
				if findDay() == i && listM.qsb[i] == false {
					maroon := color.NRGBA{R: 0, G: 0, B: 255, A: 255}
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
			btnZr := material.Button(th, &zrBtn, "择日")
			btnZr.Font = text.Font{Typeface: "Noto"}
			return btnZr.Layout(gtx)
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn4 := material.Button(th, &ed.btn4, "关于")
			btn4.Font = text.Font{Typeface: "Noto"}
			return btn4.Layout(gtx)
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
	//ui.edits.num = len(listM.s)

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

//选择自定义字体
func yaheiTTf() *material.Theme {
	f, err := os.Open("./yahei.ttf")
	if err != nil {
		log.Fatal(err)
	}
	ttc, err := opentype.ParseCollectionReaderAt(f)
	if err != nil {
		log.Fatal(err)
	}
	th := material.NewTheme([]text.FontFace{{Face: ttc}})
	return th
}

//谷歌字体
func NotoBlack() *material.Theme {
	f, err := os.Open("./NotoSerifCJKsc-Black.otf")
	if err != nil {
		log.Fatal(err)
	}
	ttc, err := opentype.ParseCollectionReaderAt(f)
	if err != nil {
		log.Fatal(err)
	}
	th := material.NewTheme([]text.FontFace{{Face: ttc}})
	return th
}

////
func ymd(t time.Time) string {
	lunar := cal.NewLunarMonth(t.Year(), int(t.Month()), t.Day())
	var lunars string
	if lunar.LeapM != 0 && lunar.LeapRmc != "" {
		lunars = fmt.Sprintf("%d年闰%s月(%s)%s %s",
			lunar.LY, cal.AliasLunarMonth(lunar.LeapM), lunar.LeapYdx, lunar.LeapRmc, lunar.Week)
	} else if lunar.LeapRmc == "" {
		lunars = fmt.Sprintf("%d年%s月(%s)%s %s",
			lunar.LY, cal.AliasLunarMonth(lunar.LM), lunar.Ydx, lunar.LRmc, lunar.Week)
	}
	gz := cal.NewGanZhiInfo(t.Year(), int(t.Month()), t.Day(), t.Hour())
	gzs := gz.YearGZ + "年 " + gz.MonthGZ + "月 " + gz.DayGZ + "日 " + gz.HourGZ + "时"
	nys := cal.NaYin(gz.YearGZ, gz.MonthGZ, gz.DayGZ, gz.HourGZ)

	arrt := cal.YueJiangJQArrT(t.Year())
	jq := cal.NewYueJiangJQ(arrt)
	todayJQs := jq.TodayJQ(t.Year(), t)
	yuejiang := jq.YueJiang(t)
	yjs := "月将: " + yuejiang.Name + " " + yuejiang.Star

	sw := cal.NewShuoWangTS(t.Year(), int(t.Month()), t.Day())
	shuo := "朔: " + sw.ShuoTS
	wang := "望: " + sw.WangTS
	shang := "上弦: " + sw.ShangXianTS
	xia := "下弦: " + sw.XiaXianTS
	s1 := gzs + "\n" + nys + "\n" + lunars + "\n" + todayJQs + "\n" + yjs + "\n\n"
	s2 := shuo + "\n" + wang + "\n" + shang + "\n" + xia + "\n"

	dmjs := cal.Dmj(t.Year(), int(t.Month()), t.Day())

	ts := taiSui(gz.YearGZ)
	s := s1 + s2 + ts + dmjs
	return s
}

//月历
func yueli(y, m, d int) M {
	yl := cal.NewYueLi(y, m, d)
	lunar := yl.LunarD
	solar := yl.SolarD
	gz := yl.GzD
	q := yl.RiQin
	qb := yl.QiShaB
	return M{s: solar, l: lunar, g: gz, qin: q, qsb: qb}
}

//24节气
func jq24(y int) string {
	jqt := cal.YueJiangJQArrT(y)
	yjq := cal.NewYueJiangJQ(jqt)
	return yjq.JQ24()
}

//生肖太岁
func taiSui(ygz string) string {
	zhi := taisui.ZhiTaiSui(ygz)
	chong := taisui.ChongTaiSui(ygz)
	xing := taisui.XingTaiSui(ygz)
	hai := taisui.HaiTaiSui(ygz)
	s := zhi + "\n" + xing + "\n" + chong + "\n" + hai + "\n"
	return s
}

type M struct {
	s   []string
	l   []string
	g   []string
	qin []string
	qsb []bool
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
一个简单干支历
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
	zrWindos
}
type Edit struct {
	ed1, ed2, ed3, ed4 widget.Editor
	btn                widget.Clickable
	btn2               widget.Clickable
	btn3               widget.Clickable
	num                int
	btn4               widget.Clickable //关于
}
