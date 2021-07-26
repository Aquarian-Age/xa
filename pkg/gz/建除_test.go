/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package gz

import (
	"fmt"
	"strings"
	"testing"
)

func TestJianChu(t *testing.T) {
	arr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	arrx := [][]string{
		{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}, //建
		{"丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子"},
		{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"},
		{"卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅"},
		{"辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯"},
		{"巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰"},
		{"午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳"},
		{"未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午"},
		{"申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未"},
		{"酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申"},
		{"戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉"}, //开
		{"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"}, //闭
	}
	want := [][]string{
		{"建", "除", "满", "平", "定", "执", "破", "危", "成", "收", "开", "闭"},
		{"闭", "建", "除", "满", "平", "定", "执", "破", "危", "成", "收", "开"},
		{"开", "闭", "建", "除", "满", "平", "定", "执", "破", "危", "成", "收"},
		{"收", "开", "闭", "建", "除", "满", "平", "定", "执", "破", "危", "成"},
		{"成", "收", "开", "闭", "建", "除", "满", "平", "定", "执", "破", "危"},
		{"危", "成", "收", "开", "闭", "建", "除", "满", "平", "定", "执", "破"},
		{"破", "危", "成", "收", "开", "闭", "建", "除", "满", "平", "定", "执"},
		{"执", "破", "危", "成", "收", "开", "闭", "建", "除", "满", "平", "定"},
		{"定", "执", "破", "危", "成", "收", "开", "闭", "建", "除", "满", "平"},
		{"平", "定", "执", "破", "危", "成", "收", "开", "闭", "建", "除", "满"},
		{"满", "平", "定", "执", "破", "危", "成", "收", "开", "闭", "建", "除"},
		{"除", "满", "平", "定", "执", "破", "危", "成", "收", "开", "闭", "建"},
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arrx); j++ {
			jcs := JianChu(arr[i], arr[j])
			if !strings.EqualFold(jcs, want[i][j]) {
				t.Errorf("func JianChu(%s %s)=%s want:%s", arr[i], arr[j], jcs, want[i][j])
			}
		}

	}
}
func TestGetRiJianChu(t *testing.T) {
	jz60 := GetJzArr()
	for i := 0; i < len(jz60); i++ {
		mgz := "甲寅"
		mgz = "乙卯"
		dgz := jz60[i]
		s := GetRiJianChu(mgz, dgz)
		fmt.Printf("%s-%s %s\n", mgz, dgz, s)
	}

}
