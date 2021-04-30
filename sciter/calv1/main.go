package main

import (
	"encoding/json"
	"fmt"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"liangzi.local/cal/cal"
	"log"
	"strconv"
	"time"
)

func main() {
	w, err := window.New(sciter.SW_TITLEBAR|
		sciter.SW_RESIZEABLE|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_ENABLE_DEBUG,
		&sciter.Rect{Left: 0, Top: 0, Right: 500, Bottom: 500})
	if err != nil {
		fmt.Println(err)
	}
	//w.LoadFile("cal.html")
	w.LoadHtml(html, "")
	root, _ := w.GetRootElement()
	setEl(root)
	setWinHandler(w)
	w.Show()
	w.Run()
}

//默认直接显示不需要点击
func setEl(root *sciter.Element) {
	p1, _ := root.SelectById("p1")
	y := time.Now().Local().Year()
	m := int(time.Now().Local().Month())
	d := time.Now().Local().Day()
	h := time.Now().Local().Hour()
	gzobj := cal.NewGanZhiInfo(y, m, d, h)
	ygz := gzobj.YearGZ
	mgz := gzobj.MonthGZ
	dgz := gzobj.DayGZ
	hgz := gzobj.HourGZ
	sf := fmt.Sprintf("%s年-%s月-%s日-%s时", ygz, mgz, dgz, hgz)
	p1.SetValue(sciter.NewValue(sf))
}

func setWinHandler(w *window.Window) {
	w.DefineFunction("ymdToGZ", ymdToGZ)
}
func ymdToGZ(args ...*sciter.Value) *sciter.Value {
	ly, lm, ld, lh := args[0].String(), args[1].String(), args[2].String(), args[3].String()
	y, m, d, h := ConvStoInt(ly, lm, ld, lh)
	gzinfo := cal.NewGanZhiInfo(y, m, d, h)
	gzb, err := json.Marshal(gzinfo)
	if err != nil {
		fmt.Println(err)
	}
	gz := string(gzb)
	return sciter.NewValue(gz)
}
func ConvStoInt(ys, ms, ds, hs string) (int, int, int, int) {
	y, err := strconv.Atoi(ys)
	if err != nil {
		log.Fatal("年份時間解析:", err)
	}

	m, err := strconv.Atoi(ms)
	if err != nil {
		log.Fatal("月份時間解析:", err)
	}
	d, err := strconv.Atoi(ds)
	if err != nil {
		log.Fatal("日期時間解析:", err)
	}
	h, err := strconv.Atoi(hs)
	if err != nil {
		log.Fatal("時辰解析:", err)
	}

	return y, m, d, h
}
