package gz

import (
	"fmt"
	"github.com/starainrt/astro/basic"
	"github.com/starainrt/astro/calendar"
	"sort"
	"time"
)

// NewTMGanZhi 干支精确到分钟
func NewTMGanZhi(year, month, day, hour, min int) *GanZhi {
	ygz, mgz, jieQiArrT, jieQiNames, zhongQiArrT, zhongQiNames := TMMonthGanZhi(year, month, day, hour, min)
	dgz, dayGan := dayGanZhi(year, month, day)
	hgz := GetHourGZ(dayGan, hour)
	return &GanZhi{
		year:         year,
		month:        month,
		day:          day,
		hour:         hour,
		min:          min,
		Ygz:          ygz,
		Mgz:          mgz,
		Dgz:          dgz,
		Hgz:          hgz,
		JieQiArrT:    jieQiArrT,
		JieQiNames:   jieQiNames,
		ZhongQiArrT:  zhongQiArrT,
		ZhongQiNames: zhongQiNames,
	}
}

func TMFixLiChun(lct, cust time.Time) bool {
	var b bool
	if cust.Equal(lct) || cust.After(lct) {
		b = true //当前时间在立春之后
	} else {
		b = false //当前时间在立春之前
	}
	return b
}

// TMjieQi 本年节气数组 精确到分钟 0:冬至 1:小寒 2:大寒 3:立春 4:雨水...冬至 立春节气时间
func TMjieQi(year int) ([]time.Time, time.Time) {
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
		xt = time.Date(xt.Year(), xt.Month(), xt.Day(), xt.Hour(), xt.Minute(), 0, 0, time.Local) //精确到小时
		arr = append(arr, xt)
	}
	return arr, arr[3] //立春时间
}

// TMYearGanZhi 年干 年支
func TMYearGanZhi(year int, lcb bool) (string, string) {
	return yearGZ(year, lcb)
}

// TMMonthGanZhi
//返回 年干支 月干支 节气时间 节气名称 中气时间 中气名称
func TMMonthGanZhi(year, month, day, hour, min int) (string, string, []time.Time, []string, []time.Time, []string) {
	cust := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	arr, lct := TMjieQi(year)
	lcb := TMFixLiChun(lct, cust)
	yearGan, yearZhi := yearGZ(year, lcb)

	//小寒  立春  惊蛰  清明  立夏  芒种  小暑  立秋  白露  寒露  立冬  大雪
	var jqArr []time.Time //节气 0:上年小寒 1:立春
	var zqArr []time.Time //中气
	for i := 0; i < len(arr); i++ {
		if i%2 == 1 {
			jqArr = append(jqArr, arr[i])
		}
		if i%2 == 0 {
			zqArr = append(zqArr, arr[i])
		}
	}

	arr, _ = TMjieQi(year + 1)
	for i := 0; i < len(arr); i++ {
		if i%2 == 1 {
			jqArr = append(jqArr, arr[i])
		}
		if i%2 == 0 && i != 0 {
			zqArr = append(zqArr, arr[i])
		}
	}

	//节(len=24 不含中气)
	jieQiNames := []string{"小寒", "立春", "惊蛰", "清明", "立夏", "芒种", "小暑", "立秋", "白露", "寒露", "立冬", "大雪", "小寒", "立春"}
	zhongQiNames := []string{"冬至", "大寒", "雨水", "春分", "谷雨", "小满", "夏至", "大暑", "处暑", "秋分", "霜降", "小雪", "冬至", "大寒"}
	jieQiArrT := jqArr[:14]
	zhongQiArrT := zqArr[:14]

	var jqb bool //当前时间等于节气或者在节气之后
	var index int
	for i := 0; i < len(jqArr); i++ {
		if cust.Equal(jqArr[i]) || cust.After(jqArr[i]) {
			index = i
			jqb = true
		}
	}

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
	case true:
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
	return yearGZs, monthGZs, jieQiArrT, jieQiNames, zhongQiArrT, zhongQiNames
}

// TMJieQi 当前节气 精确到分钟
func (obj *GanZhi) TMJieQi() string {
	year := obj.year
	arr := jq24(year)
	arr1 := jq24(year + 1)
	arr = append(arr, arr1[1:3]...)
	Jmc = append(Jmc, Jmc[1:3]...)

	var jqs string //当前时间节气
	ct := time.Date(obj.year, time.Month(obj.month), obj.day, obj.hour, obj.min, 0, 0, time.Local)
	for i := 0; i < len(arr); i++ {
		xt := arr[i]
		xth := time.Date(xt.Year(), xt.Month(), xt.Day(), xt.Hour(), xt.Minute(), 0, 0, time.Local)
		if xth.Equal(ct) || xth.After(ct) {
			index := i - 1
			if index > 24 {
				index = i + 1
			}
			xts := arr[index].Format("2006-01-02 15:04:05")
			jqs = fmt.Sprintf("%s: %s", Jmc[index], xts)
			break
		}
	}
	return jqs
}
