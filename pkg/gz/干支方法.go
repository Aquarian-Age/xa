package gz

import (
	"fmt"
	"github.com/Aquarian-Age/xa/pkg/pub"
	"github.com/starainrt/astro/calendar"
	"github.com/starainrt/astro/moon"
	"strings"
	"time"
)

var (
	xunKongMap = map[string]string{"甲子": "戌亥", "甲戌": "申酉", "甲申": "午未", "甲午": "辰巳", "甲辰": "寅卯", "甲寅": "子丑"}
	naYinMap   = map[string]string{
		"甲子": "海中金", "乙丑": "海中金",
		"丙寅": "炉中火", "丁卯": "炉中火",
		"戊辰": "大林木", "己巳": "大林木",
		"庚午": "路旁土", "辛未": "路旁土",
		"壬申": "剑锋金", "癸酉": "剑锋金",

		"甲戌": "山头火", "乙亥": "山头火",
		"丙子": "涧下水", "丁丑": "涧下水",
		"戊寅": "城头土", "己卯": "城头土",
		"庚辰": "白蜡金", "辛巳": "白蜡金",
		"壬午": "杨柳木", "癸未": "杨柳木",

		"甲申": "泉中水", "乙酉": "泉中水", //己酉-->乙酉
		"丙戌": "屋上土", "丁亥": "屋上土",
		"戊子": "霹雳火", "己丑": "霹雳火",
		"庚寅": "松柏木", "辛卯": "松柏木",
		"壬辰": "长流水", "癸巳": "长流水",

		"甲午": "沙中金", "乙未": "沙中金",
		"丙申": "山下火", "丁酉": "山下火",
		"戊戌": "平地木", "己亥": "平地木",
		"庚子": "壁上土", "辛丑": "壁上土",
		"壬寅": "金箔金", "癸卯": "金箔金",

		"甲辰": "覆灯火", "乙巳": "覆灯火", //己巳-->乙巳
		"丙午": "天河水", "丁未": "天河水",
		"戊申": "大鄢土", "己酉": "大鄢土", //yān
		"庚戌": "钗钏金", "辛亥": "钗钏金",
		"壬子": "桑柘木", "癸丑": "桑柘木",

		"甲寅": "大溪水", "乙卯": "大溪水",
		"丙辰": "沙中土", "丁巳": "沙中土",
		"戊午": "天上火", "己未": "天上火",
		"庚申": "石榴木", "辛酉": "石榴木",
		"壬戌": "大海水", "癸亥": "大海水",
	}
)

// NewGan 干
func (obj *GanZhi) NewGan(gan string) Gan {
	return Gan(gan)
}

// NewZhi 支
func (obj *GanZhi) NewZhi(zhi string) Zhi {
	return Zhi(zhi)
}

// DiSiHu 地四户(除 定 危 开)
func (obj *GanZhi) DiSiHu() *DiSiHuStruct {
	return DiSiHu(obj.Hgz)
}

// DiSiHuZhi DiSiHu 地四户 危 除 定 开 对应的地支
//月建加时支上 危除定开下临方既是
//建除满平一顺流,定执破危相接去,成收开闭掌中周,除定危开为四户,此方有难来逃避
func (obj *GanZhi) DiSiHuZhi() (string, string, string, string) {
	hz := pub.GetZhiS(obj.Hgz) //时支
	zhis := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	jcs := []string{"建", "除", "满", "平", "定", "执", "破", "危", "成", "收", "开", "闭"}
	hourArr := pub.SortArr(hz, zhis)

	var a, b, c, d string
	for i := 0; i < len(jcs); i++ {
		if strings.EqualFold(jcs[i], "危") {
			a = hourArr[i]
		}
		if strings.EqualFold(jcs[i], "除") {
			b = hourArr[i]
		}
		if strings.EqualFold(jcs[i], "定") {
			c = hourArr[i]
		}
		if strings.EqualFold(jcs[i], "开") {
			d = hourArr[i]
		}
	}
	return a, b, c, d
}

func (obj *GanZhi) XunShou(xgz string) (string, []string) {
	return XunShou(xgz)
}

// XunShou 旬首
func XunShou(xgz string) (string, []string) {
	var s string
	var arr []string
	switch xgz {
	case "甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉":
		s = "甲子"
		arr = []string{"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉"}
	case "甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未":
		s = "甲戌"
		arr = []string{"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未"}
	case "甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳":
		s = "甲申"
		arr = []string{"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳"}
	case "甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯":
		s = "甲午"
		arr = []string{"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯"}
	case "甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑":
		s = "甲辰"
		arr = []string{"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑"}
	case "甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥":
		s = "甲寅"
		arr = []string{"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"}
	}
	return s, arr
}

// XunKong 旬空
func XunKong(xgz string) string {
	xunShous, _ := XunShou(xgz)
	return xunKongMap[xunShous]
}

