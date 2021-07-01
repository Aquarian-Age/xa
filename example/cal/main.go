package main

import (
	"fmt"
	"github.com/Aquarian-Age/xa/pkg/gz"
	"github.com/starainrt/astro/basic"
	"github.com/ying32/govcl/vcl"
	"strconv"
	"time"
)

type TMainForm struct {
	*vcl.TForm
	labely                                                                                *vcl.TLabel
	input                                                                                 *vcl.TEdit
	BtnOK                                                                                 *vcl.TButton
	BtnM1, BtnM2, BtnM3, BtnM4, BtnM5, BtnM6, BtnM7, BtnM8, BtnM9, BtnM10, BtnM11, BtnM12 *vcl.TButton //阳历　1-12月

	label, label2, label3, label4, label5, label6, label7         *vcl.TLabel
	label8, label9, label10, label11, label12, label13, label14   *vcl.TLabel
	label15, label16, label17, label18, label19, label20, label21 *vcl.TLabel
	label22, label23, label24, label25, label26, label27, label28 *vcl.TLabel
	label29, label30, label31                                     *vcl.TLabel
	//
	Ltoday *vcl.TLabel
	about  *vcl.TButton
}

var (
	mainForm *TMainForm
)

func main() {
	vcl.RunApp(&mainForm)
}

