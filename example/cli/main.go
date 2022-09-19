package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Aquarian-Age/xa/pkg/gz"
	"github.com/Aquarian-Age/xa/pkg/pub"
	"github.com/jwalton/gchalk"
	"github.com/starainrt/astro/calendar"
)

var (
	tn   = time.Now().Local()
	t    = flag.String("t", tn.Format(layout), layout)
	star = flag.Bool("star", false, "显示阳历 干支历 建除 黄黑 日禽 阴历")
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

	if *star {

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
			s := fmt.Sprintf("%d月%d日 周%s %s月%s日 %s %s %s %s",
				int(tx.Month()), i, week, mgz, dgz, jc, hh, riqin, moon)
			if tn.Day() == tx.Day() {
				fmt.Println(gchalk.Blue(s))
			} else {
				fmt.Println(s)
			}

		}
	}
}
