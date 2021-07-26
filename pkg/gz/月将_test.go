/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package gz

import (
	"github.com/Aquarian-Age/xa/pkg/pub"
	"strings"
	"testing"
)

func TestGetYueJiang(t *testing.T) {
	want := []struct {
		yj, name, ts string
	}{
		{"丑", "大吉", "2020-12-21"},
		{"子", "神后", "2021-01-20"},
		{"亥", "登明", "2021-02-18"},
		{"戌", "河魁", "2021-03-20"},
		{"酉", "从魁", "2021-04-20"},
		{"申", "传送", "2021-05-21"},
		{"未", "小吉", "2021-06-21"},
		{"午", "胜光", "2021-07-22"},
		{"巳", "太乙", "2021-08-23"},
		{"辰", "天罡", "2021-09-23"},
		{"卯", "太冲", "2021-10-23"},
		{"寅", "功曹", "2021-11-22"},
		{"丑", "大吉", "2021-12-21"},
		{"子", "神后", "2022-01-20"},
		{"亥", "登明", "2022-02-19"},
		{"戌", "河魁", "2022-03-20"},
		{"酉", "从魁", "2022-04-20"},
		{"申", "传送", "2022-05-21"},
		{"未", "小吉", "2022-06-21"},
		{"午", "胜光", "2022-07-23"},
		{"巳", "太乙", "2022-08-23"},
		{"辰", "天罡", "2022-09-23"},
		{"卯", "太冲", "2022-10-23"},
		{"寅", "功曹", "2022-11-22"},
		{"丑", "大吉", "2022-12-22"},
	}
	_, zqt := getJieArr(2021)
	for i := 0; i < len(zqt); i++ {
		obj := NewGanZhi(zqt[i].Year(), int(zqt[i].Month()), zqt[i].Day(), zqt[i].Hour())
		yj := GetYueJiang(zqt[i].Year(), int(zqt[i].Month()), zqt[i].Day(), pub.GetZhiS(obj.MGZ))
		if !strings.EqualFold(yj.YueJiang, want[i].yj) && strings.EqualFold(yj.Name, want[i].name) && strings.EqualFold(yj.T, want[i].ts) {
			t.Errorf("func GetYueJiang(%d %d %d %s)=%v want %v\n",
				zqt[i].Year(), int(zqt[i].Month()), zqt[i].Day(), pub.GetZhiS(obj.MGZ), yj, want[i])
		}
	}
}
