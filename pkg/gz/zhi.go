package gz

import "math"

// AliasZhi 日支 1+mod(JD正午+1,12) 这里传入的JD是时间精确到日计算
func AliasZhi(jd float64) string {
	n := int(math.Ceil(math.Mod(jd+1, 12)))
	zhis := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	if n > 11 {
		n -= 12
	}
	return zhis[n]
}
