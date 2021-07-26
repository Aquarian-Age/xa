/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package cal

import (
	"time"
)

//农历信息
func NewLunar(y, m, d int) *Lunar {
	cust := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	allshuo := AllShuo(y)
	jqarrOBj := NewJQArr(y)
	leapmb := LeapmB(y)
	zqArrT := jqarrOBj.ZhongQiArrT()
	index, _, _ := m1(Data(y))
	shuoWangFArr := NewShuoWangF(index, Data(y))
	moonSWFarr := MoonShuoWangF(shuoWangFArr)
	moonSwArrT := MoonShuoWangT(moonSWFarr) //两冬至间的时间 从这里判断 阴历月是不是有不含中气的月份
	//阴历闰月判断
	var leapmN int
	var leapmT time.Time
	zqt := zqArrT[12]
	zqt = time.Date(zqt.Year(), zqt.Month(), zqt.Day(), 0, 0, 0, 0, time.Local)
	if cust.After(zqt) {
		y += 1
		index, _, _ = m1(Data(y))
		shuoWangFArr = NewShuoWangF(index, Data(y))
		moonSWFarr = MoonShuoWangF(shuoWangFArr)
		moonSwArrT = MoonShuoWangT(moonSWFarr)
		jqarrOBj = NewJQArr(y)

		leapmb = LeapmB(y)
		zqArrT = jqarrOBj.ZhongQiArrT()
		leapmN, leapmT = findLeapmN(leapmb, zqArrT, moonSwArrT)
	} else {
		//农历闰月数字及朔的阳历时间戳
		leapmN, leapmT = findLeapmN(leapmb, zqArrT, moonSwArrT)
	}
	mtarr := findLunarMN(allshuo)
	ly, lm := LunarMN(cust, mtarr, leapmb, leapmT)
	ydxarr := YueDaXiao(allshuo)
	ydx := findYdx(leapmb, leapmN, ydxarr, lm)
	ld, sdayN := LunarDay(cust, ydx, allshuo)
	week := WeekDay(int(cust.Weekday()))
	leaprmc, leapSdayN, leapydxN := leapYdx(leapmT, cust, leapmN, mtarr, ydxarr)
	var ydxs, leapydxNs string
	if ydx == 29 || leapydxN == 29 {
		ydxs = "小"
		leapydxNs = "小"
	} else if ydx == 30 || leapydxN == 30 {
		ydxs = "大"
		leapydxNs = "大"
	}

	lymd := new(Lunar)
	lymd = &Lunar{
		LY:       ly,
		LM:       lm,
		LRmc:     ld,
		Ydx:      ydxs,
		LeapM:    leapmN,
		LeapRmc:  leaprmc,
		LeapYdx:  leapydxNs,
		Sday:     sdayN,
		LeapSday: leapSdayN,
		Week:     week,
	}
	return lymd
}

func WeekDay(n int) (w string) {
	switch n {
	case 0:
		w = "周日"
	case 1:
		w = "周一"
	case 2:
		w = "周二"
	case 3:
		w = "周三"
	case 4:
		w = "周四"
	case 5:
		w = "周五"
	case 6:
		w = "周六"
	}
	return
}

//农历日名称 农历日对应的阳历日期 闰月月大小
func leapYdx(leapmT, cust time.Time, leapmN int, mtarr []time.Time, ydxarr []int) (string, int, int) {
	var ldayArr = []int{} //农历月数组 长度等于农历月大小数字
	var dayArr = []int{}  //阳历月数组 长度等于农历月大小数字
	var leapydx int
	var idx int
	for j := 0; j < 15; j++ {
		index := j + 11
		if index > 12 {
			index -= 12
		}
		if leapmN == index-1 && leapmT.Equal(mtarr[j]) {
			leapydx = ydxarr[j]
			idx = j
			break
		}
	}
	var leapRmc string
	var leapSdayN int

	for i := 0; i < leapydx-1; i++ {
		t := mtarr[idx]
		t = time.Date(t.Year(), t.Month(), t.Day()+i, 0, 0, 0, 0, time.Local)
		dayjd := SolarYmdToJD(t.Year(), int(t.Month()), t.Day())
		dayt := JdToLocalTime(float64(dayjd)) //阳历时间
		dayt = time.Date(dayt.Year(), dayt.Month(), dayt.Day(), 0, 0, 0, 0, time.Local)

		if dayt.Equal(cust) {
			leapRmc = Rmc[i]
			leapSdayN = dayt.Day()
		}
		ldayArr = append(ldayArr, i+1)
		dayArr = append(dayArr, dayt.Day())
	}
	return leapRmc, leapSdayN, leapydx
}

