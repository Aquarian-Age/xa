package gz

import (
	"fmt"
	"github.com/Aquarian-Age/xa/pkg/pub"
	"github.com/starainrt/astro/calendar"
	"github.com/starainrt/astro/moon"
	"time"
)

// FuTou 干支的符头--甲/己(默认为日干支)
func (obj *GanZhi) FuTou(dgz string) string {
	if dgz != "" {
		return fuTou(dgz)
	}
	return fuTou(obj.DGZ)
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
	tx := time.Date(obj.year, time.Month(obj.month), obj.day, obj.hour, 0, 0, 0, time.Local)
	moons = fmt.Sprintf("阴历: %s", moons)
	phase := moon.Phase(tx)
	yueXiang := fmt.Sprintf("月相: %5f", phase)
	return moons, yueXiang
}

// GetNaYin 年-月-日-时 纳因
func (obj *GanZhi) GetNaYin() string {
	return GetNaYin(obj.YGZ, obj.MGZ, obj.DGZ, obj.HGZ)
}

//日建除
//func (obj *GanZhi) RiJianChu() string {
//	return GetRiJianChu(obj.MGZ, obj.DGZ)
//}

// JianChu 日建除
func (obj *GanZhi) JianChu() string {
	return JianChu(obj.MGZ, obj.DGZ)
}

////日黄黑
//func (obj *GanZhi) RiHuangHei() string {
//	return GetRiHuangHei(obj.MGZ, obj.DGZ)
//}

// RiHuangHei1 日黄黑
func (obj *GanZhi) RiHuangHei1() string {
	return HuangHei(obj.MGZ, obj.DGZ)
}

////时黄黑
//func (obj *GanZhi) ShiHuangHei() string {
//	return GetRiHuangHei(obj.DGZ, obj.HGZ)
//}

// ShiHuangHei1 时黄黑
func (obj *GanZhi) ShiHuangHei1() string {
	return HuangHei(obj.DGZ, obj.HGZ)
}

// RiQin 日禽
func (obj *GanZhi) RiQin(weekN int) string {
	return GetRiQin(weekN, obj.DGZ)
}

// YueJiangStruct 月将
func (obj GanZhi) YueJiangStruct() *YJ {
	return NewYueJiang(obj.year, obj.month, obj.day, obj.MGZ)
}

// YueJiang 月将
//返回月将对应的地支 月将对应的神将名称 月将所对应的中气时间戳/中气名称
func (obj *GanZhi) YueJiang() (string, string, time.Time, string) {
	zhis := pub.GetZhiS(obj.MGZ)
	return yueJiang(obj.year, obj.month, obj.day, zhis)
}

// GuiRenYear 贵人诀 默认传入年干支
func (obj *GanZhi) GuiRenYear() (string, string) {
	return GuiRenJue(obj.YGZ)
}

// GuiRenDay 贵人诀　日干支
func (obj *GanZhi) GuiRenDay() (string, string) {
	return GuiRenJue(obj.DGZ)
}

// JieQi 当前节气名称:节气时间
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

// Jq24 24节气
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
