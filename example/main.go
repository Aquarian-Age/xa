package main

import (
	"flag"
	"fmt"

	"github.com/Aquarian-Age/xa/pkg/gz"
	"github.com/starainrt/astro/basic"
	"github.com/starainrt/astro/calendar"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"

	"math"
	"os"
	"sort"
	"time"
	_ "time/tzdata"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TForm1 struct {
	*vcl.TForm
	title0, title1, title2, title3, title4, title5, title6 *vcl.TLabel
	//
	label1, label2, label3, label4, label5, label6, label7        *vcl.TLabel
	label8, label9, label10, label11, label12, label13, label14   *vcl.TLabel
	label15, label16, label17, label18, label19, labe20, label21  *vcl.TLabel
	label22, label23, label24, label25, label26, label27, label28 *vcl.TLabel
	label29, label30, label31, label32, label33, label34, label35 *vcl.TLabel
	label36, label37, label38, label39, label40, label41, label42 *vcl.TLabel
	//
	gzLabel1, gzLabel2, gzLabel3, gzLabel4, gzLabel5, gzLabel6, gzLabel7        *vcl.TLabel
	gzLabel8, gzLabel9, gzLabel10, gzLabel11, gzLabel12, gzLabel13, gzLabel14   *vcl.TLabel
	gzLabel15, gzLabel16, gzLabel17, gzLabel18, gzLabel19, gzlabe20, gzLabel21  *vcl.TLabel
	gzLabel22, gzLabel23, gzLabel24, gzLabel25, gzLabel26, gzLabel27, gzLabel28 *vcl.TLabel
	gzLabel29, gzLabel30, gzLabel31, gzLabel32, gzLabel33, gzLabel34, gzLabel35 *vcl.TLabel
	gzLabel36, gzLabel37, gzLabel38, gzLabel39, gzLabel40, gzLabel41, gzLabel42 *vcl.TLabel
	//
	moonLabel1, moonLabel2, moonLabel3, moonLabel4, moonLabel5, moonLabel6, moonLabel7        *vcl.TLabel
	moonLabel8, moonLabel9, moonLabel10, moonLabel11, moonLabel12, moonLabel13, moonLabel14   *vcl.TLabel
	moonLabel15, moonLabel16, moonLabel17, moonLabel18, moonLabel19, moonLabe20, moonLabel21  *vcl.TLabel
	moonLabel22, moonLabel23, moonLabel24, moonLabel25, moonLabel26, moonLabel27, moonLabel28 *vcl.TLabel
	moonLabel29, moonLabel30, moonLabel31, moonLabel32, moonLabel33, moonLabel34, moonLabel35 *vcl.TLabel
	moonLabel36, moonLabel37, moonLabel38, moonLabel39, moonLabel40, moonLabel41, moonLabel42 *vcl.TLabel

	btna, btnb, btnc             *vcl.TButton
	labels, gzLabels, moonLabels []*vcl.TLabel
	pubLabel                     *vcl.TLabel
	sheng, shi                   *vcl.TEdit
}

var (
	form1        *TForm1
	printVersion bool
	T            = time.Now().Local()
	monthNow     int
	weekname     = []string{"日", "一", "二", "三", "四", "五", "六"}
	moonMonthMap = map[int]string{1: "正月", 2: "二月", 3: "三月", 4: "四月", 5: "五月", 6: "六月",
		7: "七月", 8: "八月", 9: "九月", 10: "十月", 11: "冬月", 12: "廿月"}
	dayMap = map[int]string{1: "初一", 2: "初二", 3: "初三", 4: "初四", 5: "初五", 6: "初六", 7: "初七",
		8: "初八", 9: "初九", 10: "初十", 11: "十一", 12: "十二", 13: "十三", 14: "十四",
		15: "十五", 16: "十六", 17: "十七", 18: "十八", 19: "十九", 20: "廿十", 21: "廿一",
		22: "廿二", 23: "廿三", 24: "廿四", 25: "廿五", 26: "廿六", 27: "廿七", 28: "廿八",
		29: "廿九", 30: "三十"}
	Lefts    = []int32{10, 50, 90, 130, 170, 210, 250}
	dayArray = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
		15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
		29, 30, 31}
	top0 = int32(10)
	top  = int32(40)
	top1 = top + 80
	top2 = top1 + 80
	top3 = top2 + 80
	top4 = top3 + 80
	top5 = top4 + 80
	year int
)
var (
	Version   = ""
	GoVersion = ""
	Mail      = ""
)