//农历日名称 对应的阳历日数字 农历日数组 阳历日数组 干支数组
func LunarDay(cust time.Time, ydx int, allshuo []float64) (string, int) {
	cust = time.Date(cust.Year(), cust.Month(), cust.Day(), 0, 0, 0, 0, time.Local)
	var idx int    //月大小数组索引值
	var rmc string //日名称
	var sdayN int  //日名称阳历时间

	var ldayArr = []int{}   //农历月数组 长度等于农历月大小数字
	var dayArr = []int{}    //阳历月数组 长度等于农历月大小数字
	var dgzArr = []string{} //日干支
	for i := 0; i < 15; i++ {
		jd0 := allshuo[i]
		jd1 := allshuo[i+1]
		t := JdToLocalTime(jd0)
		t1 := JdToLocalTime(jd1)
		t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
		t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
		if (cust.Equal(t) || cust.After(t)) && cust.Before(t1) {
			//使用月大小计算
			index := i + 11
			if index > 12 {
				index -= 12
			}
			idx = i
			break
		}
	}

	for j := 0; j <= ydx-1; j++ { //j农历日数字
		jd1 := allshuo[idx] //月初jd
		t := JdToLocalTime(jd1)
		t = time.Date(t.Year(), t.Month(), t.Day()+j, 0, 0, 0, 0, time.Local)
		dayjd := SolarYmdToJD(t.Year(), int(t.Month()), t.Day())
		dayt := JdToLocalTime(float64(dayjd)) //阳历时间
		dayt = time.Date(dayt.Year(), dayt.Month(), dayt.Day(), 0, 0, 0, 0, time.Local)
		ldayArr = append(ldayArr, j+1)
		dayArr = append(dayArr, dayt.Day())
		//干支
		dgz, _ := DayGZ(t.Year(), int(t.Month()), t.Day())
		dgzArr = append(dgzArr, dgz)

		if cust.Equal(dayt) {
			rmc = Rmc[j]
			sdayN = dayt.Day()
		}
	}

	return rmc, sdayN
}

//本月月大小
func findYdx(leapmb bool, leapmN int, ydxarr []int, lm int) int {
	var ydx int
	switch leapmb {
	case true:
		for i := 0; i < len(ydxarr); i++ {
			index := i + 11
			if index > 12 {
				index -= 12
			}
			if lm > leapmN {
				ydx = ydxarr[index-1]
				break
			} else if i < 2 { //上一年
				ydx = ydxarr[i]
			} else if i >= 2 {
				ydx = ydxarr[lm+1]
				break
			}
		}
	case false:
		for i := 0; i < len(ydxarr); i++ {
			if lm == i {
				ydx = ydxarr[i+1]
				break
			}
		}
	}
	return ydx
}

//从上年十一月开始 到本年十一月结束 十三个月的月大小
//计算农历月大小 这里时间精确到日
func YueDaXiao(allshuo []float64) []int {
	var ydx []int
	for i := 0; i < 15; i++ {
		t0jd := allshuo[i]
		t0 := JdToLocalTime(t0jd)
		t0 = time.Date(t0.Year(), t0.Month(), t0.Day(), 0, 0, 0, 0, time.Local)
		t1jd := allshuo[i+1]
		t1 := JdToLocalTime(t1jd)
		t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
		MaxDayNumber := t1.Sub(t0).Hours() / 24
		ydx = append(ydx, int(MaxDayNumber))
	}
	return ydx
}

