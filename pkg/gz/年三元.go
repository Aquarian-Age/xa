package gz

// GetYuanS  当前年份的x元
func GetYuanS(year int) string {
	n := yearYuanNumber(year)
	switch n {
	case 0:
		return "下元"
	case 1:
		return "上元"
	case 2:
		return "中元"
	}

	return ""
}

//年奇门三元 有效时间范围(1601~3498)
//1:上元 2:中元 0:下元
func yearYuanNumber(year int) int {
	var lastYearOfJiaZi int //上一个甲子年
	//找到当前年份的上一个甲子年
	for i := year; i <= 9999; i-- {
		if i%60 == 4 { //甲子年余数都是4
			lastYearOfJiaZi = i
			break
		}
	}
	//计算上一个甲子年的三元数字
	n := (lastYearOfJiaZi / 60) % 3
	//当前年份的甲子数
	N := n + 1
	if N > 2 {
		N -= 3
	}
	return n
}
