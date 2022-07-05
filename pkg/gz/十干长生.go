package gz

import (
	"fmt"
	"strings"
)

var (
	jiaChangShengMap  = map[string]string{"丑": "冠带", "亥": "长生", "午": "死", "卯": "帝旺", "子": "沐浴", "寅": "临官", "巳": "病", "戌": "养", "未": "墓", "申": "绝", "辰": "衰", "酉": "胎"}
	yiChangShengMap   = map[string]string{"丑": "衰", "亥": "死", "午": "长生", "卯": "临官", "子": "病", "寅": "帝旺", "巳": "沐浴", "戌": "墓", "未": "养", "申": "胎", "辰": "冠带", "酉": "绝"}
	bingChangShengMap = map[string]string{"丑": "养", "亥": "绝", "午": "帝旺", "卯": "沐浴", "子": "胎", "寅": "长生", "巳": "临官", "戌": "墓", "未": "衰", "申": "病", "辰": "冠带", "酉": "死"}
	dingChangShengMap = map[string]string{"丑": "墓", "亥": "胎", "午": "临官", "卯": "病", "子": "绝", "寅": "死", "巳": "帝旺", "戌": "养", "未": "冠带", "申": "沐浴", "辰": "衰", "酉": "长生"}
	wuChangShengMap   = map[string]string{"丑": "养", "亥": "绝", "午": "帝旺", "卯": "沐浴", "子": "胎", "寅": "长生", "巳": "临官", "戌": "墓", "未": "衰", "申": "病", "辰": "冠带", "酉": "死"}
	jiChangShengMap   = map[string]string{"丑": "墓", "亥": "胎", "午": "临官", "卯": "病", "子": "绝", "寅": "死", "巳": "帝旺", "戌": "养", "未": "冠带", "申": "沐浴", "辰": "衰", "酉": "长生"}
	gengChangShengMap = map[string]string{"丑": "墓", "亥": "病", "午": "沐浴", "卯": "胎", "子": "死", "寅": "绝", "巳": "长生", "戌": "衰", "未": "冠带", "申": "临官", "辰": "养", "酉": "帝旺"}
	xinChangShengMap  = map[string]string{"丑": "养", "亥": "沐浴", "午": "病", "卯": "绝", "子": "长生", "寅": "胎", "巳": "死", "戌": "冠带", "未": "衰", "申": "帝旺", "辰": "墓", "酉": "临官"}
	renChangShengMap  = map[string]string{"丑": "衰", "亥": "临官", "午": "胎", "卯": "死", "子": "帝旺", "寅": "病", "巳": "绝", "戌": "冠带", "未": "养", "申": "长生", "辰": "墓", "酉": "沐浴"}
	guiChangShengMap  = map[string]string{"丑": "冠带", "亥": "帝旺", "午": "绝", "卯": "长生", "子": "临官", "寅": "沐浴", "巳": "胎", "戌": "衰", "未": "墓", "申": "死", "辰": "养", "酉": "病"}
)

// ChangShengZhiS 干 支 返回干在支的长生状态
func ChangShengZhiS(gan, zhi string) string {
	s := ChangSheng(gan, zhi)
	return fmt.Sprintf("%s在%s位:%s", gan, zhi, s)
}

// ChangShengMonthS 干在月支的长生状态
func ChangShengMonthS(gan, zhi string) string {
	s := ChangSheng(gan, zhi)
	return fmt.Sprintf("%s在%s月:%s", gan, zhi, s)
}

// GanZhiChangSheng 干支长生 返回干与支在十二长生的关系 比如甲子 甲在子位沐浴 则甲临沐浴
func GanZhiChangSheng(xgz string) string {
	gs := xgz[:3]
	zs := xgz[3:6]
	cs := ChangSheng(gs, zs)
	return gs + "临" + cs
}

// ChangSheng 十干长生 传入干 支 返回干在支的十二长生名称
func ChangSheng(gan, zhi string) string {
	var s string
	switch gan {
	case "甲":
		s = jiaChangShengMap[zhi]
	case "乙":
		s = yiChangShengMap[zhi]
	case "丙":
		s = bingChangShengMap[zhi]
	case "丁":
		s = dingChangShengMap[zhi]
	case "戊":
		s = wuChangShengMap[zhi]
	case "己":
		s = jiChangShengMap[zhi]
	case "庚":
		s = gengChangShengMap[zhi]
	case "辛":
		s = xinChangShengMap[zhi]
	case "壬":
		s = renChangShengMap[zhi]
	case "癸":
		s = guiChangShengMap[zhi]
	}
	return s
}

var localGn8ZhiMap = map[int]string{1: "子", 8: "丑", 3: "卯", 4: "巳", 9: "午", 2: "未", 7: "酉", 6: "亥"}
var localGn4ZhiMap = map[int]string{8: "寅", 4: "辰", 2: "申", 6: "戌"}

// ChangSheng8 八宫的十干长生 传入干 本宫数字(坎1 艮8 震3....)
func ChangSheng8(gan string, localGn int) string {
	name := ChangSheng(gan, localGn8ZhiMap[localGn])
	if name == "冠带" {
		name = "冠"
	}
	if name == "长生" {
		name = "生"
	}
	if name == "帝旺" {
		name = "旺"
	}
	if name == "沐浴" {
		name = "沐"
	}
	if name == "临官" {
		name = "临"
	}
	return name
}

// ChangSheng4 四维宫的十干长生 传入干 本宫数字( 艮8 巽4 坤2 乾6)
func ChangSheng4(gan string, localGn int) string {
	name := ChangSheng(gan, localGn4ZhiMap[localGn])
	if name == "冠带" {
		name = "冠"
	}
	if name == "长生" {
		name = "生"
	}
	if name == "帝旺" {
		name = "旺"
	}
	if name == "沐浴" {
		name = "沐"
	}
	if name == "临官" {
		name = "临"
	}
	return name
}

// ChangShengZhi 干在十二长生位置对应的地支
//传入天干 十二长生名称(长生 沐浴 冠带 临官 帝旺 衰 病 死 墓 绝 胎 养) 返回对应的地支
func ChangShengZhi(gan, name string) (zhis string) {
	arr := findArr(gan, changShengMap)
	index := findName(name, changShengNames)
	switch gan {
	case "甲":
		zhis = arr[index]
	case "乙":
		zhis = arr[index]
	case "丙":
		zhis = arr[index]
	case "丁":
		zhis = arr[index]
	case "戊":
		zhis = arr[index]
	case "己":
		zhis = arr[index]
	case "庚":
		zhis = arr[index]
	case "辛":
		zhis = arr[index]
	case "壬":
		zhis = arr[index]
	case "癸":
		zhis = arr[index]
	}
	return
}
func findArr(name string, xmap map[string][]string) []string {
	var array []string
	for k, arr := range xmap {
		if strings.EqualFold(name, k) {
			array = arr
			break
		}
	}
	return array
}
func findName(name string, arr []string) int {
	var index int
	for i := 0; i < len(arr); i++ {
		if strings.EqualFold(name, arr[i]) {
			index = i
			break
		}
	}
	return index
}