//农历全年月数字这里闰月和非闰月是同一个数字 闰月
func LunarMN(cust time.Time, mtarr []time.Time, leapmb bool, leapmT time.Time) (int, int) {
	var lunarM int
	var lunarYn int
	cust = time.Date(cust.Year(), cust.Month(), cust.Day(), 0, 0, 0, 0, time.Local)

	if (cust.After(mtarr[0]) || cust.Equal(mtarr[0])) && cust.Before(mtarr[1]) {
		//这里年份-1
		lunarYn = cust.Year() - 1
		lunarM = 11
	} else if (cust.Equal(mtarr[1]) || cust.After(mtarr[1])) && cust.Before(mtarr[2]) {
		//这里年份-1
		lunarYn = cust.Year() - 1
		lunarM = 12
	} else {
		lunarYn = cust.Year()
	}
	lnb := LeapmB(cust.Year() + 1)
	for i := 2; i < 14; i++ {
		if (cust.Equal(mtarr[i]) || cust.After(mtarr[i])) && cust.Before(mtarr[i+1]) {
			lunarM = i - 1
			if lnb == true && lunarM >= 12 { //2033
				lunarM -= 1
			}
			break
		}
	}
	if leapmb == true && (cust.Equal(leapmT) || cust.After(leapmT)) {
		lunarM -= 1
	}
	return lunarYn, lunarM
}
func findLunarMN(allshuo []float64) []time.Time {
	//上一年十一月到本年十月
	var mt []time.Time
	for i := 0; i < 12; i++ {
		jd := allshuo[i]
		st := JdToLocalTime(jd)
		st = time.Date(st.Year(), st.Month(), st.Day(), 0, 0, 0, 0, time.Local)
		lm := i + 11
		if lm > 12 {
			lm -= 12
		}
		mt = append(mt, st)
	}
	//本年十一月 十二月
	for i := 12; i <= 14; i++ {
		jd := allshuo[i]
		st := JdToLocalTime(jd)
		st = time.Date(st.Year(), st.Month(), st.Day(), 0, 0, 0, 0, time.Local)
		mt = append(mt, st)
	}
	return mt
}

//农历闰月数字 闰月所在朔的阳历时间戳 朔望时间戳数组的索引值
func findLeapmN(leapmb bool, zqarrt []time.Time, swarrt []*ShuoWangT) (int, time.Time) {
	var leapMonthN int       //农历闰月数字 这里是几就是闰几月
	var leapMonthT time.Time //闰月朔(农历初一)的阳历时间戳
	//这里需要cust 如果cust 在冬至时间戳之后 2033年情况需要另行判断
	if leapmb == true {
	I:
		for i := 0; i < len(zqarrt); i++ {
			zqt := zqarrt[i]
			zqt = time.Date(zqt.Year(), zqt.Month(), zqt.Day(), 0, 0, 0, 0, time.Local)
			for j := 0; j < len(swarrt)-1; j++ {
				if i == j {
					shuot0 := swarrt[j].ShuoT //本月初一
					shuot0 = time.Date(shuot0.Year(), shuot0.Month(), shuot0.Day(), 0, 0, 0, 0, time.Local)
					shuot1 := swarrt[j+1].ShuoT //下月初一
					shuot1 = time.Date(shuot1.Year(), shuot1.Month(), shuot1.Day(), 0, 0, 0, 0, time.Local)

					if (zqt.After(shuot0)) && (zqt.Equal(shuot1) || zqt.After(shuot1)) {
						leapMonthN = j - 2
						if leapMonthN < 0 {
							leapMonthN += 12
						}
						leapMonthT = shuot0
						break I
					}
					break
				}
			}
		}
		if leapMonthN == 0 { //2033 0-->11
			leapMonthN = 11
		}
	}
	return leapMonthN, leapMonthT
}
func Alias(m int) string {
	var index int
	for i := 0; i < len(Ymc); i++ {
		if i == m-1 {
			index = i
			break
		}
	}
	return Ymc[index]
}
