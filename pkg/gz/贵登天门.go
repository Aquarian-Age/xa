/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package gz

import "strings"

//传入月将　日时干支 返回昼夜等天门时辰
func TianMen(yj, dgz string) (string, string) {
	return dtm(yj, dgz)
}

//k:月将
//v十干0:甲 1:乙
func dtm(yj, dgz string) (string, string) {

	tmap := map[string][]struct{ dan, xi string }{
		"亥": { //月将
			{"卯", "酉"}, //甲日
			{"", "戌"},
			{"", "亥"}, //丙日
			{"", "丑"}, //丁日
			{"酉", "卯"},
			{"", "寅"}, //己日
			{"酉", "卯"},
			{"申", ""}, //辛日
			{"未", ""},
			{"巳", ""}, //癸日
		},
		"戌": {
			{"", ""}, //甲日
			{"", "酉"},
			{"", "戌"}, //丙日
			{"", "子"},
			{"申", "寅"}, //戊日
			{"酉", "丑"},
			{"申", "寅"}, //庚日
			{"未", "卯"},
			{"午", ""}, //壬日
			{"辰", ""},
		},
		"酉": {
			{"", ""}, //甲日
			{"", ""},
			{"", ""}, //丙日
			{"酉", "亥"},
			{"未", "丑"}, //戊日
			{"申", "子"},
			{"未", "丑"}, //庚日
			{"午", "寅"},
			{"巳", ""}, //壬日
			{"卯", ""},
		},
		"申": {
			{"", ""}, //甲日
			{"", ""},
			{"戌", ""}, //丙日
			{"申", "戌"},
			{"午", "子"}, //戊日
			{"未", "亥"},
			{"午", "子"}, //庚日
			{"巳", "丑"},
			{"辰", "寅"}, //壬日
			{"寅", ""},
		},
		"未": {
			{"", ""}, //甲日
			{"戌", ""},
			{"酉", ""}, //丙日
			{"未", ""},
			{"巳", "亥"}, //戊日
			{"午", "戌"},
			{"巳", "亥"}, //庚日
			{"辰", "子"},
			{"卯", "丑"}, //壬日
			{"", ""},
		},
		"午": {
			{"", ""}, //甲日
			{"酉", ""},
			{"申", ""}, //丙日
			{"午", ""},
			{"辰", "戌"}, //戊日
			{"巳", ""},
			{"辰", "戌"}, //庚日
			{"卯", "亥"},
			{"寅", "子"}, //壬日
			{"", "寅"},
		},
		"巳": {
			{"酉", ""}, //甲日
			{"申", ""},
			{"未", ""}, //丙日
			{"巳", ""},
			{"卯", "酉"}, //戊日
			{"辰", ""},
			{"卯", "酉"}, //庚日
			{"", "戌"},
			{"", "亥"}, //壬日
			{"", "丑"},
		},
		"辰": {
			{"申", "寅"}, //甲日
			{"未", "卯"},
			{"午", ""}, //丙日
			{"辰", ""},
			{"", ""}, //戊日
			{"卯", ""},
			{"", ""}, //庚日
			{"", "酉"},
			{"", "戌"}, //壬日
			{"", "子"},
		},
		"卯": {
			{"未", "丑"}, //甲日
			{"午", "寅"},
			{"巳", "卯"}, //丙日
			{"卯", ""},
			{"", ""}, //戊日
			{"", ""},
			{"", ""}, //庚日
			{"", ""},
			{"", "酉"}, //壬日
			{"酉", "亥"},
		},
		"寅": {
			{"午", "子"}, //甲日
			{"巳", "丑"},
			{"辰", "寅"}, //丙日
			{"", ""},
			{"", ""}, //戊日
			{"", ""},
			{"", ""}, //庚日
			{"", ""},
			{"", "申"}, //壬日
			{"申", "戌"},
		},
		"丑": {
			{"巳", "亥"}, //甲日
			{"辰", "子"},
			{"", "丑"}, //丙日
			{"", "卯"},
			{"", ""}, //戊日
			{"", "辰"},
			{"", ""}, //庚日
			{"", ""},
			{"", ""}, //壬日
			{"未", "酉"},
		},
		"子": {
			{"辰", "戌"}, //甲日
			{"卯", "亥"},
			{"", "子"}, //丙日
			{"", "寅"},
			{"", ""}, //戊日
			{"", "卯"},
			{"", ""}, //庚日
			{"", ""},
			{"申", ""}, //壬日
			{"午", "申"},
		},
	}
	gan := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	var index int
	for i := 0; i < len(gan); i++ {
		if strings.ContainsAny(dgz, gan[i]) {
			index = i
			break
		}
	}

	var dan, mu string //贵登天门时
	for k, v := range tmap {
		if strings.EqualFold(yj, k) {
			for i := 0; i < len(v); i++ {
				if index == i {
					dan = v[i].dan
					mu = v[i].xi
					break
				}
			}
			break
		}
	}
	return dan, mu
}
