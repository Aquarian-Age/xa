package pub

import "strings"

//取天干 传干支 返回干
func GetGanS(gz string) string {
	gan := [11]string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	var g string
	for i := 0; i < len(gan); i++ {
		if strings.ContainsAny(gz, gan[i]) {
			g = gan[i]
			break
		}
	}
	return g
}

//取地支 传干支返回支
func GetZhiS(gz string) string {
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	var z string
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(gz, zhi[i]) {
			z = zhi[i]
			break
		}
	}
	return z
}

//顺序排地支 传入对应的地支 原地支数组 返回排序后的地支数组
func SortArr(zhi string, zhiArr []string) []string {
	for i := 0; i < len(zhiArr); i++ {
		if strings.EqualFold(zhi, zhiArr[i]) {
			head := zhiArr[:i]
			end := zhiArr[i:]
			zhiArr = append(end, head...)
			break
		}
	}
	return zhiArr
}
