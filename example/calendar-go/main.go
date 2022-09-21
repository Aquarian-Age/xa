package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Aquarian-Age/xa/pkg/gz"
	"github.com/Aquarian-Age/xa/pkg/pub"
	"github.com/charmbracelet/lipgloss"
	"github.com/jwalton/gchalk"
	"github.com/starainrt/astro/calendar"
)

var (
	tn   = time.Now().Local()
	t    = flag.String("t", tn.Format(layout), layout)
	jq   = flag.Bool("jq", false, "显示二十四节气时间")
	star = flag.Bool("star", false, "显示阳历 干支历 建除 黄黑 日禽 阴历")
	cal  = flag.Bool("cal", false, "显示阴/阳历 干支")
)

const layout = "2006-01-02 15:04:05"

func main() {
	flag.Parse()
	tx, err := time.Parse(layout, *t)
	if err != nil {
		log.Println(err)
	}

	gzo := gz.NewTMGanZhi(tx.Year(), int(tx.Month()), tx.Day(), tx.Hour(), tx.Minute())

	info := gzo.Info()
	fmt.Println(info.String())

	if *jq {
		jieQi(tx)
	}
	if *cal {
		calFunc(tx)
	}

	if *star {
		starFunc(tx)
	}

}
func calFunc(tx time.Time) {
	y := tx.Year()
	m := int(tx.Month())
	var b bool
	if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
		b = true
	} else {
		b = false
	}
	txa := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.Local)
	weekn := int(txa.Weekday())

	var allDay int
	var days []int
	switch m {
	case 2:
		if b == true {
			allDay = 29
		} else {
			allDay = 28
		}
	case 1, 3, 5, 7, 8, 10, 12:
		allDay = 31
	case 4, 6, 9, 11:
		allDay = 30
	}

	for i := 1; i <= allDay; i++ {
		days = append(days, i)
	}
	lastMonthDay := days[:weekn]
	days = append(lastMonthDay, days...)

	var str1, str2, str3, str4, str5, str6 []lipgloss.Style
	var gzsty1, gzsty2, gzsty3, gzsty4, gzsty5, gzsty6 []lipgloss.Style
	var moonsty1, moonsty2, moonsty3, moonsty4, moonsty5, moonsty6 []lipgloss.Style
	const width = 6
	for i := 0; i < len(days); i++ {
		switch i {
		case 0, 1, 2, 3, 4, 5, 6:
			tx := time.Date(y, time.Month(m), int(days[i]), 0, 0, 0, 0, time.Local)
			strx := lipgloss.NewStyle().Width(width).SetString(strconv.Itoa(tx.Day())).Align(lipgloss.Center)
			str1 = append(str1, strx)

			dgz := gz.GetDayGZ(tx.Year(), int(tx.Month()), tx.Day())
			gzx := lipgloss.NewStyle().Width(width).SetString(dgz).Align(lipgloss.Center)
			gzsty1 = append(gzsty1, gzx)

			_, _, _, moon := calendar.ChineseLunar(tx)
			moon = moon[6:12]
			moonx := lipgloss.NewStyle().Width(width).SetString(moon).Align(lipgloss.Center)
			moonsty1 = append(moonsty1, moonx)
			//}
		case 7, 8, 9, 10, 11, 12, 13:
			tx := time.Date(y, time.Month(m), int(days[i]), 0, 0, 0, 0, time.Local)
			strx := lipgloss.NewStyle().Width(width).SetString(strconv.Itoa(tx.Day())).Align(lipgloss.Center)
			str2 = append(str2, strx)

			dgz := gz.GetDayGZ(tx.Year(), int(tx.Month()), tx.Day())
			gzx := lipgloss.NewStyle().Width(width).SetString(dgz).Align(lipgloss.Center)
			gzsty2 = append(gzsty2, gzx)

			_, _, _, moon := calendar.ChineseLunar(tx)
			moon = moon[6:12]
			moonx := lipgloss.NewStyle().Width(width).SetString(moon).Align(lipgloss.Center)
			moonsty2 = append(moonsty2, moonx)
		case 14, 15, 16, 17, 18, 19, 20:
			tx := time.Date(y, time.Month(m), int(days[i]), 0, 0, 0, 0, time.Local)
			strx := lipgloss.NewStyle().Width(width).SetString(strconv.Itoa(tx.Day())).Align(lipgloss.Center)
			str3 = append(str3, strx)

			dgz := gz.GetDayGZ(tx.Year(), int(tx.Month()), tx.Day())
			gzx := lipgloss.NewStyle().Width(width).SetString(dgz).Align(lipgloss.Center)
			gzsty3 = append(gzsty3, gzx)

			_, _, _, moon := calendar.ChineseLunar(tx)
			moon = moon[6:12]
			moonx := lipgloss.NewStyle().Width(width).SetString(moon).Align(lipgloss.Center)
			moonsty3 = append(moonsty3, moonx)
		case 21, 22, 23, 24, 25, 26, 27:
			tx := time.Date(y, time.Month(m), int(days[i]), 0, 0, 0, 0, time.Local)
			strx := lipgloss.NewStyle().Width(width).SetString(strconv.Itoa(tx.Day())).Align(lipgloss.Center)
			str4 = append(str4, strx)

			dgz := gz.GetDayGZ(tx.Year(), int(tx.Month()), tx.Day())
			gzx := lipgloss.NewStyle().Width(width).SetString(dgz).Align(lipgloss.Center)
			gzsty4 = append(gzsty4, gzx)

			_, _, _, moon := calendar.ChineseLunar(tx)
			moon = moon[6:12]
			moonx := lipgloss.NewStyle().Width(width).SetString(moon).Align(lipgloss.Center)
			moonsty4 = append(moonsty4, moonx)
		case 28, 29, 30, 31, 32, 33, 34:
			tx := time.Date(y, time.Month(m), int(days[i]), 0, 0, 0, 0, time.Local)
			strx := lipgloss.NewStyle().Width(width).SetString(strconv.Itoa(tx.Day())).Align(lipgloss.Center)
			str5 = append(str5, strx)

			dgz := gz.GetDayGZ(tx.Year(), int(tx.Month()), tx.Day())
			gzx := lipgloss.NewStyle().Width(width).SetString(dgz).Align(lipgloss.Center)
			gzsty5 = append(gzsty5, gzx)

			_, _, _, moon := calendar.ChineseLunar(tx)
			moon = moon[6:12]
			moonx := lipgloss.NewStyle().Width(width).SetString(moon).Align(lipgloss.Center)
			moonsty5 = append(moonsty5, moonx)
		case 35, 36, 37, 38, 39, 40, 41:
			tx := time.Date(y, time.Month(m), int(days[i]), 0, 0, 0, 0, time.Local)
			strx := lipgloss.NewStyle().Width(width).SetString(strconv.Itoa(tx.Day())).Align(lipgloss.Center)
			str6 = append(str6, strx)

			dgz := gz.GetDayGZ(tx.Year(), int(tx.Month()), tx.Day())
			gzx := lipgloss.NewStyle().Width(width).SetString(dgz).Align(lipgloss.Center)
			gzsty6 = append(gzsty6, gzx)

			_, _, _, moon := calendar.ChineseLunar(tx)
			moon = moon[6:12]
			moonx := lipgloss.NewStyle().Width(width).SetString(moon).Align(lipgloss.Center)
			moonsty6 = append(moonsty6, moonx)
		}
	}
	weeksty := []lipgloss.Style{
		lipgloss.NewStyle().Width(width).SetString("周日").Align(lipgloss.Center),
		lipgloss.NewStyle().Width(width).SetString("周一").Align(lipgloss.Center),
		lipgloss.NewStyle().Width(width).SetString("周二").Align(lipgloss.Center),
		lipgloss.NewStyle().Width(width).SetString("周三").Align(lipgloss.Center),
		lipgloss.NewStyle().Width(width).SetString("周四").Align(lipgloss.Center),
		lipgloss.NewStyle().Width(width).SetString("周五").Align(lipgloss.Center),
		lipgloss.NewStyle().Width(width).SetString("周六").Align(lipgloss.Center),
	}
	for i := 0; i < len(weeksty); i++ {
		fmt.Printf("%v", weeksty[i].String())
	}
	fmt.Println()
	//
	for i := 0; i < len(str1); i++ {
		if i < weekn {
			fmt.Printf("%v", str1[i].SetString(""))
		} else {
			fmt.Printf("%v", str1[i].String())
		}
	}
	fmt.Println()
	for i := 0; i < len(gzsty1); i++ {
		if i < weekn {
			fmt.Printf("%v", gzsty1[i].SetString(""))
		} else {
			fmt.Printf("%v", gzsty1[i].String())
		}
	}
	fmt.Println()
	for i := 0; i < len(moonsty1); i++ {
		if i < weekn {
			fmt.Printf("%v", moonsty1[i].SetString(""))
		} else {
			fmt.Printf("%v", moonsty1[i].String())
		}
	}
	fmt.Println()
	fmt.Println()
	//
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%v", str2[i].String())
	}
	fmt.Println()
	for i := 0; i < len(gzsty2); i++ {
		fmt.Printf("%v", gzsty2[i].String())
	}
	fmt.Println()
	for i := 0; i < len(moonsty2); i++ {
		fmt.Printf("%v", moonsty2[i].String())
	}
	fmt.Println()
	fmt.Println()
	//
	for i := 0; i < len(str3); i++ {
		fmt.Printf(str3[i].String())
	}
	fmt.Println()
	for i := 0; i < len(gzsty3); i++ {
		fmt.Printf("%v", gzsty3[i].String())
	}
	fmt.Println()
	for i := 0; i < len(moonsty3); i++ {
		fmt.Printf("%v", moonsty3[i].String())
	}
	fmt.Println()
	fmt.Println()
	//
	for i := 0; i < len(str4); i++ {
		fmt.Printf("%v", str4[i].String())
	}
	fmt.Println()
	for i := 0; i < len(gzsty4); i++ {
		fmt.Printf("%v", gzsty4[i].String())
	}
	fmt.Println()
	for i := 0; i < len(moonsty4); i++ {
		fmt.Printf("%v", moonsty4[i].String())
	}
	fmt.Println()
	fmt.Println()
	//
	for i := 0; i < len(str5); i++ {
		fmt.Printf("%v", str5[i].String())
	}
	fmt.Println()
	for i := 0; i < len(gzsty5); i++ {
		fmt.Printf("%v", gzsty5[i].String())
	}
	fmt.Println()
	for i := 0; i < len(moonsty5); i++ {
		fmt.Printf("%v", moonsty5[i].String())
	}
	fmt.Println()
	fmt.Println()
	//
	for i := 0; i < len(str6); i++ {
		fmt.Printf("%v", str6[i].String())
	}
	fmt.Println()
	for i := 0; i < len(gzsty6); i++ {
		fmt.Printf("%v", gzsty6[i].String())
	}
	fmt.Println()
	for i := 0; i < len(moonsty6); i++ {
		fmt.Printf("%v", moonsty6[i].String())
	}
	fmt.Println()
	fmt.Println()
}

