package main

import (
	"fmt"
	"github.com/Aquarian-Age/xa/pkg/gz"
	"github.com/starainrt/astro/basic"
	"github.com/starainrt/astro/calendar"
	"math"
	"sort"
	"time"
)

func dayGZ(year, month, day int) string {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	jd := calendar.Date2JDE(t)
	jdI := int(math.Ceil(jd))
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
	monthNow int
	weekname = []string{"日", "一", "二", "三", "四", "五", "六"}
	dayMap   = map[int]string{1: "初一", 2: "初二", 3: "初三", 4: "初四", 5: "初五", 6: "初六", 7: "初七",
		8: "初八", 9: "初九", 10: "初十", 11: "十一", 12: "十二", 13: "十三", 14: "十四",
		15: "十五", 16: "十六", 17: "十七", 18: "十八", 19: "十九", 20: "廿十", 21: "廿一",
		22: "廿二", 23: "廿三", 24: "廿四", 25: "廿五", 26: "廿六", 27: "廿七", 28: "廿八",
		29: "廿九", 30: "三十"}
	Lefts    = []int{10, 50, 90, 130, 170, 210, 250}
	dayArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
		15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
		29, 30, 31}

	Gans    = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	Zhi     = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	jqnames = []string{
		"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
		"春分", "清明", "谷雨", "立夏", "小满", "芒种",
		"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
		"秋分", "寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
	}
	jqtmap = make(map[time.Time]string)
)

func (ui *UI) showLables(t time.Time) {
	monthNow = int(t.Month())
	lastmonth := monthNow - 1
	if lastmonth == 0 {
		lastmonth = 12
	}
	//判断闰月
	year := t.Year()
	b := (year%4 == 0 && year%100 != 0) || year%400 == 0
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
	wn := int(t.Weekday())
	jqt(t)

	var xday []int
	_days := dayArray[:wn]
	xday = append(xday, _days...)
	for i := 0; i < allDay; i++ {
		xday = append(xday, i+1)
	}

	//使文所有label的文本为空
	for i := 0; i < len(ui.labels); i++ {
		ui.labels[i].SetText("")
	}

	var moon string
	for i := 0; i < len(xday); i++ {
		tx := time.Date(year, time.Month(monthNow), xday[i], 0, 0, 0, 0, time.Local)

		dayx := fmt.Sprintf(" %d ", xday[i])
		dgz := dayGZ(year, monthNow, xday[i])
		starNames := getRiQin(tx)
		_, moonday, _, _ := calendar.ChineseLunar(tx)
		if jqs, ok := jqtmap[tx]; ok {
			moon = jqs
		} else {
			moon = dayMap[moonday]
		}

		s := dayx + starNames + "\n" + dgz + "\n" + moon
		if i >= wn {
			ui.labels[i].SetText(s)
		}
	}
}

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

//时间精确到日
func getRiQin(tx time.Time) string {
	jd := calendar.Date2JDE(tx)
	weekn := gz.WeekNumber(jd)
	aliaszhi := gz.AliasZhi(jd)
	return gz.AliasStarName(weekn, aliaszhi)[:3]
}
