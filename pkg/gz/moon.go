package gz

import (
	"fmt"
	"github.com/starainrt/astro/calendar"
	"github.com/starainrt/astro/moon"
	"time"
)

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

type ShuoWang struct {
	LastWang  time.Time `json:"last_wang"`
	NextWang  time.Time `json:"next_wang"`
	Shuo      time.Time `json:"shuo"`
	Wang      time.Time `json:"wang"`
	ShangXian time.Time `json:"shang_xian"`
	XiaXian   time.Time `json:"xia_xian"`
	LastShuo  time.Time `json:"last_shuo"`
	NextShuo  time.Time `json:"next_shuo"`
}

// Moons 朔望
func (obj *GanZhi) Moons() *ShuoWang {
	cst := time.FixedZone("CST", 8*3600)
	date := time.Date(obj.year, time.Month(obj.month), obj.day, obj.hour, 0, 0, 0, cst)

	nextShuoT := moon.NextShuoYue(date)          //下一朔月
	lastShuoT := moon.LastShuoYue(date)          //上一个朔月
	shuoT := moon.ClosestShuoYue(date)           //最近的朔月
	wangT := moon.ClosestWangYue(date)           // 最近的望月时间
	shangXianT := moon.ClosestShangXianYue(date) //最近的上弦月时间
	xianXianT := moon.ClosestXiaXianYue(date)    //最近的下弦月时间
	lastWangT := moon.LastWangYue(date)          //上一个望月
	nextWangT := moon.NextWangYue(date)          //下一个望月

	return &ShuoWang{
		lastWangT,
		nextWangT,
		shuoT,
		wangT,
		shangXianT,
		xianXianT,
		lastShuoT,
		nextShuoT,
	}
}

// FormatTime 朔望时间格式化为字符串
func (sw *ShuoWang) FormatTime(xt time.Time) string {
	return xt.Format("2006-01-02 15:04:05")
}