func PrintVersion() {
	fmt.Printf("App: %s\n", Version)
	fmt.Printf("%s\n", GoVersion)
	fmt.Printf("Mail: %s\n", Mail)
	os.Exit(0)
}

func main() {
	flag.BoolVar(&printVersion, "version", false, "print program build version")
	flag.Parse()
	if printVersion {
		PrintVersion()
	}

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&form1)
	vcl.Application.Run()
}
func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.SetWidth(500)
	f.SetHeight(550)
	f.ScreenCenter()
	f.SetCaption("简易农历")
	f.SetBorderStyle(types.BsSingle) //固定窗口大小
	f.SetShowHint(true)
	f.Title()
	f.initLabels(sender)
	f.initedit(sender)

	f.showLabels(T)
	f.OnBtnClicked(sender)
	f.about()
}
func (f *TForm1) about() {
	link := vcl.NewLinkLabel(f)
	link.SetParent(f)
	link.SetLeft(f.Width() - 40)
	link.SetTop(f.Height() - 30)
	link.SetCaption("<a href=\"https://github.com/Aquarian-Age/ccal/releases\">下载</a>")
	link.SetOnLinkClick(func(sender vcl.IObject, link string, linktype types.TSysLinkType) {
		rtl.SysOpen(link)
	})
}

func (f *TForm1) Title() {
	titles := []*vcl.TLabel{f.title0, f.title1, f.title2, f.title3, f.title4, f.title5, f.title6}
	for i := 0; i < len(titles); i++ {
		titles[i] = vcl.NewLabel(f)
		titles[i].SetParent(f)
		titles[i].SetLeft(Lefts[i])
		titles[i].SetTop(top0)
		titles[i].SetCaption(weekname[i])
	}
}
func (f *TForm1) initLabels(sender vcl.IObject) {
	labels := []*vcl.TLabel{
		f.label1, f.label2, f.label3, f.label4, f.label5, f.label6, f.label7,
		f.label8, f.label9, f.label10, f.label11, f.label12, f.label13, f.label14,
		f.label15, f.label16, f.label17, f.label18, f.label19, f.labe20, f.label21,
		f.label22, f.label23, f.label24, f.label25, f.label26, f.label27, f.label28,
		f.label29, f.label30, f.label31, f.label32, f.label33, f.label34, f.label35,
		f.label36, f.label37, f.label38, f.label39, f.label40, f.label41, f.label42,
	}
	gzLabels := []*vcl.TLabel{
		f.gzLabel1, f.gzLabel2, f.gzLabel3, f.gzLabel4, f.gzLabel5, f.gzLabel6, f.gzLabel7,
		f.gzLabel8, f.gzLabel9, f.gzLabel10, f.gzLabel11, f.gzLabel12, f.gzLabel13, f.gzLabel14,
		f.gzLabel15, f.gzLabel16, f.gzLabel17, f.gzLabel18, f.gzLabel19, f.labe20, f.gzLabel21,
		f.gzLabel22, f.gzLabel23, f.gzLabel24, f.gzLabel25, f.gzLabel26, f.gzLabel27, f.gzLabel28,
		f.gzLabel29, f.gzLabel30, f.gzLabel31, f.gzLabel32, f.gzLabel33, f.gzLabel34, f.gzLabel35,
		f.gzLabel36, f.gzLabel37, f.gzLabel38, f.gzLabel39, f.gzLabel40, f.gzLabel41, f.gzLabel42,
	}
	moonLabels := []*vcl.TLabel{
		f.moonLabel1, f.moonLabel2, f.moonLabel3, f.moonLabel4, f.moonLabel5, f.moonLabel6, f.moonLabel7,
		f.moonLabel8, f.moonLabel9, f.moonLabel10, f.moonLabel11, f.moonLabel12, f.moonLabel13, f.moonLabel14,
		f.moonLabel15, f.moonLabel16, f.moonLabel17, f.moonLabel18, f.moonLabel19, f.labe20, f.moonLabel21,
		f.moonLabel22, f.moonLabel23, f.moonLabel24, f.moonLabel25, f.moonLabel26, f.moonLabel27, f.moonLabel28,
		f.moonLabel29, f.moonLabel30, f.moonLabel31, f.moonLabel32, f.moonLabel33, f.moonLabel34, f.moonLabel35,
		f.moonLabel36, f.moonLabel37, f.moonLabel38, f.moonLabel39, f.moonLabel40, f.moonLabel41, f.moonLabel42,
	}
	for i := 0; i < len(labels); i++ {
		labels[i] = vcl.NewLabel(f)
		labels[i].SetParent(f)
		//labels[i].SetAlignment(types.AsrCenter)

		gzLabels[i] = vcl.NewLabel(f)
		gzLabels[i].SetParent(f)
		gzLabels[i].SetColor(colors.ClWhite)

		moonLabels[i] = vcl.NewLabel(f)
		moonLabels[i].SetParent(f)
		//moonLabels[i].SetAlign(types.AlNone)
		moonLabels[i].SetColor(colors.ClWhite)
	}
	f.labels = labels
	f.gzLabels = gzLabels
	f.moonLabels = moonLabels

	f.pubLabel = vcl.NewLabel(f)
	f.pubLabel.SetParent(f)
}
func (f *TForm1) initedit(sender vcl.IObject) {
	f.sheng = vcl.NewEdit(f)
	f.sheng.SetParent(f)
	f.sheng.SetLeft(f.Width() - 200)
	f.sheng.SetTop(f.Height() - 35)
	f.sheng.SetWidth(70)
	f.sheng.SetHint("输入省名称自动计算省经纬度")

	f.shi = vcl.NewEdit(f)
	f.shi.SetParent(f)
	f.shi.SetLeft(f.Width() - 120)
	f.shi.SetTop(f.Height() - 35)
	f.shi.SetWidth(70)
	f.shi.SetHint("输入城市名称自动计算经纬度")

}
func (f *TForm1) OnBtnClicked(sender vcl.IObject) {
	f.btna = vcl.NewButton(f)
	f.btna.SetParent(f)
	f.btna.SetLeft(f.Width() - 30)
	f.btna.SetWidth(20)
	f.btna.SetTop(10)
	f.btna.SetCaption("+")
	f.btna.SetOnClick(f.btnaClick)

	f.btnb = vcl.NewButton(f)
	f.btnb.SetParent(f)
	f.btnb.SetLeft(f.btna.Left() - f.btna.Width())
	f.btnb.SetWidth(20)
	f.btnb.SetTop(f.btna.Top())
	f.btnb.SetCaption("-")
	f.btnb.SetOnClick(f.btnbClick)

	f.btnc = vcl.NewButton(f)
	f.btnc.SetParent(f)
	f.btnc.SetLeft(f.btnb.Left() - f.btnb.Width())
	f.btnc.SetTop(f.btna.Top())
	f.btnc.SetWidth(20)
	f.btnc.SetCaption("今")
	f.btnc.SetOnClick(f.btncClick)
}

