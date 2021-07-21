package cal

import "strings"

//周几的值宿=当前局的基数+7*当前周几+当前周几 如果大于28取余数即为28宿索引号
// 当日值宿名称 传入当前周数字(0周日) dayz:当日地支　返回 值宿名称　七煞日( b=true七煞 false非七煞)　七煞名称
func Star(weekN int, dgz string) (name string) {
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	var dayz string
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(dgz, zhi[i]) {
			dayz = zhi[i]
			break
		}
	}
	switch dayz {
	case "亥", "卯", "未": //木局-->周日:昴日鸡值宿
		_w0 := 17                   //周日值宿对应的基数
		wn := _w0 + 7*weekN + weekN //当前周数字对应的28宿索引
		if wn > 28 {
			wn = wn % 28
		}
		name = XingSu28[wn] //当日值宿名称
	case "巳", "酉", "丑": //金局-->周日:房日兔值宿
		_w0 := 3
		wn := _w0 + 7*weekN + weekN
		if wn > 28 {
			wn = wn % 28
		}
		name = XingSu28[wn]
	case "寅", "午", "戌": //火局-->周日:星日马值宿
		_w0 := 24
		wn := _w0 + 7*weekN + weekN
		if wn > 28 {
			wn = wn % 28
		}
		name = XingSu28[wn]
	case "申", "子", "辰": //水局-->周日:虚日鼠值宿
		_w0 := 10
		wn := _w0 + 7*weekN + weekN
		if wn > 28 {
			wn = wn % 28
		}
		name = XingSu28[wn]
	}
	return
}