const (
	topm      = 10
	leftLabel = 10
	leftInput = 110
	leftok    = 80 + iota*45
	M1
	M2
	M3
	M4
	M5
	M6
	M7
	M8
	M9
	M10
	M11
	M12

	leftAbout  = 500
	leftLtoday = 570
	wm         = 40
	hm         = 30
)

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetWidth(800)
	f.SetHeight(600)
	f.SetCaption("二十八宿日历")
	f.Font().SetSize(12)
	//
	f.label = vcl.NewLabel(f)
	f.label.SetParent(f)
	f.label.SetLeft(leftLabel)
	f.label.SetTop(topm)
	f.label.SetWidth(wm)
	f.label.SetHeight(hm)
	f.label.SetTextBuf("输入年月日时\n2021010203")
	//
	f.input = vcl.NewEdit(f)
	f.input.SetParent(f)
	f.input.SetTop(topm)
	f.input.SetLeft(leftInput)
	f.input.SetWidth(100)
	f.input.SetHeight(30)
	f.input.SetTextBuf("2021010203")
	//
	f.BtnOK = vcl.NewButton(f)
	f.BtnOK.SetParent(f)
	f.BtnOK.SetTop(topm)
	f.BtnOK.SetLeft(leftok)
	f.BtnOK.SetWidth(wm)
	f.BtnOK.SetHeight(hm)
	f.BtnOK.SetCaption("OK")
	f.BtnOK.SetOnClick(f.OnBtnOK)
	//
	f.BtnM1 = vcl.NewButton(f)
	f.BtnM1.SetParent(f)
	f.BtnM1.SetTop(topm)
	f.BtnM1.SetLeft(M1)
	f.BtnM1.SetWidth(wm)
	f.BtnM1.SetHeight(hm)
	f.BtnM1.SetCaption("1月")
	f.BtnM1.SetOnClick(f.OnBtnM1)
	//
	f.BtnM2 = vcl.NewButton(f)
	f.BtnM2.SetParent(f)
	f.BtnM2.SetTop(topm)
	f.BtnM2.SetLeft(M2)
	f.BtnM2.SetWidth(wm)
	f.BtnM2.SetHeight(hm)
	f.BtnM2.SetCaption("2月")
	f.BtnM2.SetOnClick(f.OnBtnM2)
	//
	f.BtnM3 = vcl.NewButton(f)
	f.BtnM3.SetParent(f)
	f.BtnM3.SetTop(topm)
	f.BtnM3.SetLeft(M3)
	f.BtnM3.SetWidth(wm)
	f.BtnM3.SetHeight(hm)
	f.BtnM3.SetCaption("3月")
	f.BtnM3.SetOnClick(f.OnBtnM3)
	//
	f.BtnM4 = vcl.NewButton(f)
	f.BtnM4.SetParent(f)
	f.BtnM4.SetTop(topm)
	f.BtnM4.SetLeft(M4)
	f.BtnM4.SetWidth(wm)
	f.BtnM4.SetHeight(hm)
	f.BtnM4.SetCaption("4月")
	f.BtnM4.SetOnClick(f.OnBtnM4)
	//
	f.BtnM5 = vcl.NewButton(f)
	f.BtnM5.SetParent(f)
	f.BtnM5.SetTop(topm)
	f.BtnM5.SetLeft(M5)
	f.BtnM5.SetWidth(wm)
	f.BtnM5.SetHeight(hm)
	f.BtnM5.SetCaption("5月")
	f.BtnM5.SetOnClick(f.OnBtnM5)
	//
	f.BtnM6 = vcl.NewButton(f)
	f.BtnM6.SetParent(f)
	f.BtnM6.SetTop(topm)
	f.BtnM6.SetLeft(M6)
	f.BtnM6.SetWidth(wm)
	f.BtnM6.SetHeight(hm)
	f.BtnM6.SetCaption("6月")
	f.BtnM6.SetOnClick(f.OnBtnM6)
	//
	f.BtnM7 = vcl.NewButton(f)
	f.BtnM7.SetParent(f)
	f.BtnM7.SetTop(topm)
	f.BtnM7.SetLeft(M7)
	f.BtnM7.SetWidth(wm)
	f.BtnM7.SetHeight(hm)
	f.BtnM7.SetCaption("7月")
	f.BtnM7.SetOnClick(f.OnBtnM7)
	//
	f.BtnM8 = vcl.NewButton(f)
	f.BtnM8.SetParent(f)
	f.BtnM8.SetTop(topm)
	f.BtnM8.SetLeft(M8)
	f.BtnM8.SetWidth(wm)
	f.BtnM8.SetHeight(hm)
	f.BtnM8.SetCaption("8月")
	f.BtnM8.SetOnClick(f.OnBtnM8)
	//
	f.BtnM9 = vcl.NewButton(f)
	f.BtnM9.SetParent(f)
	f.BtnM9.SetTop(topm)
	f.BtnM9.SetLeft(M9)
	f.BtnM9.SetWidth(wm)
	f.BtnM9.SetHeight(hm)
	f.BtnM9.SetCaption("9月")
	f.BtnM9.SetOnClick(f.OnBtnM9)
	//
	f.BtnM10 = vcl.NewButton(f)
	f.BtnM10.SetParent(f)
	f.BtnM10.SetTop(topm)
	f.BtnM10.SetLeft(M10)
	f.BtnM10.SetWidth(wm)
	f.BtnM10.SetHeight(hm)
	f.BtnM10.SetCaption("10月")
	f.BtnM10.SetOnClick(f.OnBtnM10)
	//
	f.BtnM11 = vcl.NewButton(f)
	f.BtnM11.SetParent(f)
	f.BtnM11.SetTop(topm)
	f.BtnM11.SetLeft(M11)
	f.BtnM11.SetWidth(wm)
	f.BtnM11.SetHeight(hm)
	f.BtnM11.SetCaption("11月")
	f.BtnM11.SetOnClick(f.OnBtnM11)
	//
	f.BtnM12 = vcl.NewButton(f)
	f.BtnM12.SetParent(f)
	f.BtnM12.SetTop(topm)
	f.BtnM12.SetLeft(M12)
	f.BtnM12.SetWidth(wm)
	f.BtnM12.SetHeight(hm)
	f.BtnM12.SetCaption("12月")
	f.BtnM12.SetOnClick(f.OnBtnM12)
	//---------------------------------------
	f.Ltoday = vcl.NewLabel(f)
	f.Ltoday.SetParent(f)
	f.Ltoday.SetTop(570)
	f.Ltoday.SetLeft(leftLtoday)
	f.Ltoday.SetWidth(90)
	f.Ltoday.SetHeight(hm)
	f.Ltoday.Font().SetSize(14)
	f.Ltoday.SetTextBuf("今日阳历时间: " + time.Now().Format("2006-01-02"))
	//
	f.about = vcl.NewButton(f)
	f.about.SetParent(f)
	f.about.SetTop(570)
	f.about.SetLeft(leftAbout)
	f.about.SetWidth(70)
	f.about.SetHeight(hm)
	f.about.SetCaption("开源地址")
	f.about.Font().SetSize(12)
	f.about.SetOnClick(f.OnBtnAboutClick)
	//--------------------------
	f.today()
}
func info(y, m, d, h int) string {
	solar := fmt.Sprintf("%d月%d日", m, d) //阳历
	//
	t := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	wi := int(t.Weekday())
	warr := []string{"日", "一", "二", "三", "四", "五", "六"}
	var ws string
	for i := 0; i < len(warr); i++ {
		if wi == i {
			ws = "周" + warr[i]
			break
		}
	}
	//
	_, _, _, moon := basic.GetLunar(y, m, d) //阴历
	mgz := gz.GetMonthGZ(y, m, d, h)         //月干支
	dgz := gz.GetDayGZ(y, m, d)              //日干支
	jianChu := gz.GetRiJianChu(mgz, dgz)     //日建除
	huangHei := gz.GetRiHuangHei(mgz, dgz)   //日黄黑

	wd := int(t.Weekday())
	riQin := gz.GetRiQin(wd, dgz)

	s := solar + ws + "\n" + moon + "\n" + dgz + " " + jianChu + " " + huangHei + "\n" + riQin
	return s
}

