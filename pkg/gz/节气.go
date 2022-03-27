package gz

import (
	"github.com/starainrt/astro/basic"
	"github.com/starainrt/astro/calendar"
	"sort"
	"time"
)

// 冬至-冬至 len=25 索引单数为节 1,3,5,7,9,11,13,15,17,19,21,23 偶数为气
func jq24(year int) []time.Time {
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
		arr = append(arr, xt)
	}
	return arr
}

//传入阳历年数字 返回本年立春阳历时间戳 12节时间戳数组(上一年冬至到本年冬至)
//获取本年立春时间戳
func getJie12T(year int) (time.Time, []time.Time, []time.Time) {
	year -= 1 //k:1-->上一年冬至时间 k:25-->本年冬至时间 k:4--本年立春
	jq := basic.GetOneYearJQ(year)
	var keys []int
	for k := range jq {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	//k:1上一年冬至...k4:立春... k:25本年冬至
	/*
		"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
		"春分", "清明", "谷雨", "立夏", "小满", "芒种",
		"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
		"秋分", "寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
	*/
	var zqArr []time.Time  //中气(0:上一年冬至 到本年冬至)
	var jieArr []time.Time //12节
	var lct time.Time
	for _, v := range keys {
		if v%2 == 1 { //中气
			zqArr = append(zqArr, calendar.JDE2Date(jq[v]))
		}
		if v%2 == 0 { //节
			jieArr = append(jieArr, calendar.JDE2Date(jq[v]))
		}
		if v == 4 {
			lct = calendar.JDE2Date(jq[v])
		}
	}
	//12中气
	// 冬至  大寒  雨水  春分  谷雨  小满  夏至  大暑  处暑  秋分  霜降  小雪 冬至
	//12节
	// 小寒  立春  惊蛰  清明  立夏  芒种  小暑  立秋  白露  寒露  立冬  大雪
	//排序后对应的k值 2 4 6 8 10 12 14 16 18 20 22 24
	return lct, jieArr, zqArr
}

//正月立春节 二月惊蛰节 三月清明节 四月立夏节 五月忙钟节 六月小暑节
//七月立秋节 八月白露节 九月寒露节 十月立东节 冬月大雪节 腊月小寒节
//12节  0:上一年小寒 1今年立春...11大雪 12:本年小寒 13:下年立春
// 小寒  立春  惊蛰  清明  立夏  芒种  小暑  立秋  白露  寒露  立冬  大雪
//12中气 0:上一年冬至　12:本年冬至时间戳 13:下一年大寒
// 冬至  大寒  雨水  春分  谷雨  小满  夏至  大暑  处暑  秋分  霜降  小雪
//2年的节气和中气时间戳　时间精确到秒
//上一年小寒到下一年节气的时间戳数组 len=24 上一年冬至到本年冬至中气时间戳数组 len=25
func getJieArr(year int) ([]time.Time, []time.Time) {
	_, j12arr, zq1Arr := getJie12T(year)
	_, j4arr, zq2Arr := getJie12T(year + 1)
	var zqArrT []time.Time //12中气
	zq2Arr = zq2Arr[1:]    //去掉数组中本年冬至重复的时间戳
	zqArrT = append(zqArrT, zq1Arr...)
	zqArrT = append(zqArrT, zq2Arr...)
	var arrT []time.Time //12节气
	arrT = append(arrT, j12arr...)
	arrT = append(arrT, j4arr...)
	return arrT, zqArrT
}

//true节气之后 false节气之前 节气计算精确到日
func findJie(cust time.Time, jarrT []time.Time) (bool, int) {
	cust = time.Date(cust.Year(), cust.Month(), cust.Day(), 0, 0, 0, 0, time.Local)
	var b bool
	var index int
	for i := 0; i < len(jarrT)-1; i++ {
		j0 := jarrT[i]
		j1 := jarrT[i+1]
		j0 = time.Date(j0.Year(), j0.Month(), j0.Day(), 0, 0, 0, 0, time.Local) //精确到日
		j1 = time.Date(j1.Year(), j1.Month(), j1.Day(), 0, 0, 0, 0, time.Local)
		if (cust.Equal(j0) || cust.After(j0)) && cust.Before(j1) {
			index = i
			b = true //节气之后
			break
		}
	}
	return b, index
}
