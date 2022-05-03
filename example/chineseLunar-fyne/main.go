package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Aquarian-Age/xa/pkg/gz"
)

type UI struct {
	ent                       *widget.Entry
	selectm, selectd, selecth *widget.Select
	btn                       *widget.Button
	btnAbout                  *widget.Button
	selectCon                 *fyne.Container
	btnCon                    *fyne.Container
	Tieles                    []*widget.Label
	TitleCon                  *fyne.Container
	labels                    []*widget.Label
	label1Con                 *fyne.Container
	label2Con                 *fyne.Container
	label3Con                 *fyne.Container
	label4Con                 *fyne.Container
	label5Con                 *fyne.Container
	label6Con                 *fyne.Container
	labelInfo                 *widget.Label
}

var T = time.Now().Local()

func main() {
	a := app.New()
	a.SetIcon(resourceIconPng)
	a.Settings().SetTheme(&Biu{})
	w := a.NewWindow("简易星历")
	w.Resize(fyne.Size{Width: 360, Height: 640})
	w.CenterOnScreen()

	ui := newUI()
	ui.LabelLayout()
	ui.OnBtnClicked()
	ui.Run(w)
}
func newUI() *UI {
	inputYear := strconv.Itoa(T.Year())
	ent := widget.NewEntry()
	ent.SetText(inputYear)

	months := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	days := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}
	hours := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}

	selectm := widget.NewSelect(months, func(s string) {
	})
	selectm.PlaceHolder = "month"
	selectd := widget.NewSelect(days, func(s string) {
	})
	selectd.PlaceHolder = "day"
	selecth := widget.NewSelect(hours, func(s string) {
	})
	selecth.PlaceHolder = "hour"

	if selectm.SelectedIndex() == -1 {
		selectm.SetSelectedIndex(int(T.Month()) - 1)
	}
	if selectd.SelectedIndex() == -1 {
		selectd.SetSelectedIndex(T.Day() - 1)
	}
	if selecth.SelectedIndex() == -1 {
		selecth.SetSelectedIndex(T.Hour())
	}
	b := widget.NewButton("btn", func() {
	})
	ba := widget.NewButton("About", func() {
		subw := fyne.CurrentApp().NewWindow("")
		subw.CenterOnScreen()
		texts := `简易星历
v1.0.3
mail: bGlhbmd6aUB5YW5kZXguY29tCg==
本软件(简易星历)为免费软件
`
		label := widget.NewLabelWithStyle(texts, fyne.TextAlignLeading, fyne.TextStyle{})

		link, _ := url.Parse("https://github.com/Aquarian-Age/xa/tree/master/example/chineseLunar-fyne")
		hyperlink := widget.NewHyperlink("源码", link)
		hyperlink.TextStyle = fyne.TextStyle{Bold: true}

		linkfont, _ := url.Parse("http://wenq.org/wqy2/index.cgi")
		hyperlinkFont := widget.NewHyperlink("文泉驿字体", linkfont)

		label.Wrapping = fyne.TextWrapWord
		hyperlink.Alignment = fyne.TextAlignLeading

		showx := container.New(layout.NewVBoxLayout(), label, hyperlink, hyperlinkFont)
		subw.SetContent(showx)
		subw.Resize(fyne.Size{Width: 360, Height: 640})
		subw.Show()
	})
	//tieles
	title0 := widget.NewLabel("周日")
	title1 := widget.NewLabel("周一")
	title2 := widget.NewLabel("周二")
	title3 := widget.NewLabel("周三")
	title4 := widget.NewLabel("周四")
	title5 := widget.NewLabel("周五")
	title6 := widget.NewLabel("周六")
	titles := []*widget.Label{title0, title1, title2, title3, title4, title5, title6}

	//labels
	labels := []*widget.Label{
		label1, label2, label3, label4, label5, label6, label7,
		label8, label9, label10, label11, label12, label13, label14,
		label15, label16, label17, label18, label19, label20, label21,
		label22, label23, label24, label25, label26, label27, label28,
		label29, label30, label31, label32, label33, label34, label35,
		label36, label37, label38, label39, label40, label41, label42,
	}
	labelInfo := widget.NewLabel("")

	selectc := container.New(layout.NewGridLayout(4), ent, selectm, selectd, selecth)
	btnc := container.New(layout.NewGridLayout(2), b, ba)
	titlec := container.New(layout.NewGridLayout(7), title0, title1, title2, title3, title4, title5, title6)

	return &UI{
		ent:       ent,
		selectm:   selectm,
		selectd:   selectd,
		selecth:   selecth,
		btn:       b,
		btnAbout:  ba,
		selectCon: selectc,
		btnCon:    btnc,
		Tieles:    titles,
		TitleCon:  titlec,
		labels:    labels,
		labelInfo: labelInfo,
	}
}

