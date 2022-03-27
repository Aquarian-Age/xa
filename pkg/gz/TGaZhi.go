package gz

import (
	"github.com/starainrt/astro/basic"
	"github.com/starainrt/astro/calendar"
	"sort"
	"time"
)

// NewTGanZhi 月干支精确到时辰
func NewTGanZhi(year, month, day, hour int) *GanZhi {
	ygz, mgz := TMonthGanZhi(year, month, day, hour)
	dgz, dayGan := dayGanZhi(year, month, day)
	hgz := GetHourGZ(dayGan, hour)
	return &GanZhi{
		year:  year,
		month: month,
		day:   day,
		hour:  hour,
		Ygz:   ygz,
		Mgz:   mgz,
		Dgz:   dgz,
		Hgz:   hgz,
	}
}

var ()

func TFixLiChun(lct, cust time.Time) bool {
	var b bool
	if cust.Equal(lct) || cust.After(lct) {
		b = true //当前时间在立春之后
	} else {
		b = false //当前时间在立春之前
	}
	return b
}
func TjieQi(year int) ([]time.Time, time.Time) {
	year -= 1 //k:1-->上一年冬至时间 k:25-->本年冬至时间 k:4--本年立春
	jq := basic.GetOneYearJQ(year)
	var keys []int
	for k := range jq {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var arr []time.Time
	for _, v := range keys {
		xt := calendar.JDE2Date(jq[v])
		xt = time.Date(xt.Year(), xt.Month(), xt.Day(), xt.Hour(), 0, 0, 0, time.Local) //精确到小时
		arr = append(arr, xt)
	}
	return arr, arr[3] //立春时间
}

// TYearGanZhi 年干 年支
func TYearGanZhi(year int, lcb bool) (string, string) {
	return yearGZ(year, lcb)
}

func TMonthGanZhi(year, month, day, hour int) (string, string) {
	cust := time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.Local)
	arr, lct := TjieQi(year)
	lcb := TFixLiChun(lct, cust)
	//fmt.Println("fix lichun: ", lct.String()[:19], lcb)
	yearGan, yearZhi := yearGZ(year, lcb)
	//fmt.Println("年干:", yearGan)

	//小寒  立春  惊蛰  清明  立夏  芒种  小暑  立秋  白露  寒露  立冬  大雪
	var jqArr []time.Time //0:上年小寒 1:立春
	for i := 0; i < len(arr); i++ {
		if i%2 == 1 {
			jqArr = append(jqArr, arr[i])
		}
	}
	//fmt.Println("jqArr: ", jqArr)
	arr, _ = TjieQi(year + 1)
	for i := 0; i < len(arr); i++ {
		if i%2 == 1 {
			jqArr = append(jqArr, arr[i])
		}
	}
	//fmt.Println(len(jqArr), jqArr)
	arrName := []string{"小寒", "立春", "惊蛰", "清明", "立夏", "芒种", "小暑", "立秋", "白露", "寒露", "立冬", "大雪"}
	arrName = append(arrName, arrName...)
	//fmt.Println(len(arrName), arrName)
	//
	var jqb bool //当前时间等于节气或者在节气之后
	//var jqt time.Time
	var index int
	for i := 0; i < len(jqArr); i++ {
		if cust.Equal(jqArr[i]) || cust.After(jqArr[i]) {
			index = i
			//jqt = jqArr[i]
			jqb = true
		}
	}
	//fmt.Println(year, month, day, hour)
	//fmt.Printf("jqb:%t jqt:%v index:%d\n", jqb, jqt.Format("2006-01-02 15:04:05"), index)

	jiaJiArr := []string{"丙子", "丁丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉", "甲戌", "乙亥", "丙子", "丁丑"}
	yiGengArr := []string{"戊子", "己丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未", "甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑"}
	bingXinArr := []string{"庚子", "辛丑", "庚寅", "辛卯", "壬辰", "癸巳", "甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑"}
	dingRenArr := []string{"壬子", "癸丑", "壬寅", "癸卯", "甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑"}
	wuGuiArr := []string{"甲子", "乙丑", "甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥", "甲子", "乙丑"}

	var monthGZs string
	switch jqb {
	case false:
		switch yearGan {
		case "甲", "己":
			monthGZs = jiaJiArr[index]
		case "乙", "庚":
			monthGZs = yiGengArr[index]
		case "丙", "辛":
			monthGZs = bingXinArr[index]
		case "丁", "壬":
			monthGZs = dingRenArr[index]
		case "戊", "癸":
			monthGZs = wuGuiArr[index]
		}
		//if index == 0 { //当前时间在上年(子月:阴历11月)
		//	fmt.Printf("%s年 月干支:%s\n", yearGan, dingRenArr[index])
		//	fmt.Printf("%s年 月干支:%s\n", yearGan, wuGuiArr[index])
		//}
	case true:
		//fmt.Printf("%s年 月干支:%s\n", yearGan, dingRenArr[index+1])
		//fmt.Printf("%s年 月干支:%s\n", yearGan, wuGuiArr[index+1])
		index += 1
		switch yearGan {
		case "甲", "己":
			monthGZs = jiaJiArr[index]
		case "乙", "庚":
			monthGZs = yiGengArr[index]
		case "丙", "辛":
			monthGZs = bingXinArr[index]
		case "丁", "壬":
			monthGZs = dingRenArr[index]
		case "戊", "癸":
			monthGZs = wuGuiArr[index]
		}
	}
	yearGZs := yearGan + yearZhi
	//fmt.Printf("阳历: %d-%d-%d %dH  %s年 月干支:%s\n", year, month, day, hour, yearGZs, monthGZs)
	return yearGZs, monthGZs
}
