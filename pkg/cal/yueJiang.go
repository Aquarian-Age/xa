package cal

import (
	"strings"
	"time"
)

//二十四节气名称
var JMC = []string{
	"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
	"春分", "清明", "谷雨", "立夏", "小满", "芒种",
	"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
	"秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
}

//从上年冬至到下年立春的节气名称　节气时间戳
// 冬至 小寒 大寒 立春 雨水 惊蛰 春分 清明 谷雨 立夏 小满 芒种
//夏至 小暑 大暑 立秋 处暑 白露 秋分 寒露 霜降 立冬 小雪 大雪
//冬至 小寒 大寒 立春
type YueJiangJQ struct {
	Name []string    //节气名称数组
	Time []time.Time //节气时间数组 精确到分钟秒钟
}

//月将名称 月将对应的地支 十二宫星
type YueJiang struct {
	Name  string `json:"name"`   //月将
	Zhi   string `json:"zhi"`    //月将对应的地支
	Star  string `json:"star"`   //黄道十二宫星名称
	JieQi string `json:"jie_qi"` //节气
}

//节气时间数组　精确到分钟
//上一年冬至(含)开始 0:冬至　1:小寒　2:大寒...
func YueJiangJQT(y int) []time.Time { //y:农历年数字
	var j24t []time.Time
	data := Data(y)
	for i := 1; i <= 25; i++ { //<=25冬至到冬至
		hs := data[0] + data[i]  //UTC合朔
		hs8 := JdToLocalTime(hs) //CST+8
		j24t = append(j24t, hs8)
	}

	var j24t1 []time.Time
	data1 := Data(y + 1)
	for j := 1; j <= 25; j++ {
		hs := data1[0] + data1[j] //UTC合朔
		hs8 := JdToLocalTime(hs)  //CST+8
		j24t1 = append(j24t1, hs8)
	}

	//去重
	var kx int
I:
	for k := 0; k < len(j24t); k++ {
		xk := j24t[k]
		for k1 := 0; k1 < len(j24t1); k1++ {
			xk1 := j24t1[k1]
			if b := xk.Equal(xk1); b == true {
				kx = k1
				break I
			}
		}
	}
	//上一年冬至(含)开始 0:冬至　1:小寒　2:大寒...
	alljqt := append(j24t, j24t1[kx+1:]...)
	return alljqt
}

//上一年冬至到下一年立春的节气
func NewYueJiangJQ(jqt []time.Time) *YueJiangJQ {
	jq := new(YueJiangJQ)
	var jqmc []string     //节气名称
	var jqmct []time.Time //节气时间戳
	//i=28到下一年立春
	for i := 0; i < 28; i++ {
		x := i
		if x > 23 {
			x = i - 24
		}
		jqmc = append(jqmc, JMC[x])
		jqmct = append(jqmct, jqt[i])
	}
	jq = &YueJiangJQ{
		Name: jqmc,
		Time: jqmct,
	}
	return jq
}

//全年24节气 文本自动换行
func (jq *YueJiangJQ) JQ24() string {
	format := "2006-01-02 15:04:05"
	var jqArr []string
	for i := 3; i <= 26; i++ {
		jqx := jq.Name[i] + ":" + jq.Time[i].Format(format)
		jqArr = append(jqArr, jqx)
	}
	jqs := strings.Join(jqArr, "\n")
	return jqs
}