func (f *TForm1) btnaClick(sender vcl.IObject) {
	monthNow += 1
	if monthNow > 12 {
		monthNow = 1
		year += 1
	}
	t := time.Date(year, time.Month(monthNow), 1, 0, 0, 0, 0, time.Local)
	f.showLabels(t)
	//f.showBtnClicked(sender)
	years := fmt.Sprintf("%d年%d月%d日\n", year, monthNow, 1)
	f.pubLabel.Refresh()
	f.pubLabel.SetLeft(290)
	f.pubLabel.SetTop(10)
	f.pubLabel.SetCaption(years)
}

func (f *TForm1) btnbClick(sender vcl.IObject) {
	monthNow -= 1
	if monthNow <= 0 {
		monthNow = 12
		year -= 1
	}
	t := time.Date(year, time.Month(monthNow), 1, 0, 0, 0, 0, time.Local)
	f.showLabels(t)
	//f.showBtnClicked(sender)
	years := fmt.Sprintf("%d年%d月%d日\n", year, monthNow, 1)
	f.pubLabel.Refresh()
	f.pubLabel.SetLeft(290)
	f.pubLabel.SetTop(10)
	f.pubLabel.SetCaption(years)
}

func (f *TForm1) btncClick(sender vcl.IObject) {
	//t := time.Date(year, time.Month(monthNow), 1, 0, 0, 0, 0, time.Local)
	//t := time.Now().Local()
	f.showLabels(T)
	years := fmt.Sprintf("%d年%d月%d日\n", year, monthNow, T.Day())
	f.pubLabel.Refresh()
	f.pubLabel.SetLeft(290)
	f.pubLabel.SetTop(10)
	f.pubLabel.SetCaption(years)
}

