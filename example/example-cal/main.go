/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package main

import (
	"fmt"
	"github.com/Aquarian-Age/xa/pkg/cal"
	"time"
)

var t = time.Now().Local()

func main() {
	t = time.Date(2033, time.Month(12), 22, 0, 0, 0, 0, time.Local)
	lunar := cal.NewLunar(t.Year(), int(t.Month()), t.Day())
	var lunars string
	if lunar.LeapM != 0 && lunar.LeapRmc != "" {
		lunars = fmt.Sprintf("%d年闰%s月(%s)%s %s",
			lunar.LY, cal.Alias(lunar.LeapM), lunar.LeapYdx, lunar.LeapRmc, lunar.Week)
	} else if lunar.LeapRmc == "" {
		lunars = fmt.Sprintf("%d年%s月(%s)%s %s",
			lunar.LY, cal.Alias(lunar.LM), lunar.Ydx, lunar.LRmc, lunar.Week)
	}
	fmt.Println(lunars) //阴历信息

	gz := cal.NewCal(t.Year(), int(t.Month()), t.Day(), t.Hour())
	gzs := gz.YearGZ + "年 " + gz.MonthGZ + "月 " + gz.DayGZ + "日 " + gz.HourGZ + "时"
	fmt.Println(gzs) //阳历干支
	nys := gz.NaYin()
	fmt.Println(nys) //纳因

	sw := cal.NewShuoWangTS(t.Year(), int(t.Month()), t.Day())
	shuo := "朔: " + sw.ShuoTS
	wang := "望: " + sw.WangTS
	shang := "上弦: " + sw.ShangXianTS
	xia := "下弦: " + sw.XiaXianTS
	sws := shuo + "\n" + wang + "\n" + shang + "\n" + xia + "\n"
	fmt.Println(sws) //朔望时间

	jqt := cal.YueJiangJQT(t.Year())
	yjq := cal.NewYueJiangJQ(jqt)
	jqs := yjq.JQ24() //24节气
	fmt.Println(jqs)
	//for i := 0; i < len(yjq.Time); i++ {
	//	fmt.Printf("%s %v\n", yjq.Name[i], yjq.Time[i].Format("2006-01-02 15:04:05")) //２４节气
	//}
	weekn := int(t.Weekday())
	stars := cal.Star(weekn, gz.DayGZ)
	fmt.Println(stars) //日禽

	dmjs := cal.Dmj(t.Year(), int(t.Month()), t.Day())
	fmt.Println(dmjs)
}