func (ui *UI) LabelLayout() {
	ui.showLables(T)
	gwl := layout.NewGridWrapLayout(fyne.Size{Width: 45, Height: 35})
	ui.label1Con = container.New(gwl, label1, label2, label3, label4, label5, label6, label7)
	ui.label2Con = container.New(gwl, label8, label9, label10, label11, label12, label13, label14)
	ui.label3Con = container.New(gwl, label15, label16, label17, label18, label19, label20, label21)
	ui.label4Con = container.New(gwl, label22, label23, label24, label25, label26, label27, label28)
	ui.label5Con = container.New(gwl, label29, label30, label31, label32, label33, label34, label35)
	ui.label6Con = container.New(layout.NewGridWrapLayout(fyne.Size{Width: 45, Height: 10}), label36, label37, label38, label39, label40, label41, label42)
}
func (ui *UI) OnBtnClicked() {
	ent := ui.ent
	selectm := ui.selectm
	selectd := ui.selectd
	selecth := ui.selecth
	ui.btn.OnTapped = func() {
		years := ent.Text
		year, err := strconv.Atoi(years)
		if err != nil {
			fmt.Println(err)
		}
		month := selectm.SelectedIndex() + 1
		day := selectd.SelectedIndex() + 1
		hour := selecth.SelectedIndex()
		tx := time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.Local)
		ui.showLables(tx)
		gzo := gz.NewGanZhi(year, month, day, hour)
		solars := tx.String()[:16]
		moons, _ := gzo.GetLunar()
		info := fmt.Sprintf("\n%s %s %s %s\n", gzo.Ygz, gzo.Mgz, gzo.Dgz, gzo.Hgz)
		s := solars + info + moons
		ui.labelInfo.Refresh()
		ui.labelInfo.SetText(s)
	}
}
func (ui *UI) Run(w fyne.Window) {
	connect := container.New(layout.NewVBoxLayout(), ui.selectCon, ui.btnCon, ui.TitleCon,
		ui.label1Con, ui.label2Con, ui.label3Con, ui.label4Con, ui.label5Con, ui.label6Con,
		ui.labelInfo)
	w.SetContent(connect)
	w.ShowAndRun()
}

var (
	label1  = widget.NewLabel("")
	label2  = widget.NewLabel("")
	label3  = widget.NewLabel("")
	label4  = widget.NewLabel("")
	label5  = widget.NewLabel("")
	label6  = widget.NewLabel("")
	label7  = widget.NewLabel("")
	label8  = widget.NewLabel("")
	label9  = widget.NewLabel("")
	label10 = widget.NewLabel("")
	label11 = widget.NewLabel("")
	label12 = widget.NewLabel("")
	label13 = widget.NewLabel("")
	label14 = widget.NewLabel("")
	label15 = widget.NewLabel("")
	label16 = widget.NewLabel("")
	label17 = widget.NewLabel("")
	label18 = widget.NewLabel("")
	label19 = widget.NewLabel("")
	label20 = widget.NewLabel("")
	label21 = widget.NewLabel("")
	label22 = widget.NewLabel("")
	label23 = widget.NewLabel("")
	label24 = widget.NewLabel("")
	label25 = widget.NewLabel("")
	label26 = widget.NewLabel("")
	label27 = widget.NewLabel("")
	label28 = widget.NewLabel("")
	label29 = widget.NewLabel("")
	label30 = widget.NewLabel("")
	label31 = widget.NewLabel("")
	label32 = widget.NewLabel("")
	label33 = widget.NewLabel("")
	label34 = widget.NewLabel("")
	label35 = widget.NewLabel("")
	label36 = widget.NewLabel("")
	label37 = widget.NewLabel("")
	label38 = widget.NewLabel("")
	label39 = widget.NewLabel("")
	label40 = widget.NewLabel("")
	label41 = widget.NewLabel("")
	label42 = widget.NewLabel("")
)
