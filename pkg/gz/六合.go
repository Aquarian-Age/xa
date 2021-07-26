/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package gz

import "strings"

//返回xgz的地支合宫
func LiuHe(xgz string) string {
	hmap := map[string]string{
		"午": "未", "未": "午",
		"子": "丑", "丑": "子",
		"寅": "亥", "亥": "寅",
		"卯": "戌", "戌": "卯",
		"辰": "酉", "酉": "辰",
		"巳": "申", "申": "巳",
	}
	var s string
	for k, v := range hmap {
		if strings.Contains(xgz, k) {
			s = v
			break
		}
	}
	return s
}
