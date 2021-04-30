package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"liangzi.local/qx"
	"liangzi.local/qx/pkg/x"
	"liangzi.local/qx/pkg/xjbfs"
)

func init() {
	_ = os.Setenv("FYNE_FONT", "font/SourceHanSerifCN-SemiBold.ttf")
	//_ = os.Setenv("FYNE_FONT", "/system/fonts/SourceHanSerifCN-SemiBold.ttf") //apk字体
}
func main() {
	a := app.New()
	w := a.NewWindow("演禽")                    //标题
	a.Settings().SetTheme(theme.LightTheme()) //白色主题

	t := time.Now().Local()
	ys := fmt.Sprintf("%d", t.Year())
	ent := widget.NewEntry()
	ent.SetText(ys)

	//select options
	months := func() []string {
		var s []string
		for i := 1; i <= 12; i++ {
			is := fmt.Sprintf("%d月", i)
			s = append(s, is)
		}
		return s
	}
	days := func() []string {
		var s []string
		for i := 1; i <= 31; i++ {
			is := fmt.Sprintf("%d日", i)
			s = append(s, is)
		}
		return s
	}

	hours := func() []string {
		var s []string
		for i := 0; i <= 23; i++ {
			is := fmt.Sprintf("%d时", i)
			s = append(s, is)
		}
		return s
	}
	selectm := widget.NewSelect(months(), func(s string) {
	})
	selectm.PlaceHolder = "month"
	selectd := widget.NewSelect(days(), func(s string) {
	})
	selectd.PlaceHolder = "day"
	selecth := widget.NewSelect(hours(), func(s string) {
	})
	selecth.PlaceHolder = "hour"

	label := widget.NewLabel("")
	label3 := widget.NewLabel("")
	label2 := widget.NewLabel("")
	label1 := widget.NewLabel("")
	label1.Wrapping = fyne.TextWrapWord //长文本自动换行

	btnOK := widget.NewButton("OK", func() {
		ys := ent.Text
		if selectm.SelectedIndex() == -1 {
			selectm.SetSelectedIndex(int(t.Month()) - 1)
		}
		if selectd.SelectedIndex() == -1 {
			selectd.SetSelectedIndex(t.Day() - 1)
		}
		if selecth.SelectedIndex() == -1 {
			selecth.SetSelectedIndex(t.Hour())
		}
		year, err := strconv.Atoi(ys)
		if err != nil {
			fmt.Println(err)
		}
		month := selectm.SelectedIndex() + 1
		day := selectd.SelectedIndex() + 1
		hour := selecth.SelectedIndex()
		//显示到主页面
		ymdh := fmt.Sprintf("%d年-%d月-%d日-%d时", year, month, day, hour)
		label.SetText(ymdh)
		yq := get(year, month, day, hour)
		label3.SetText(yq.GetGanZhi() + "\n" +
			yq.GetQiKe() + "\n" +
			yq.GetJiangTou() + "\n" +
			yq.GetFanQinHuoYao() + "\n" +
			yq.GetFuJiang() + "\n" +
			yq.QiSha())

		suoBoH := yq.SuoBoJueH()
		xjbf := getXjbf(yq.Ygz, yq.Mgz, yq.Dgz, yq.Hgz)
		rijc := xjbf.JianChu()
		jcjx := x.UseJianChu(rijc, yq.RiQin)

		starArr := qx.Star420()
		riQinN, riQinindex, _ := qx.FindStarIndexAndGzIndex(yq.Dgz, yq.RiQin, starArr)
		ds := fmt.Sprintf("日禽%s所属%d元 本元六十索引值%d\n", yq.RiQin, riQinN, riQinindex)
		indexYuan, index, _ := qx.FindStarIndexAndGzIndex(yq.Hgz, yq.ShiQin, starArr)
		hs := fmt.Sprintf("时禽%s所属%d元 本元六十索引值%d\n", yq.ShiQin, indexYuan, index) //
		label2.SetText(suoBoH + "\n" + rijc + " " + jcjx + "\n" + ds + hs)

		yiji := yq.Yiji()
		//bmjs := yq.TXbenMingJiShen()
		qiucai := yq.QiuCai()
		jitanbing := xjbfs.JiTanBing(yq.Dgz)
		liuXiongRi := yq.LiuEeQin()
		tianNan := yq.TianNanRi()
		shangJi := yq.Shangji()
		sanXiongXing := yq.SanXiongXing()
		msg1, msg2 := x.JiuXingFangWei(year)

		label1.SetText(yq.GetSimpleJX() + "\n" + yiji + "\n" + qiucai + "\n" + jitanbing +
			"\n" + liuXiongRi + "\n\n" + tianNan + "\n" + shangJi + "\n" + sanXiongXing +
			"\n" + msg1 + "\n" + msg2)
	})
	////img
	//img := canvas.NewImageFromFile("img/震邪符咒.png")
	//img.FillMode = canvas.ImageFillContain
	//c := container.New(layout.NewGridWrapLayout(fyne.NewSize(150, 150)), img)

	//
	hbox := container.New(layout.NewHBoxLayout(), ent, selectm, selectd, selecth, btnOK)

	//
	vbox := container.New(layout.NewVBoxLayout(), label, label3, label2, label1)

	w.SetContent(container.New(layout.NewVBoxLayout(), hbox, vbox))
	w.Resize(fyne.Size{380, 420})
	w.ShowAndRun()

}
