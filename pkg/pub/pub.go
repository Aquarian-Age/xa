package pub

import "strings"

//取地支 传干支返回支
func GetZhiS(gz string) string {
	zhi := [13]string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	var z string
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(gz, zhi[i]) {
			z = zhi[i]
			break
		}
	}
	return z
}
