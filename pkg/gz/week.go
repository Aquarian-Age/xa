package gz

import "math"

// WeekNumber 星期数(7是周日)  这里传入的JD是时间精确到日计算
func WeekNumber(jd float64) int {
	n := int(math.Ceil(math.Mod(jd+1, 7)))
	return n
}