func (f *TForm1) onLabelxClick(sender vcl.IObject) {
	x := vcl.AsLabel(sender)
	tl := time.Date(year, time.Month(monthNow), x.Tag(), 0, 0, 0, 0, time.Local)
	//fmt.Println(tl.String()[:19])

	years := fmt.Sprintf("%d年%d月%d日\n", year, monthNow, x.Tag())
	gzo := gz.NewGanZhi(year, monthNow, x.Tag(), 0)
	ganZhis := fmt.Sprintf("%s年 %s月 %s日\n", gzo.YGZ, gzo.MGZ, gzo.DGZ)
	nayins := gz.NaYin(gzo.YGZ) + " " + gz.NaYin(gzo.MGZ) + " " + gz.NaYin(gzo.DGZ) + "\n"
	yjo := gzo.YueJiangStruct()
	zhongqis := fmt.Sprintf("中气: %s\n中气时间: %s\n", yjo.ZhongQiName, yjo.ZhongQiT)
	yuejiangs := fmt.Sprintf("月将: %s(%s)\n", yjo.Zhi, yjo.Name)
	jianchus := fmt.Sprintf("建除: %s\n", gzo.JianChuDay())
	huangheis := fmt.Sprintf("黄黑: %s\n", gzo.RiHuangHei1())
	riqins := fmt.Sprintf("日禽: %s\n", gzo.RiQin(int(tl.Weekday())))
	lunar, moon := gzo.GetLunar()
	moons := fmt.Sprintf("%s\n%s\n", lunar, moon)
	s := years + ganZhis + nayins + zhongqis + yuejiangs + jianchus + huangheis + riqins + moons

	jing := f.sheng.Text()
	wei := f.shi.Text()
	jws := jingweix(jing, wei)
	s += jws
	f.pubLabel.Refresh()
	f.pubLabel.SetLeft(290)
	f.pubLabel.SetTop(10)
	f.pubLabel.Font().SetColor(colors.ClSkyblue)
	f.pubLabel.Font().SetSize(12)
	f.pubLabel.SetCaption(s)
}
func jingweix(jing, wei string) string {
	j, w := GetJingWei(jing, wei)
	so := GetSun(j, w)
	s1 := `日出: ` + so.ShengQi
	s2 := `日落: ` + so.LuoXia
	s3 := `中天: ` + so.ZhongTian
	s4 := `黄赤交角: ` + so.HuangChi
	s5 := `太阳真黄径: ` + so.HuangJing
	s6 := `日地距离: ` + so.JuLi
	s := fmt.Sprintf("%s %s\n经度:%f\n纬度:%f\n%s\n%s\n%s\n%s\n%s\n%s\n", jing, wei, j, w, s1, s2, s3, s4, s5, s6)
	return s
}

// func ShuoWang(gzo *gz.GanZhi) string {
// 	sw := gzo.Moons()
// 	lastWang := sw.FormatTime(sw.LastWang)
// 	shuo := sw.FormatTime(sw.Shuo)
// 	wang := sw.FormatTime(sw.Wang)
// 	shangXian := sw.FormatTime(sw.ShangXian)
// 	xiaXian := sw.FormatTime(sw.XiaXian)
// 	nextWang := sw.FormatTime(sw.NextWang)
// 	s := fmt.Sprintf("上一望月: %s\n最近朔月: %s\n最近望月: %s\n最近上弦: %s\n最近下弦: %s\n下一望月: %s\n",
// 		lastWang, shuo, wang, shangXian, xiaXian, nextWang)
// 	return s
// }

//func (f *TForm1) showBtnClicked(sender vcl.IObject) {
//	years := fmt.Sprintf("%d年%d月%d日\n", year, monthNow, 1)
//	f.pubLabel.Refresh()
//	f.pubLabel.SetLeft(290)
//	f.pubLabel.SetTop(10)
//	f.pubLabel.SetCaption(years)
//}

