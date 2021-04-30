package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"eliasnaur.com/font/roboto/robotoregular"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/unit"
	"github.com/gonoto/notosans"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"liangzi.local/cal/cal"
)

func init() {
}

type (
	D = layout.Dimensions
	C = layout.Context
)

var (
	listday = &layout.List{Axis: layout.Vertical}
	t       = time.Now().Local()
)

func main() {
	go func() {
		w := app.NewWindow(app.Size(unit.Dp(600), unit.Dp(800)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := mth()
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			render(gtx, th)
			e.Frame(gtx.Ops)
		}
	}
}

//////////////////

func render(gtx C, th *material.Theme) D {
	return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx C) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				l := material.H6(th, "干支历")
				//l.Font = text.Font{Weight: text.Bold}
				l.Font = text.Font{Typeface: "Noto"}
				l.Alignment = text.Middle
				return l.Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				return layout.Inset{Top: unit.Sp(20)}.Layout(gtx, func(gtx C) D {
					lunar := ymd(t)
					l := material.H6(th, lunar)
					l.Font = text.Font{Typeface: "Noto"}
					return l.Layout(gtx)
				})
			}),
			layout.Flexed(1, func(gtx C) D {
				return layout.UniformInset(unit.Sp(20)).Layout(gtx, func(gtx C) D {
					return listday.Layout(gtx, 1, func(gtx C, i int) D {
						l := material.Body1(th, yueli(t.Year(), int(t.Month()), t.Day()))
						l.Font = text.Font{Typeface: "Noto"}
						return l.Layout(gtx)
					})
				})
			}),
		)
	})
}

//返回一个中文字体主题
func mth() *material.Theme {
	fonts := gofont.Collection()
	fonts = appendTTF(fonts, text.Font{Typeface: "Roboto"}, robotoregular.TTF)
	fonts = appendOTC(fonts, text.Font{Typeface: "Noto"}, notosans.OTC())
	th := material.NewTheme(fonts)

	return th
}
func appendTTF(collection []text.FontFace, fnt text.Font, ttf []byte) []text.FontFace {
	face, err := opentype.Parse(ttf)
	if err != nil {
		panic(fmt.Errorf("failed to parse font: %v", err))
	}
	return append(collection, text.FontFace{Font: fnt, Face: face})
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
	//t := time.Now().Local()
	lunar := cal.NewLunarMonth(t.Year(), int(t.Month()), t.Day())
	var info string
	if lunar.LeapSday == 0 {
		info = fmt.Sprintf("%d年-%d月(%s)-%s %s",
			lunar.LY, lunar.LM, lunar.Ydx, lunar.LRmc, lunar.Week)
	} else {
		info = fmt.Sprintf("%d年-闰%d月(%s)-%s %s",
			lunar.LY, lunar.LeapM, lunar.LeapYdx, lunar.LeapRmc, lunar.Week)
	}
	gz := cal.NewGanZhiInfo(t.Year(), int(t.Month()), t.Day(), t.Hour())
	gzs := gz.YearGZ + "-" + gz.MonthGZ + "-" + gz.DayGZ + "-" + gz.HourGZ

	sw := cal.NewShuoWangTS(t.Year(), int(t.Month()), t.Day())
	shuo := "朔: " + sw.ShuoTS
	wang := "望: " + sw.WangTS
	shang := "上弦: " + sw.ShangXianTS
	xia := "下弦: " + sw.XiaXianTS
	sws := shuo + "\n" + wang + "\n" + shang + "\n" + xia
	s := info + "\n" + gzs + "\n" + sws
	return s
}

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
