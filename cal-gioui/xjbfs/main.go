package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gonoto/notosans"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"liangzi.local/cal/cal"
	"liangzi.local/xjbfs"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("干支历择日"),
			app.Size(unit.Dp(360), unit.Dp(640))) //360X640 HTC Ultra
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

var (
	listday                = &layout.List{Axis: layout.Vertical}
	ed1, ed2, ed3, ed4     widget.Editor
	btn1, btn2, btn3, btn4 widget.Clickable
	t                      = time.Now().Local()
	y                      = t.Year()
	m                      = int(t.Month())
	d                      = t.Day()
	h                      = t.Hour()
	ymds                   = ymd(t)
	show                   = Show()
)

func loop(w *app.Window) error {
	ticker := time.NewTicker(time.Second)
	th := mth()
	var ops op.Ops
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.FrameEvent:
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
		ymds = ymd(t)
		show = Show()
	}
	//上一小时
	if btn2.Clicked() {
		h -= 1
		if h < 1 {
			d -= 1
			h = 23
		}
		//fmt.Println("btn2-->", y, m, d, h)
		if ((y%4 == 0 && y%100 != 0) || y%400 == 0) && m == 3 {
			if d < 1 {
				m = 2
				d = 1
			}
			//	fmt.Println("btn2leap3-->", y, m, d, h)
		}
		if !((y%4 == 0 && y%100 != 0) || y%400 == 0) && m == 3 {
			if d < 1 {
				m = 2
				d = 1
			}
			//	fmt.Println("btn2nolea3-->", y, m, d, h)
		}
		if m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
			if d < 1 {
				m -= 1
				if m < 1 {
					m += 12
				}
				d = 30
			}
			//	fmt.Println("30-->", y, m, d, h)
		}
		if m == 1 || m == 2 || m == 4 || m == 6 || m == 9 || m == 11 {
			if d < 1 {
				m -= 1
				if m < 1 {
					m += 12
				}
				d = 31
			}
			//	fmt.Println("31-->", y, m, d, h)
		}
		if m > 12 {
			y -= 1
		}
		t = time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
		ymds = ymd(t)
		show = Show()
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
		//fmt.Println(">", y, m, d, h)
		t = time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
		ymds = ymd(t)
		show = Show()
	}
	if btn4.Clicked() {
		ymds = ymd(t)
		show = about
	}
	//----------------------------------------------------------
	inset := layout.UniformInset(unit.Dp(3)) //限制宽度
	inset.Top = unit.Dp(43)                  //距离上面输入小部件的间隔
	return inset.Layout(gtx, func(gtx C) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return layout.Inset{Top: unit.Sp(14)}.Layout(gtx, func(gtx C) D {
					lunar := ymd(t)
					l := material.H6(th, lunar)
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
						l := material.Body1(th, show)
						l.Font = text.Font{Typeface: "Noto"}
						return l.Layout(gtx)
					})
				})
			}),
		)
	})
}
func convEdit(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("convEdit:", err)
	}
	return i
}

//返回一个中文字体主题
func mth() *material.Theme {
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

////////
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
	//月将
	yj := jq.YueJiang(t)
	yjg := yj.Star
	yjm := "月将: " + yj.Name + " " + yjg
	s1 := gzs + "\n" + lunars + "\n" + todayJQs + "\n" + yjm

	return s1
}

//生肖太岁
func taiSui(ygz string) string {
	zhi := xjbfs.ZhiTaiSui(ygz)
	chong := xjbfs.ChongTaiSui(ygz)
	xing := xjbfs.XingTaiSui(ygz)
	hai := xjbfs.HaiTaiSui(ygz)
	s := zhi + " " + xing + "\n" + chong + " " + hai
	return s
}

func Show() string {
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
	//日建除
	x8 := fmt.Sprintf("\n☺日建除☺: %s\n", jcday)
	x9 := fmt.Sprintf("♠日黄黑♠: %s\n", hhd)
	x6 := fmt.Sprintf("♤日吉♤: %s\n", mjs)
	x7 := fmt.Sprintf("♠日凶♠: %s\n", mxs)
	//日表
	rjs, rxs := xjbf.GetRiBaoS(jz60) // xjbf.RiBiao(jz60)

	ri := x8 + x9 + x6 + x7
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

var about = `
协纪辩方书择日
朔 望 上弦 下弦 
月将 节气 十二星宫
金符经七煞
GUI源码: https://github.com/Aquarian-Age/ccal
mail: liangzi@yandex.com
本人保留版权!
`

func yueli(y, m, d int) string {

	yl := cal.NewYueLi(y, m, d)
	l := yl.LunarD
	s := yl.SolarD
	gz := yl.GzD
	var s1, l1, g1, s2, l2, g2, s3, l3, g3, s4, l4, g4, s5, l5, g5 []string
	for i := 0; i < 6; i++ {
		s1 = append(s1, s[i])
		l1 = append(l1, l[i])
		g1 = append(g1, gz[i])
	}
	for i := 6; i < 12; i++ {
		s2 = append(s2, s[i])
		l2 = append(l2, l[i])
		g2 = append(g2, gz[i])
	}
	for i := 12; i < 18; i++ {
		s3 = append(s3, s[i])
		l3 = append(l3, l[i])
		g3 = append(g3, gz[i])
	}
	for i := 18; i < 24; i++ {
		s4 = append(s4, s[i])
		l4 = append(l4, l[i])
		g4 = append(g4, gz[i])
	}
	for i := 24; i < len(l); i++ {
		s5 = append(s5, s[i])
		l5 = append(l5, l[i])
		g5 = append(g5, gz[i])

	}

	rs1 := strings.Join(s1, " | ")
	rl1 := strings.Join(l1, " | ")
	rg1 := strings.Join(g1, " | ")
	rs2 := strings.Join(s2, " | ")
	rl2 := strings.Join(l2, " | ")
	rg2 := strings.Join(g2, " | ")
	rs3 := strings.Join(s3, " | ")
	rl3 := strings.Join(l3, " | ")
	rg3 := strings.Join(g3, " | ")
	rs4 := strings.Join(s4, " | ")
	rl4 := strings.Join(l4, " | ")
	rg4 := strings.Join(g4, " | ")
	rs5 := strings.Join(s5, " | ")
	rl5 := strings.Join(l5, " | ")
	rg5 := strings.Join(g5, " | ")

	rs := "\n" + "---------------------------------------------------------" + "\n" +
		rs1 + "\n" + rl1 + "\n" + rg1 +
		"\n" + "---------------------------------------------------------" + "\n" +
		rs2 + "\n" + rl2 + "\n" + rg2 +
		"\n" + "---------------------------------------------------------" + "\n" +
		rs3 + "\n" + rl3 + "\n" + rg3 +
		"\n" + "---------------------------------------------------------" + "\n" +
		rs4 + "\n" + rl4 + "\n" + rg4 +
		"\n" + "---------------------------------------------------------" + "\n" +
		rs5 + "\n" + rl5 + "\n" + rg5

	return rs + "\n\n" + "UI地址: https://github.com/Aquarian-Age/ccal"
}