const (
	top1   = 50
	tn     = 100
	top2   = top1 + tn
	top3   = top2 + tn
	top4   = top3 + tn
	top5   = top4 + tn
	width  = 150
	height = 60

	w     = 100
	left1 = 20
	left2 = left1 + w
	left3 = left2 + w
	left4 = left3 + w
	left5 = left4 + w
	left6 = left5 + w
	left7 = left6 + w
)

func (f *TMainForm) today() {
	//------------------------------------------
	t := time.Now().Local()
	y := t.Year()
	m := int(t.Month())
	d := t.Day()
	h := t.Hour()
	f.label = vcl.NewLabel(f)
	f.label.SetParent(f)
	f.label.SetTop(top1)
	f.label.SetWidth(width)
	f.label.SetHeight(height)
	f.label.SetLeft(left1)
	f.label.SetTextBuf(info(y, m, 1, h))
	//
	f.label2 = vcl.NewLabel(f)
	f.label2.SetParent(f)
	f.label2.SetTop(top1)
	f.label2.SetWidth(width)
	f.label2.SetHeight(height)
	f.label2.SetLeft(left2)
	f.label2.SetTextBuf(info(y, m, 2, h))
	////
	f.label3 = vcl.NewLabel(f)
	f.label3.SetParent(f)
	f.label3.SetTop(top1)
	f.label3.SetWidth(width)
	f.label3.SetHeight(height)
	f.label3.SetLeft(left3)
	f.label3.SetTextBuf(info(y, m, 3, h))
	////
	f.label4 = vcl.NewLabel(f)
	f.label4.SetParent(f)
	f.label4.SetTop(top1)
	f.label4.SetWidth(width)
	f.label4.SetHeight(height)
	f.label4.SetLeft(left4)
	f.label4.SetTextBuf(info(y, m, 4, h))
	////
	f.label5 = vcl.NewLabel(f)
	f.label5.SetParent(f)
	f.label5.SetTop(top1)
	f.label5.SetWidth(width)
	f.label5.SetHeight(height)
	f.label5.SetLeft(left5)
	f.label5.SetTextBuf(info(y, m, 5, h))
	////
	f.label6 = vcl.NewLabel(f)
	f.label6.SetParent(f)
	f.label6.SetTop(top1)
	f.label6.SetWidth(width)
	f.label6.SetHeight(height)
	f.label6.SetLeft(left6)
	f.label6.SetTextBuf(info(y, m, 6, h))
	////
	f.label7 = vcl.NewLabel(f)
	f.label7.SetParent(f)
	f.label7.SetTop(top1)
	f.label7.SetWidth(width)
	f.label7.SetHeight(height)
	f.label7.SetLeft(left7)
	f.label7.SetTextBuf(info(y, m, 7, h))
	//-------------------------------------
	f.label8 = vcl.NewLabel(f)
	f.label8.SetParent(f)
	f.label8.SetTop(top2)
	f.label8.SetWidth(width)
	f.label8.SetHeight(height)
	f.label8.SetLeft(left1)
	f.label8.SetTextBuf(info(y, m, 8, h))
	////
	f.label9 = vcl.NewLabel(f)
	f.label9.SetParent(f)
	f.label9.SetTop(top2)
	f.label9.SetWidth(width)
	f.label9.SetHeight(height)
	f.label9.SetLeft(left2)
	f.label9.SetTextBuf(info(y, m, 9, h))
	////
	f.label10 = vcl.NewLabel(f)
	f.label10.SetParent(f)
	f.label10.SetTop(top2)
	f.label10.SetWidth(width)
	f.label10.SetHeight(height)
	f.label10.SetLeft(left3)
	f.label10.SetTextBuf(info(y, m, 10, h))
	////
	f.label11 = vcl.NewLabel(f)
	f.label11.SetParent(f)
	f.label11.SetTop(top2)
	f.label11.SetWidth(width)
	f.label11.SetHeight(height)
	f.label11.SetLeft(left4)
	f.label11.SetTextBuf(info(y, m, 11, h))
	////
	f.label12 = vcl.NewLabel(f)
	f.label12.SetParent(f)
	f.label12.SetTop(top2)
	f.label12.SetWidth(width)
	f.label12.SetHeight(height)
	f.label12.SetLeft(left5)
	f.label12.SetTextBuf(info(y, m, 12, h))
	////
	f.label13 = vcl.NewLabel(f)
	f.label13.SetParent(f)
	f.label13.SetTop(top2)
	f.label13.SetWidth(width)
	f.label13.SetHeight(height)
	f.label13.SetLeft(left6)
	f.label13.SetTextBuf(info(y, m, 13, h))
	////
	f.label14 = vcl.NewLabel(f)
	f.label14.SetParent(f)
	f.label14.SetTop(top2)
	f.label14.SetWidth(width)
	f.label14.SetHeight(height)
	f.label14.SetLeft(left7)
	f.label14.SetTextBuf(info(y, m, 14, h))
	//---------------
	f.label15 = vcl.NewLabel(f)
	f.label15.SetParent(f)
	f.label15.SetTop(top3)
	f.label15.SetWidth(width)
	f.label15.SetHeight(height)
	f.label15.SetLeft(left1)
	f.label15.SetTextBuf(info(y, m, 15, h))
	////-------
	f.label16 = vcl.NewLabel(f)
	f.label16.SetParent(f)
	f.label16.SetTop(top3)
	f.label16.SetWidth(width)
	f.label16.SetHeight(height)
	f.label16.SetLeft(left2)
	f.label16.SetTextBuf(info(y, m, 16, h))
	////
	f.label17 = vcl.NewLabel(f)
	f.label17.SetParent(f)
	f.label17.SetTop(top3)
	f.label17.SetWidth(width)
	f.label17.SetHeight(height)
	f.label17.SetLeft(left3)
	f.label17.SetTextBuf(info(y, m, 17, h))
	////
	f.label18 = vcl.NewLabel(f)
	f.label18.SetParent(f)
	f.label18.SetTop(top3)
	f.label18.SetWidth(width)
	f.label18.SetHeight(height)
	f.label18.SetLeft(left4)
	f.label18.SetTextBuf(info(y, m, 18, h))
	////
	f.label19 = vcl.NewLabel(f)
	f.label19.SetParent(f)
	f.label19.SetTop(top3)
	f.label19.SetWidth(width)
	f.label19.SetHeight(height)
	f.label19.SetLeft(left5)
	f.label19.SetTextBuf(info(y, m, 19, h))
	////
	f.label20 = vcl.NewLabel(f)
	f.label20.SetParent(f)
	f.label20.SetTop(top3)
	f.label20.SetWidth(width)
	f.label20.SetHeight(height)
	f.label20.SetLeft(left6)
	f.label20.SetTextBuf(info(y, m, 20, h))
	////
	f.label21 = vcl.NewLabel(f)
	f.label21.SetParent(f)
	f.label21.SetTop(top3)
	f.label21.SetWidth(width)
	f.label21.SetHeight(height)
	f.label21.SetLeft(left7)
	f.label21.SetTextBuf(info(y, m, 21, h))
	////
	f.label22 = vcl.NewLabel(f)
	f.label22.SetParent(f)
	f.label22.SetTop(top4)
	f.label22.SetWidth(width)
	f.label22.SetHeight(height)
	f.label22.SetLeft(left1)
	f.label22.SetTextBuf(info(y, m, 22, h))
	////
	f.label23 = vcl.NewLabel(f)
	f.label23.SetParent(f)
	f.label23.SetTop(top4)
	f.label23.SetWidth(width)
	f.label23.SetHeight(height)
	f.label23.SetLeft(left2)
	f.label23.SetTextBuf(info(y, m, 23, h))
	////
	f.label24 = vcl.NewLabel(f)
	f.label24.SetParent(f)
	f.label24.SetTop(top4)
	f.label24.SetWidth(width)
	f.label24.SetHeight(height)
	f.label24.SetLeft(left3)
	f.label24.SetTextBuf(info(y, m, 24, h))
	////
	f.label25 = vcl.NewLabel(f)
	f.label25.SetParent(f)
	f.label25.SetTop(top4)
	f.label25.SetWidth(width)
	f.label25.SetHeight(height)
	f.label25.SetLeft(left4)
	f.label25.SetTextBuf(info(y, m, 25, h))
	////
	f.label26 = vcl.NewLabel(f)
	f.label26.SetParent(f)
	f.label26.SetTop(top4)
	f.label26.SetWidth(width)
	f.label26.SetHeight(height)
	f.label26.SetLeft(left5)
	f.label26.SetTextBuf(info(y, m, 26, h))
	////
	f.label27 = vcl.NewLabel(f)
	f.label27.SetParent(f)
	f.label27.SetTop(top4)
	f.label27.SetWidth(width)
	f.label27.SetHeight(height)
	f.label27.SetLeft(left6)
	f.label27.SetTextBuf(info(y, m, 27, h))
	////
	f.label28 = vcl.NewLabel(f)
	f.label28.SetParent(f)
	f.label28.SetTop(top4)
	f.label28.SetWidth(width)
	f.label28.SetHeight(height)
	f.label28.SetLeft(left7)
	f.label28.SetTextBuf(info(y, m, 28, h))
	//--------------------------------------
	f.label29 = vcl.NewLabel(f)
	f.label29.SetParent(f)
	f.label29.SetTop(top5)
	f.label29.SetWidth(width)
	f.label29.SetHeight(height)
	f.label29.SetLeft(left1)
	f.label29.SetTextBuf(info(y, m, 29, h))
	////
	f.label30 = vcl.NewLabel(f)
	f.label30.SetParent(f)
	f.label30.SetTop(top5)
	f.label30.SetWidth(width)
	f.label30.SetHeight(height)
	f.label30.SetLeft(left2)
	f.label30.SetTextBuf(info(y, m, 30, h))
	////
	f.label31 = vcl.NewLabel(f)
	f.label31.SetParent(f)
	f.label31.SetTop(top5)
	f.label31.SetWidth(width)
	f.label31.SetHeight(height)
	f.label31.SetLeft(left3)
	f.label31.SetTextBuf(info(y, m, 31, h))
	//-------------
	if m == 4 || m == 6 || m == 9 || m == 11 {
		f.label31.Free()
		f.label31 = vcl.NewLabel(f)
		f.label31.SetParent(f)
		f.label31.SetTop(top5)
		f.label31.SetWidth(width)
		f.label31.SetHeight(height)
		f.label31.SetLeft(left3)
		f.label31.SetTextBuf("")
	}
	b := (y%4 == 0 && y%100 != 0) || y%400 == 0
	//闰年2月
	if m == 2 && b == true {
		f.label30.Free()
		f.label30 = vcl.NewLabel(f)
		f.label30.SetParent(f)
		f.label30.SetTop(top5)
		f.label30.SetWidth(width)
		f.label30.SetHeight(height)
		f.label30.SetLeft(left2)
		f.label30.SetTextBuf("")

		f.label31.Free()
		f.label31 = vcl.NewLabel(f)
		f.label31.SetParent(f)
		f.label31.SetTop(top5)
		f.label31.SetWidth(width)
		f.label31.SetHeight(height)
		f.label31.SetLeft(left3)
		f.label31.SetTextBuf("")
	}
	if m == 2 && b == false {
		f.label29.Free()
		f.label29 = vcl.NewLabel(f)
		f.label29.SetParent(f)
		f.label29.SetTop(top5)
		f.label29.SetWidth(width)
		f.label29.SetHeight(height)
		f.label29.SetLeft(left1)
		f.label29.SetTextBuf("")
		////
		f.label30.Free()
		f.label30 = vcl.NewLabel(f)
		f.label30.SetParent(f)
		f.label30.SetTop(top5)
		f.label30.SetWidth(width)
		f.label30.SetHeight(height)
		f.label30.SetLeft(left2)
		f.label30.SetTextBuf("")
		////
		f.label31.Free()
		f.label31 = vcl.NewLabel(f)
		f.label31.SetParent(f)
		f.label31.SetTop(top5)
		f.label31.SetWidth(width)
		f.label31.SetHeight(height)
		f.label31.SetLeft(left3)
		f.label31.SetTextBuf("")
	}

	//-----------------------------------------
	color := f.label.Color().RGB(255, 255, 255)
	labelArr := []*vcl.TLabel{f.label, f.label2, f.label3, f.label4, f.label5, f.label6, f.label7,
		f.label8, f.label9, f.label10, f.label11, f.label12, f.label13, f.label14,
		f.label15, f.label16, f.label17, f.label18, f.label19, f.label20, f.label21,
		f.label22, f.label23, f.label24, f.label25, f.label26, f.label27, f.label28,
		f.label29, f.label30, f.label31}
	for i := 0; i < len(labelArr); i++ {
		if i == d-1 {
			labelArr[i].SetColor(color)
			break
		}
	}
}