func (f *TForm1) showLabels(t time.Time) {
	monthNow = int(t.Month())
	lastmonth := monthNow - 1
	if lastmonth == 0 {
		lastmonth = 12
	}
	//判断闰月
	y := t.Year()
	year = y
	b := (y%4 == 0 && y%100 != 0) || y%400 == 0
	var allDay int
	switch monthNow {
	case 1, 3, 5, 7, 8, 10, 12:
		allDay = 31
	case 4, 6, 9, 11:
		allDay = 30
	case 2:
		if b == true {
			allDay = 29
		} else {
			allDay = 28
		}
	}

	t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	weeknow := int(t.Weekday())
	//fmt.Println(t.String()[:20])

	jqt(t)

	var xday []int32
	wn := int32(weeknow)
	_days := dayArray[:wn]
	xday = append(xday, _days...)
	for i := 0; i < allDay; i++ {
		xday = append(xday, int32(i+1))
	}

	labels := f.labels
	gzLabels := f.gzLabels
	moonLabels := f.moonLabels
	for i := 0; i < len(labels); i++ {
		f.labels[i].SetCaption(``)
		f.gzLabels[i].SetCaption(``)
		f.moonLabels[i].SetCaption(``)
		//
		if i >= 0 && i <= 6 {
			if i >= weeknow {
				tx := time.Date(year, time.Month(monthNow), int(xday[i]), 0, 0, 0, 0, time.Local)
				labels[i].SetLeft(Lefts[i])
				labels[i].SetTop(top)
				labels[i].SetWidth(10)
				if s, ok := jqtmap[tx]; ok {
					xs := fmt.Sprintf("%d%s", xday[i], s)
					labels[i].Font().SetSize(10)
					labels[i].Font().SetColor(colors.ClSkyblue)
					labels[i].SetCaption(xs)
				} else {
					labels[i].Font().SetSize(12)
					labels[i].Font().SetColor(colors.ClBlack)
					labels[i].SetCaption(fmt.Sprintf("  %d", xday[i]))
				}
				//labels[i].SetCaption(fmt.Sprintf("  %d", xday[i]))
				labels[i].SetTag(int(xday[i]))
				labels[i].SetOnClick(f.onLabelxClick)

				dgz := dayGZ(year, monthNow, int(xday[i]))
				gzLabels[i].SetLeft(Lefts[i])
				gzLabels[i].SetWidth(10)
				gzLabels[i].SetTop(labels[i].Top() + 25)
				gzLabels[i].SetCaption(dgz)

				//tx := time.Date(year, time.Month(monthNow), int(xday[i]), 0, 0, 0, 0, time.Local)
				_, moonday, _, _ := calendar.ChineseLunar(tx)
				moon := dayMap[moonday]
				moonLabels[i].SetLeft(Lefts[i])
				moonLabels[i].SetWidth(10)
				moonLabels[i].SetTop(gzLabels[i].Top() + 25)
				moonLabels[i].SetCaption(moon)
			}

		}
		if i >= 7 && i <= 13 {
			tx := time.Date(year, time.Month(monthNow), int(xday[i]), 0, 0, 0, 0, time.Local)
			labels[i].SetLeft(Lefts[i-7])
			labels[i].SetTop(top1)
			labels[i].SetWidth(10)
			labels[i].SetTag(int(xday[i]))
			if s, ok := jqtmap[tx]; ok {
				xs := fmt.Sprintf("%d%s", xday[i], s)
				labels[i].Font().SetSize(10)
				labels[i].Font().SetColor(colors.ClSkyblue)
				labels[i].SetCaption(xs)
			} else {
				labels[i].Font().SetSize(12)
				labels[i].Font().SetColor(colors.ClBlack)
				labels[i].SetCaption(fmt.Sprintf("  %d", xday[i]))
			}
			labels[i].SetOnClick(f.onLabelxClick)

			dgz := dayGZ(year, monthNow, int(xday[i]))
			gzLabels[i].SetLeft(Lefts[i-7])
			gzLabels[i].SetWidth(10)
			gzLabels[i].SetTop(labels[i].Top() + 25)
			gzLabels[i].SetCaption(dgz)

			_, moonday, _, _ := calendar.ChineseLunar(tx)
			moon := dayMap[moonday]
			moonLabels[i].SetLeft(Lefts[i-7])
			moonLabels[i].SetWidth(10)
			moonLabels[i].SetTop(gzLabels[i].Top() + 25)
			moonLabels[i].SetCaption(moon)
		}
		if i >= 14 && i <= 20 {
			tx := time.Date(year, time.Month(monthNow), int(xday[i]), 0, 0, 0, 0, time.Local)
			labels[i].SetLeft(Lefts[i-14])
			labels[i].SetTop(top2)
			labels[i].SetWidth(10)
			if s, ok := jqtmap[tx]; ok {
				xs := fmt.Sprintf("%d%s", xday[i], s)
				labels[i].Font().SetSize(10)
				labels[i].Font().SetColor(colors.ClSkyblue)
				labels[i].SetCaption(xs)
			} else {
				labels[i].Font().SetSize(12)
				labels[i].Font().SetColor(colors.ClBlack)
				labels[i].SetCaption(fmt.Sprintf("  %d", xday[i]))
			}
			labels[i].SetTag(int(xday[i]))
			labels[i].SetOnClick(f.onLabelxClick)

			dgz := dayGZ(year, monthNow, int(xday[i]))
			gzLabels[i].SetLeft(Lefts[i-14])
			gzLabels[i].SetWidth(10)
			gzLabels[i].SetTop(labels[i].Top() + 25)
			gzLabels[i].SetCaption(dgz)

			_, moonday, _, _ := calendar.ChineseLunar(tx)
			moon := dayMap[moonday]
			moonLabels[i].SetLeft(Lefts[i-14])
			moonLabels[i].SetWidth(10)
			moonLabels[i].SetTop(gzLabels[i].Top() + 25)
			moonLabels[i].SetCaption(moon)
		}
		if i >= 21 && i <= 27 {
			tx := time.Date(year, time.Month(monthNow), int(xday[i]), 0, 0, 0, 0, time.Local)
			labels[i].SetLeft(Lefts[i-21])
			labels[i].SetTop(top3)
			labels[i].SetWidth(10)
			if s, ok := jqtmap[tx]; ok {
				xs := fmt.Sprintf("%d%s", xday[i], s)
				labels[i].Font().SetSize(10)
				labels[i].Font().SetColor(colors.ClSkyblue)
				labels[i].SetCaption(xs)
			} else {
				labels[i].Font().SetSize(12)
				labels[i].Font().SetColor(colors.ClBlack)
				labels[i].SetCaption(fmt.Sprintf("  %d", xday[i]))
			}
			labels[i].SetTag(int(xday[i]))
			labels[i].SetOnClick(f.onLabelxClick)

			dgz := dayGZ(year, monthNow, int(xday[i]))
			gzLabels[i].SetLeft(Lefts[i-21])
			gzLabels[i].SetWidth(10)
			gzLabels[i].SetTop(labels[i].Top() + 25)
			gzLabels[i].SetCaption(dgz)

			_, moonday, _, _ := calendar.ChineseLunar(tx)
			moon := dayMap[moonday]
			moonLabels[i].SetLeft(Lefts[i-21])
			moonLabels[i].SetWidth(10)
			moonLabels[i].SetTop(gzLabels[i].Top() + 25)
			moonLabels[i].SetCaption(moon)
		}
		if i >= 28 && i <= 34 {
			if len(xday) > 34 {
				tx := time.Date(year, time.Month(monthNow), int(xday[i]), 0, 0, 0, 0, time.Local)
				labels[i].SetLeft(Lefts[i-28])
				labels[i].SetTop(top4)
				labels[i].SetWidth(10)
				if s, ok := jqtmap[tx]; ok {
					xs := fmt.Sprintf("%d%s", xday[i], s)
					labels[i].Font().SetSize(10)
					labels[i].Font().SetColor(colors.ClSkyblue)
					labels[i].SetCaption(xs)
				} else {
					labels[i].Font().SetSize(12)
					labels[i].Font().SetColor(colors.ClBlack)
					labels[i].SetCaption(fmt.Sprintf("  %d", xday[i]))
				}
				labels[i].SetTag(int(xday[i]))
				labels[i].SetOnClick(f.onLabelxClick)

				dgz := dayGZ(year, monthNow, int(xday[i]))
				gzLabels[i].SetLeft(Lefts[i-28])
				gzLabels[i].SetWidth(10)
				gzLabels[i].SetTop(labels[i].Top() + 25)
				gzLabels[i].SetCaption(dgz)

				_, moonday, _, _ := calendar.ChineseLunar(tx)
				moon := dayMap[moonday]
				moonLabels[i].SetLeft(Lefts[i-28])
				moonLabels[i].SetWidth(10)
				moonLabels[i].SetTop(gzLabels[i].Top() + 25)
				moonLabels[i].SetCaption(moon)
			} else {
				if i < len(xday) {
					tx := time.Date(year, time.Month(monthNow), int(xday[i]), 0, 0, 0, 0, time.Local)
					labels[i].SetLeft(Lefts[i-28])
					labels[i].SetTop(top4)
					labels[i].SetWidth(10)
					if s, ok := jqtmap[tx]; ok {
						xs := fmt.Sprintf("%d%s", xday[i], s)
						labels[i].Font().SetSize(10)
						labels[i].Font().SetColor(colors.ClSkyblue)
						labels[i].SetCaption(xs)
					} else {
						labels[i].Font().SetSize(12)
						labels[i].Font().SetColor(colors.ClBlack)
						labels[i].SetCaption(fmt.Sprintf("  %d", xday[i]))
					}
					labels[i].SetTag(int(xday[i]))
					labels[i].SetOnClick(f.onLabelxClick)

					dgz := dayGZ(year, monthNow, int(xday[i]))
					gzLabels[i].SetLeft(Lefts[i-28])
					gzLabels[i].SetWidth(10)
					gzLabels[i].SetTop(labels[i].Top() + 25)
					gzLabels[i].SetCaption(dgz)

					_, moonday, _, _ := calendar.ChineseLunar(tx)
					moon := dayMap[moonday]
					moonLabels[i].SetLeft(Lefts[i-28])
					moonLabels[i].SetWidth(10)
					moonLabels[i].SetTop(gzLabels[i].Top() + 25)
					moonLabels[i].SetCaption(moon)
				}
			}
		}
		if i >= 35 && len(xday)-1 <= 41 && i <= 41 {
			if i < len(xday) {
				tx := time.Date(year, time.Month(monthNow), int(xday[i]), 0, 0, 0, 0, time.Local)
				labels[i].SetLeft(Lefts[i-35])
				labels[i].SetTop(top5)
				labels[i].SetWidth(10)
				if s, ok := jqtmap[tx]; ok {
					xs := fmt.Sprintf("%d%s", xday[i], s)
					labels[i].Font().SetSize(10)
					labels[i].Font().SetColor(colors.ClSkyblue)
					labels[i].SetCaption(xs)
				} else {
					labels[i].Font().SetSize(12)
					labels[i].Font().SetColor(colors.ClBlack)
					labels[i].SetCaption(fmt.Sprintf("  %d", xday[i]))
				}
				labels[i].SetTag(int(xday[i]))
				labels[i].SetOnClick(f.onLabelxClick)

				dgz := dayGZ(year, monthNow, int(xday[i]))
				gzLabels[i].SetLeft(Lefts[i-35])
				gzLabels[i].SetWidth(10)
				gzLabels[i].SetTop(labels[i].Top() + 25)
				gzLabels[i].SetCaption(dgz)

				_, moonday, _, _ := calendar.ChineseLunar(tx)
				moon := dayMap[moonday]
				moonLabels[i].SetLeft(Lefts[i-35])
				moonLabels[i].SetWidth(10)
				moonLabels[i].SetTop(gzLabels[i].Top() + 25)
				moonLabels[i].SetCaption(moon)
			}
		}
	}
}