func starFunc(tx time.Time) {

	y := tx.Year()
	m := int(tx.Month())
	var b bool
	if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
		b = true
	} else {
		b = false
	}

	var allDay int
	switch m {
	case 2:
		if b == true {
			allDay = 29
		} else {
			allDay = 28
		}
	case 1, 3, 5, 7, 8, 10, 12:
		allDay = 31
	case 4, 6, 9, 11:
		allDay = 30
	}

	for i := 1; i <= allDay; i++ {

		tx := time.Date(y, time.Month(m), i, tn.Hour(), 0, 0, 0, time.Local)
		week := pub.WeekName(int(tx.Weekday()))
		mgz := gz.GetMonthGZ(tx.Year(), int(tx.Month()), i, tx.Hour())
		dgz := gz.GetDayGZ(tx.Year(), int(tx.Month()), i)
		jc := gz.JianChu(mgz, dgz)
		hh := gz.HuangHei(mgz, dgz)
		riqin := gz.GetRiQin(int(tx.Weekday()), dgz)
		_, _, _, moon := calendar.ChineseLunar(tx)
		s := fmt.Sprintf("%d月%d日 %s %s月%s日 %s %s %s %s",
			int(tx.Month()), i, week, mgz, dgz, jc, hh, riqin, moon)
		if tn.Day() == tx.Day() {
			fmt.Println(gchalk.Blue(s))
		} else {
			fmt.Println(s)
		}

	}
}

func jieQi(tx time.Time) {
	gzo := gz.NewTMGanZhi(tx.Year(), int(tx.Month()), tx.Day(), tx.Hour(), tx.Minute())
	jqarr := gzo.Jq24()
	for i := 0; i < len(jqarr); i++ {
		fmt.Println(jqarr[i])
	}
}