// FuTou 干支的符头--甲/己(默认为日干支)
func (obj *GanZhi) FuTou(dgz string) string {
	if dgz != "" {
		return fuTou(dgz)
	}
	return fuTou(obj.Dgz)
}
func fuTou(xgz string) string {
	var zhis string
	switch xgz {
	case "甲子", "乙丑", "丙寅", "丁卯", "戊辰":
		zhis = "子"
	case "己巳", "庚午", "辛未", "壬申", "癸酉":
		zhis = "巳"
	case "甲戌", "乙亥", "丙子", "丁丑", "戊寅":
		zhis = "戌"
	case "己卯", "庚辰", "辛巳", "壬午", "癸未":
		zhis = "卯"
	case "甲申", "乙酉", "丙戌", "丁亥", "戊子":
		zhis = "申"
	case "己丑", "庚寅", "辛卯", "壬辰", "癸巳":
		zhis = "丑"
	case "甲午", "乙未", "丙申", "丁酉", "戊戌":
		zhis = "午"
	case "己亥", "庚子", "辛丑", "壬寅", "癸卯":
		zhis = "亥"
	case "甲辰", "乙巳", "丙午", "丁未", "戊申":
		zhis = "辰"
	case "己酉", "庚戌", "辛亥", "壬子", "癸丑":
		zhis = "酉"
	case "甲寅", "乙卯", "丙辰", "丁巳", "戊午":
		zhis = "寅"
	case "己未", "庚申", "辛酉", "壬戌", "癸亥":
		zhis = "未"
	}
	return zhis
}

// GetLunar 返回阴历月日　月相
func (obj *GanZhi) GetLunar() (string, string) {
	_, _, _, moons := calendar.Lunar(obj.year, obj.month, obj.day)
	tx := time.Date(obj.year, time.Month(obj.month), obj.day, obj.hour, obj.min, 0, 0, time.Local)
	//_, _, _, moons := calendar.SolarToLunar(tx)
	moons = fmt.Sprintf("阴历: %s", moons)
	phase := moon.Phase(tx)
	yueXiang := fmt.Sprintf("月相: %5f", phase)
	return moons, yueXiang
}

// GetNaYinString 纳因(年-月-日-时)
func (obj *GanZhi) GetNaYinString() string {
	return GetNaYinString(obj.Ygz, obj.Mgz, obj.Dgz, obj.Hgz)
}

// NaYin 干支纳因
func (obj *GanZhi) NaYin(gzx string) string {
	return naYinMap[gzx]
}

// NaYin 指定干支的纳因 传入干支 返回对应纳因
func NaYin(gzx string) string {
	return naYinMap[gzx]
}

// GetNaYinString 年月日时 的干支纳音
func GetNaYinString(gzx ...string) string {
	var arr []string
	for i := 0; i < len(gzx); i++ {
		s := NaYin(gzx[i])
		arr = append(arr, s)
	}
	return strings.Join(arr, "-")
}

//日建除
//func (obj *GanZhi) RiJianChu() string {
//	return GetRiJianChu(obj.Mgz, obj.Dgz)
//}

// JianChu 日建除
func (obj *GanZhi) JianChu() string {
	return JianChu(obj.Mgz, obj.Dgz)
}

////日黄黑
//func (obj *GanZhi) RiHuangHei() string {
//	return GetRiHuangHei(obj.Mgz, obj.Dgz)
//}

// RiHuangHei1 日黄黑
func (obj *GanZhi) RiHuangHei1() string {
	return HuangHei(obj.Mgz, obj.Dgz)
}

////时黄黑
//func (obj *GanZhi) ShiHuangHei() string {
//	return GetRiHuangHei(obj.Dgz, obj.Hgz)
//}

// ShiHuangHei1 时黄黑
func (obj *GanZhi) ShiHuangHei1() string {
	return HuangHei(obj.Dgz, obj.Hgz)
}

// RiQin 日禽
func (obj *GanZhi) RiQin(weekN int) string {
	return GetRiQin(weekN, obj.Dgz)
}

// YueJiangStruct 月将
func (obj GanZhi) YueJiangStruct() *YJ {
	return NewYueJiang(obj.year, obj.month, obj.day, obj.Mgz)
}

// YueJiang 月将
//返回月将对应的地支 月将对应的神将名称 月将所对应的中气时间戳/中气名称
func (obj *GanZhi) YueJiang() (string, string, time.Time, string) {
	zhis := pub.GetZhiS(obj.Mgz)
	return yueJiang(obj.year, obj.month, obj.day, zhis)
}

// GuiRen 贵人诀 默认传入年干支
func (obj *GanZhi) GuiRen(xgz string) (string, string) {
	return GuiRen(xgz)
}

// GuiRenDay 贵人诀　日干支
//func (obj *GanZhi) GuiRenDay() (string, string) {
//	return GuiRenJue(obj.Dgz)
//}

// JieQi  当前节气时间 精确到小时
func (obj *GanZhi) JieQi() string {
	year := obj.year
	arr := jq24(year)
	arr1 := jq24(year + 1)
	arr = append(arr, arr1[1:3]...)
	Jmc = append(Jmc, Jmc[1:3]...)

	var jqs string //当前时间节气
	ct := time.Date(obj.year, time.Month(obj.month), obj.day, obj.hour, 0, 0, 0, time.Local)
	for i := 0; i < len(arr); i++ {
		xt := arr[i]
		xth := time.Date(xt.Year(), xt.Month(), xt.Day(), xt.Hour(), 0, 0, 0, time.Local)
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

// Jq24 24节气数组
func (obj *GanZhi) Jq24() []string {
	year := obj.year
	arr := jq24(year)
	arr1 := jq24(year + 1)
	arr = append(arr, arr1[1:3]...)
	Jmc = append(Jmc, Jmc[1:3]...)

	var tmp []string
	for i := 0; i < len(arr); i++ {
		x := Jmc[i] + ": " + arr[i].Format("2006-01-02 :15:04:05")
		tmp = append(tmp, x)
	}
	return tmp[:len(tmp)-24]
}
func (obj *GanZhi) Jq24T() []time.Time {
	return GetJq24(obj.year)
}
func GetJq24(year int) []time.Time {
	return jq24(year)
}