func (f *TMainForm) OnBtnOK(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	months := text[4:6]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	m, err := strconv.Atoi(months)
	if err != nil {
		fmt.Println("months", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	f.ShowMx(y, m, d, h)
}

func (f *TMainForm) OnBtnM1(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	//
	f.ShowMx(y, 1, d, h)
}

func (f *TMainForm) OnBtnM2(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	f.ShowMx(y, 2, d, h)
}

func (f *TMainForm) OnBtnM3(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	f.ShowMx(y, 3, d, h)
}

func (f *TMainForm) OnBtnM4(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	f.ShowMx(y, 4, d, h)
}

func (f *TMainForm) OnBtnM5(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	f.ShowMx(y, 5, d, h)
}

func (f *TMainForm) OnBtnM6(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	f.ShowMx(y, 6, d, h)
}

func (f *TMainForm) OnBtnM7(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	f.ShowMx(y, 7, d, h)
}

func (f *TMainForm) OnBtnM8(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	f.ShowMx(y, 8, d, h)
}

func (f *TMainForm) OnBtnM9(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	f.ShowMx(y, 9, d, h)
}

func (f *TMainForm) OnBtnM10(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	f.ShowMx(y, 10, d, h)
}

func (f *TMainForm) OnBtnM11(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	f.ShowMx(y, 11, d, h)
}

func (f *TMainForm) OnBtnM12(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	days := text[6:8]
	hours := text[8:10]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}
	f.ShowMx(y, 12, d, h)
}

func (f *TMainForm) ShowMx(y int, m int, d int, h int) {
	f.label.Free()
	f.label = vcl.NewLabel(f)
	f.label.SetParent(f)
	f.label.SetTop(top1)
	f.label.SetWidth(width)
	f.label.SetHeight(height)
	f.label.SetLeft(left1)
	f.label.SetTextBuf(info(y, m, 1, h))
	//
	f.label2.Free()
	f.label2 = vcl.NewLabel(f)
	f.label2.SetParent(f)
	f.label2.SetTop(top1)
	f.label2.SetWidth(width)
	f.label2.SetHeight(height)
	f.label2.SetLeft(left2)
	f.label2.SetTextBuf(info(y, m, 2, h))
	///
	f.label3.Free()
	f.label3 = vcl.NewLabel(f)
	f.label3.SetParent(f)
	f.label3.SetTop(top1)
	f.label3.SetWidth(width)
	f.label3.SetHeight(height)
	f.label3.SetLeft(left3)
	f.label3.SetTextBuf(info(y, m, 3, h))
	////
	f.label4.Free()
	f.label4 = vcl.NewLabel(f)
	f.label4.SetParent(f)
	f.label4.SetTop(top1)
	f.label4.SetWidth(width)
	f.label4.SetHeight(height)
	f.label4.SetLeft(left4)
	f.label4.SetTextBuf(info(y, m, 4, h))
	////
	f.label5.Free()
	f.label5 = vcl.NewLabel(f)
	f.label5.SetParent(f)
	f.label5.SetTop(top1)
	f.label5.SetWidth(width)
	f.label5.SetHeight(height)
	f.label5.SetLeft(left5)
	f.label5.SetTextBuf(info(y, m, 5, h))
	////
	f.label6.Free()
	f.label6 = vcl.NewLabel(f)
	f.label6.SetParent(f)
	f.label6.SetTop(top1)
	f.label6.SetWidth(width)
	f.label6.SetHeight(height)
	f.label6.SetLeft(left6)
	f.label6.SetTextBuf(info(y, m, 6, h))
	////
	f.label7.Free()
	f.label7 = vcl.NewLabel(f)
	f.label7.SetParent(f)
	f.label7.SetTop(top1)
	f.label7.SetWidth(width)
	f.label7.SetHeight(height)
	f.label7.SetLeft(left7)
	f.label7.SetTextBuf(info(y, m, 7, h))
	//-------------------------------------
	f.label8.Free()
	f.label8 = vcl.NewLabel(f)
	f.label8.SetParent(f)
	f.label8.SetTop(top2)
	f.label8.SetWidth(width)
	f.label8.SetHeight(height)
	f.label8.SetLeft(left1)
	f.label8.SetTextBuf(info(y, m, 8, h))
	////
	f.label9.Free()
	f.label9 = vcl.NewLabel(f)
	f.label9.SetParent(f)
	f.label9.SetTop(top2)
	f.label9.SetWidth(width)
	f.label9.SetHeight(height)
	f.label9.SetLeft(left2)
	f.label9.SetTextBuf(info(y, m, 9, h))
	////
	f.label10.Free()
	f.label10 = vcl.NewLabel(f)
	f.label10.SetParent(f)
	f.label10.SetTop(top2)
	f.label10.SetWidth(width)
	f.label10.SetHeight(height)
	f.label10.SetLeft(left3)
	f.label10.SetTextBuf(info(y, m, 10, h))
	////
	f.label11.Free()
	f.label11 = vcl.NewLabel(f)
	f.label11.SetParent(f)
	f.label11.SetTop(top2)
	f.label11.SetWidth(width)
	f.label11.SetHeight(height)
	f.label11.SetLeft(left4)
	f.label11.SetTextBuf(info(y, m, 11, h))
	////
	f.label12.Free()
	f.label12 = vcl.NewLabel(f)
	f.label12.SetParent(f)
	f.label12.SetTop(top2)
	f.label12.SetWidth(width)
	f.label12.SetHeight(height)
	f.label12.SetLeft(left5)
	f.label12.SetTextBuf(info(y, m, 12, h))
	////
	f.label13.Free()
	f.label13 = vcl.NewLabel(f)
	f.label13.SetParent(f)
	f.label13.SetTop(top2)
	f.label13.SetWidth(width)
	f.label13.SetHeight(height)
	f.label13.SetLeft(left6)
	f.label13.SetTextBuf(info(y, m, 13, h))
	////
	f.label14.Free()
	f.label14 = vcl.NewLabel(f)
	f.label14.SetParent(f)
	f.label14.SetTop(top2)
	f.label14.SetWidth(width)
	f.label14.SetHeight(height)
	f.label14.SetLeft(left7)
	f.label14.SetTextBuf(info(y, m, 14, h))
	//---------------
	f.label15.Free()
	f.label15 = vcl.NewLabel(f)
	f.label15.SetParent(f)
	f.label15.SetTop(top3)
	f.label15.SetWidth(width)
	f.label15.SetHeight(height)
	f.label15.SetLeft(left1)
	f.label15.SetTextBuf(info(y, m, 15, h))
	////-------
	f.label16.Free()
	f.label16 = vcl.NewLabel(f)
	f.label16.SetParent(f)
	f.label16.SetTop(top3)
	f.label16.SetWidth(width)
	f.label16.SetHeight(height)
	f.label16.SetLeft(left2)
	f.label16.SetTextBuf(info(y, m, 16, h))
	////
	f.label17.Free()
	f.label17 = vcl.NewLabel(f)
	f.label17.SetParent(f)
	f.label17.SetTop(top3)
	f.label17.SetWidth(width)
	f.label17.SetHeight(height)
	f.label17.SetLeft(left3)
	f.label17.SetTextBuf(info(y, m, 17, h))
	////
	f.label18.Free()
	f.label18 = vcl.NewLabel(f)
	f.label18.SetParent(f)
	f.label18.SetTop(top3)
	f.label18.SetWidth(width)
	f.label18.SetHeight(height)
	f.label18.SetLeft(left4)
	f.label18.SetTextBuf(info(y, m, 18, h))
	////
	f.label19.Free()
	f.label19 = vcl.NewLabel(f)
	f.label19.SetParent(f)
	f.label19.SetTop(top3)
	f.label19.SetWidth(width)
	f.label19.SetHeight(height)
	f.label19.SetLeft(left5)
	f.label19.SetTextBuf(info(y, m, 19, h))
	////
	f.label20.Free()
	f.label20 = vcl.NewLabel(f)
	f.label20.SetParent(f)
	f.label20.SetTop(top3)
	f.label20.SetWidth(width)
	f.label20.SetHeight(height)
	f.label20.SetLeft(left6)
	f.label20.SetTextBuf(info(y, m, 20, h))
	////
	f.label21.Free()
	f.label21 = vcl.NewLabel(f)
	f.label21.SetParent(f)
	f.label21.SetTop(top3)
	f.label21.SetWidth(width)
	f.label21.SetHeight(height)
	f.label21.SetLeft(left7)
	f.label21.SetTextBuf(info(y, m, 21, h))
	////
	f.label22.Free()
	f.label22 = vcl.NewLabel(f)
	f.label22.SetParent(f)
	f.label22.SetTop(top4)
	f.label22.SetWidth(width)
	f.label22.SetHeight(height)
	f.label22.SetLeft(left1)
	f.label22.SetTextBuf(info(y, m, 22, h))
	////
	f.label23.Free()
	f.label23 = vcl.NewLabel(f)
	f.label23.SetParent(f)
	f.label23.SetTop(top4)
	f.label23.SetWidth(width)
	f.label23.SetHeight(height)
	f.label23.SetLeft(left2)
	f.label23.SetTextBuf(info(y, m, 23, h))
	////
	f.label24.Free()
	f.label24 = vcl.NewLabel(f)
	f.label24.SetParent(f)
	f.label24.SetTop(top4)
	f.label24.SetWidth(width)
	f.label24.SetHeight(height)
	f.label24.SetLeft(left3)
	f.label24.SetTextBuf(info(y, m, 24, h))
	////
	f.label25.Free()
	f.label25 = vcl.NewLabel(f)
	f.label25.SetParent(f)
	f.label25.SetTop(top4)
	f.label25.SetWidth(width)
	f.label25.SetHeight(height)
	f.label25.SetLeft(left4)
	f.label25.SetTextBuf(info(y, m, 25, h))
	////
	f.label26.Free()
	f.label26 = vcl.NewLabel(f)
	f.label26.SetParent(f)
	f.label26.SetTop(top4)
	f.label26.SetWidth(width)
	f.label26.SetHeight(height)
	f.label26.SetLeft(left5)
	f.label26.SetTextBuf(info(y, m, 26, h))
	////
	f.label27.Free()
	f.label27 = vcl.NewLabel(f)
	f.label27.SetParent(f)
	f.label27.SetTop(top4)
	f.label27.SetWidth(width)
	f.label27.SetHeight(height)
	f.label27.SetLeft(left6)
	f.label27.SetTextBuf(info(y, m, 27, h))
	////
	f.label28.Free()
	f.label28 = vcl.NewLabel(f)
	f.label28.SetParent(f)
	f.label28.SetTop(top4)
	f.label28.SetWidth(width)
	f.label28.SetHeight(height)
	f.label28.SetLeft(left7)
	f.label28.SetTextBuf(info(y, m, 28, h))
	//--------------------------------------
	f.label29.Free()
	f.label29 = vcl.NewLabel(f)
	f.label29.SetParent(f)
	f.label29.SetTop(top5)
	f.label29.SetWidth(width)
	f.label29.SetHeight(height)
	f.label29.SetLeft(left1)
	f.label29.SetTextBuf(info(y, m, 29, h))
	////
	f.label30.Free()
	f.label30 = vcl.NewLabel(f)
	f.label30.SetParent(f)
	f.label30.SetTop(top5)
	f.label30.SetWidth(width)
	f.label30.SetHeight(height)
	f.label30.SetLeft(left2)
	f.label30.SetTextBuf(info(y, m, 30, h))
	////
	f.label31.Free()
	f.label31 = vcl.NewLabel(f)
	f.label31.SetParent(f)
	f.label31.SetTop(top5)
	f.label31.SetWidth(width)
	f.label31.SetHeight(height)
	f.label31.SetLeft(left3)
	f.label31.SetTextBuf(info(y, m, 31, h))
	//-------------
	if m == 4 || m == 6 || m == 9 || m == 11 {
		f.label31.Free()
		f.label31 = vcl.NewLabel(f)
		f.label31.SetParent(f)
		f.label31.SetTop(top5)
		f.label31.SetWidth(width)
		f.label31.SetHeight(height)
		f.label31.SetLeft(left3)
		f.label31.SetTextBuf("")
	}
	b := (y%4 == 0 && y%100 != 0) || y%400 == 0
	//闰年2月
	if m == 2 && b == true {
		f.label30.Free()
		f.label30 = vcl.NewLabel(f)
		f.label30.SetParent(f)
		f.label30.SetTop(top5)
		f.label30.SetWidth(width)
		f.label30.SetHeight(height)
		f.label30.SetLeft(left2)
		f.label30.SetTextBuf("")

		f.label31.Free()
		f.label31 = vcl.NewLabel(f)
		f.label31.SetParent(f)
		f.label31.SetTop(top5)
		f.label31.SetWidth(width)
		f.label31.SetHeight(height)
		f.label31.SetLeft(left3)
		f.label31.SetTextBuf("")
	}
	if m == 2 && b == false {
		f.label29.Free()
		f.label29 = vcl.NewLabel(f)
		f.label29.SetParent(f)
		f.label29.SetTop(top5)
		f.label29.SetWidth(width)
		f.label29.SetHeight(height)
		f.label29.SetLeft(left1)
		f.label29.SetTextBuf("")

		f.label30.Free()
		f.label30 = vcl.NewLabel(f)
		f.label30.SetParent(f)
		f.label30.SetTop(top5)
		f.label30.SetWidth(width)
		f.label30.SetHeight(height)
		f.label30.SetLeft(left2)
		f.label30.SetTextBuf("")

		f.label31.Free()
		f.label31 = vcl.NewLabel(f)
		f.label31.SetParent(f)
		f.label31.SetTop(top5)
		f.label31.SetWidth(width)
		f.label31.SetHeight(height)
		f.label31.SetLeft(left3)
		f.label31.SetTextBuf("")
	}

	//-----------------------------------------
	color := f.label.Color().RGB(255, 255, 255)
	labelArr := []*vcl.TLabel{f.label, f.label2, f.label3, f.label4, f.label5, f.label6, f.label7,
		f.label8, f.label9, f.label10, f.label11, f.label12, f.label13, f.label14,
		f.label15, f.label16, f.label17, f.label18, f.label19, f.label20, f.label21,
		f.label22, f.label23, f.label24, f.label25, f.label26, f.label27, f.label28,
		f.label29, f.label30, f.label31}
	for i := 0; i < len(labelArr); i++ {
		if i == d-1 {
			labelArr[i].SetColor(color)
			break
		}
	}
}

func (f *TMainForm) OnBtnAboutClick(_ vcl.IObject) {
	vcl.ShowMessage("https://github.com/Aquarian-Age/xa")
}
