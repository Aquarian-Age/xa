/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package gz

import (
	"strings"
	"time"
)

//月将(太阳过宫)
type YJ struct {
	YueJiang string `json:"yue_jiang"`
	Name     string `json:"name"`
	T        string `json:"t"`
}

func GetYueJiang(y, m, d int, mz string) *YJ {
	yj, name, t := yueJiang(y, m, d, mz)
	ts := t.Format("2006-01-02")
	return &YJ{
		yj,
		name,
		ts,
	}
}

//传入阳历时间(年　月　日)　月支 返回月将对应的地支 月将对应的神将名称 月将所对应的中气时间戳
func yueJiang(year, month, day int, mgz string) (string, string, time.Time) {
	cust := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local) //精确到日
	_, _, zqArrT := getJie12T(year)

	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子"}             //月支从上年冬月开始
	tyj := []string{"丑", "子", "亥", "戌", "酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑"}             //天月将的地支 从子月到子月
	sj := []string{"大吉", "神后", "登明", "河魁", "从魁", "传送", "小吉", "胜光", "太乙", "天罡", "太冲", "功曹", "大吉"} //从子月到子月

	var zqt time.Time //中气时间戳　精确到日
	var yjZhi string  //月将的地支
	var yjName string //神将
	for i := 1; i < len(zhi); i++ {
		if strings.EqualFold(mgz, zhi[i]) {
			zqt = zqArrT[i]
			zqt = time.Date(zqt.Year(), zqt.Month(), zqt.Day(), 0, 0, 0, 0, time.Local)
			if cust.Equal(zqt) || cust.After(zqt) {
				yjZhi = tyj[i]
				yjName = sj[i]
				break
			} else {
				index := i - 1
				if index < 0 {
					index = 13
				}
				yjZhi = tyj[index]
				yjName = sj[index]
				zqt = zqArrT[index]
				break
			}
		}
	}

	return yjZhi, yjName, zqt
}