func dayGZ(year, month, day int) string {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	jd := calendar.Date2JDE(t)
	jdi := int(math.Ceil(jd))
	return dGz(jdi)
}
func dGz(jdI int) string {
	gn := 1 + (jdI%60-1)%10 //干
	if gn == 0 {
		gn += 10
	}
	z := 1 + +(jdI%60+1)%12 //支

	//g 日干数字
	daygM := Gans[gn-1]
	dayzM := Zhi[z-1]

	dgz := daygM + dayzM
	return dgz
}

var (
	Gans    = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	Zhi     = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	jqnames = []string{
		"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
		"春分", "清明", "谷雨", "立夏", "小满", "芒种",
		"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
		"秋分", "寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
	}
)
var jqtmap = make(map[time.Time]string)

//k:1-->上一年冬至时间 k:25-->本年冬至时间 k:4--本年立春
func jqt(t time.Time) {
	jqmap := basic.GetOneYearJQ(t.Year() - 1)
	var keys []int
	for k := range jqmap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var jqts []time.Time
	for _, v := range keys {
		jqtx := calendar.JDE2Date(jqmap[v])
		jqtx = time.Date(jqtx.Year(), jqtx.Month(), jqtx.Day(), 0, 0, 0, 0, time.Local)
		jqts = append(jqts, jqtx)
	}

	for i := 0; i < len(jqts); i++ {
		for j := 0; j < len(keys); j++ {
			if i == j {
				jqtmap[jqts[j]] = jqnames[keys[j]-1]
				break
			}
		}
	}
}
