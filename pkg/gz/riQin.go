package gz

import "strings"

var StarName = []string{
	"角木蛟", "亢金龙", "氐土貉", "房日兔", "心月狐", "尾火虎", "箕水豹", //东方青龙
	"斗木獬", "牛金牛", "女土蝠", "虚日鼠", "危月燕", "室火猪", "壁水貐", //北方玄武
	"奎木狼", "娄金狗", "胃土雉", "昴日鸡", "毕月乌", "觜火猴", "参水猿", //西方白虎
	"井木犴", "鬼金羊", "柳土獐", "星日马", "张月鹿", "翼火蛇", "轸水蚓", //南方朱雀
}

//日禽 传入周数字　日干支
func GetRiQin(wd int, dgz string) string {
	index := RiQin(wd, dgz)
	return StarName[index]
}

//日禽索引值
func RiQin(wd int, dgz string) (wn int) {
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
		_w0 := 17            //周日值宿对应的基数
		wn = _w0 + 7*wd + wd //当前周数字对应的28宿索引
		if wn > 28 {
			wn = wn % 28
		}
	case "巳", "酉", "丑": //金局-->周日:房日兔值宿
		_w0 := 3
		wn = _w0 + 7*wd + wd
		if wn > 28 {
			wn = wn % 28
		}
	case "寅", "午", "戌": //火局-->周日:星日马值宿
		_w0 := 24
		wn = _w0 + 7*wd + wd
		if wn > 28 {
			wn = wn % 28
		}
	case "申", "子", "辰": //水局-->周日:虚日鼠值宿
		_w0 := 10
		wn = _w0 + 7*wd + wd
		if wn > 28 {
			wn = wn % 28
		}
	}
	return
}