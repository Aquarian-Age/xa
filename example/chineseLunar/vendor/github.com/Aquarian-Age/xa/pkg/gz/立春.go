package gz

import "time"

//立春修正
func fixLiChun(year int, cust time.Time) (bool, time.Time) {
	lct, _, _ := getJie12T(year)
	//节气精确到日
	lct = time.Date(lct.Year(), lct.Month(), lct.Day(), 0, 0, 0, 0, time.Local)
	cust = time.Date(cust.Year(), cust.Month(), cust.Day(), 0, 0, 0, 0, time.Local)
	var b bool
	if cust.Equal(lct) || cust.After(lct) {
		b = true //当前时间在立春之后
	} else {
		b = false //当前时间在立春之前
	}
	return b, lct
}
