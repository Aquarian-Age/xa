package cal

import (
	"database/sql"
	"fmt"
	"math"
	"os"
	"time"

	"4d63.com/tz"
	jl "github.com/soniakeys/meeus/v3/julian"
)

//##############################################
//从数据库中获取数据 cal.sql获取年份同map数组 ccal获取年份从-4000~8000

type calData struct {
	year                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   int
	jd0, Z11a, J12, Z12, J01, Z01, J02, Z02, J03, Z03, J04, Z04, J05, Z05, J06, Z06, J07, Z07, J08, Z08, J09, Z09, J10, Z10, J11, Z11b, Q0_01, Q1_01, Q2_01, Q3_01, Q0_02, Q1_02, Q2_02, Q3_02, Q0_03, Q1_03, Q2_03, Q3_03, Q0_04, Q1_04, Q2_04, Q3_04, Q0_05, Q1_05, Q2_05, Q3_05, Q0_06, Q1_06, Q2_06, Q3_06, Q0_07, Q1_07, Q2_07, Q3_07, Q0_08, Q1_08, Q2_08, Q3_08, Q0_09, Q1_09, Q2_09, Q3_09, Q0_10, Q1_10, Q2_10, Q3_10, Q0_11, Q1_11, Q2_11, Q3_11, Q0_12, Q1_12, Q2_12, Q3_12, Q0_13, Q1_13, Q2_13, Q3_13, Q0_14, Q1_14, Q2_14, Q3_14, Q0_15, Q1_15, Q2_15, Q3_15 float64
}

//用数据库获取年份数据 cal:1600~3500 ccal:-4000~8000
// func Data(y int) []float64 {
// 	cfg := mysql.NewConfig()
// 	cfg.User = "root"
// 	cfg.Passwd = "passwd" //这里改你的密码
// 	cfg.Net = "tcp"
// 	cfg.Addr = "127.0.0.1:3306"
// 	cfg.DBName = "ccal" //数据库
// 	db, err := sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(err)
// 	}
// 	//fmt.Println("connected!`")
// 	dbf, err := calByYear(int64(y), db)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return dbf
// }

//查询单行 传入年数字 数据库
func calByYear(year int64, db *sql.DB) ([]float64, error) {
	var alb calData
	var dbf []float64
	row := db.QueryRow("SELECT * FROM ccal.ccal WHERE year = ?", year)
	if err := row.Scan(&alb.year, &alb.jd0, &alb.Z11a, &alb.J12, &alb.Z12, &alb.J01, &alb.Z01, &alb.J02, &alb.Z02, &alb.J03, &alb.Z03, &alb.J04, &alb.Z04, &alb.J05, &alb.Z05, &alb.J06, &alb.Z06, &alb.J07, &alb.Z07, &alb.J08, &alb.Z08, &alb.J09, &alb.Z09, &alb.J10, &alb.Z10, &alb.J11, &alb.Z11b, &alb.Q0_01, &alb.Q1_01, &alb.Q2_01, &alb.Q3_01, &alb.Q0_02, &alb.Q1_02, &alb.Q2_02, &alb.Q3_02, &alb.Q0_03, &alb.Q1_03, &alb.Q2_03, &alb.Q3_03, &alb.Q0_04, &alb.Q1_04, &alb.Q2_04, &alb.Q3_04, &alb.Q0_05, &alb.Q1_05, &alb.Q2_05, &alb.Q3_05, &alb.Q0_06, &alb.Q1_06, &alb.Q2_06, &alb.Q3_06, &alb.Q0_07, &alb.Q1_07, &alb.Q2_07, &alb.Q3_07, &alb.Q0_08, &alb.Q1_08, &alb.Q2_08, &alb.Q3_08, &alb.Q0_09, &alb.Q1_09, &alb.Q2_09, &alb.Q3_09, &alb.Q0_10, &alb.Q1_10, &alb.Q2_10, &alb.Q3_10, &alb.Q0_11, &alb.Q1_11, &alb.Q2_11, &alb.Q3_11, &alb.Q0_12, &alb.Q1_12, &alb.Q2_12, &alb.Q3_12, &alb.Q0_13, &alb.Q1_13, &alb.Q2_13, &alb.Q3_13, &alb.Q0_14, &alb.Q1_14, &alb.Q2_14, &alb.Q3_14, &alb.Q0_15, &alb.Q1_15, &alb.Q2_15, &alb.Q3_15); err != nil {
		if err == sql.ErrNoRows {
			dbf = append(dbf, alb.jd0, alb.Z11a, alb.J12, alb.Z12, alb.J01, alb.Z01, alb.J02, alb.Z02, alb.J03, alb.Z03, alb.J04, alb.Z04, alb.J05, alb.Z05, alb.J06, alb.Z06, alb.J07, alb.Z07, alb.J08, alb.Z08, alb.J09, alb.Z09, alb.J10, alb.Z10, alb.J11, alb.Z11b, alb.Q0_01, alb.Q1_01, alb.Q2_01, alb.Q3_01, alb.Q0_02, alb.Q1_02, alb.Q2_02, alb.Q3_02, alb.Q0_03, alb.Q1_03, alb.Q2_03, alb.Q3_03, alb.Q0_04, alb.Q1_04, alb.Q2_04, alb.Q3_04, alb.Q0_05, alb.Q1_05, alb.Q2_05, alb.Q3_05, alb.Q0_06, alb.Q1_06, alb.Q2_06, alb.Q3_06, alb.Q0_07, alb.Q1_07, alb.Q2_07, alb.Q3_07, alb.Q0_08, alb.Q1_08, alb.Q2_08, alb.Q3_08, alb.Q0_09, alb.Q1_09, alb.Q2_09, alb.Q3_09, alb.Q0_10, alb.Q1_10, alb.Q2_10, alb.Q3_10, alb.Q0_11, alb.Q1_11, alb.Q2_11, alb.Q3_11, alb.Q0_12, alb.Q1_12, alb.Q2_12, alb.Q3_12, alb.Q0_13, alb.Q1_13, alb.Q2_13, alb.Q3_13, alb.Q0_14, alb.Q1_14, alb.Q2_14, alb.Q3_14, alb.Q0_15, alb.Q1_15, alb.Q2_15, alb.Q3_15)
			return dbf, fmt.Errorf("calById %d: no such cal", year)
		}
		dbf = append(dbf, alb.jd0, alb.Z11a, alb.J12, alb.Z12, alb.J01, alb.Z01, alb.J02, alb.Z02, alb.J03, alb.Z03, alb.J04, alb.Z04, alb.J05, alb.Z05, alb.J06, alb.Z06, alb.J07, alb.Z07, alb.J08, alb.Z08, alb.J09, alb.Z09, alb.J10, alb.Z10, alb.J11, alb.Z11b, alb.Q0_01, alb.Q1_01, alb.Q2_01, alb.Q3_01, alb.Q0_02, alb.Q1_02, alb.Q2_02, alb.Q3_02, alb.Q0_03, alb.Q1_03, alb.Q2_03, alb.Q3_03, alb.Q0_04, alb.Q1_04, alb.Q2_04, alb.Q3_04, alb.Q0_05, alb.Q1_05, alb.Q2_05, alb.Q3_05, alb.Q0_06, alb.Q1_06, alb.Q2_06, alb.Q3_06, alb.Q0_07, alb.Q1_07, alb.Q2_07, alb.Q3_07, alb.Q0_08, alb.Q1_08, alb.Q2_08, alb.Q3_08, alb.Q0_09, alb.Q1_09, alb.Q2_09, alb.Q3_09, alb.Q0_10, alb.Q1_10, alb.Q2_10, alb.Q3_10, alb.Q0_11, alb.Q1_11, alb.Q2_11, alb.Q3_11, alb.Q0_12, alb.Q1_12, alb.Q2_12, alb.Q3_12, alb.Q0_13, alb.Q1_13, alb.Q2_13, alb.Q3_13, alb.Q0_14, alb.Q1_14, alb.Q2_14, alb.Q3_14, alb.Q0_15, alb.Q1_15, alb.Q2_15, alb.Q3_15)
		return dbf, fmt.Errorf("calById %d: %v", year, err)
	}
	dbf = append(dbf, alb.jd0, alb.Z11a, alb.J12, alb.Z12, alb.J01, alb.Z01, alb.J02, alb.Z02, alb.J03, alb.Z03, alb.J04, alb.Z04, alb.J05, alb.Z05, alb.J06, alb.Z06, alb.J07, alb.Z07, alb.J08, alb.Z08, alb.J09, alb.Z09, alb.J10, alb.Z10, alb.J11, alb.Z11b, alb.Q0_01, alb.Q1_01, alb.Q2_01, alb.Q3_01, alb.Q0_02, alb.Q1_02, alb.Q2_02, alb.Q3_02, alb.Q0_03, alb.Q1_03, alb.Q2_03, alb.Q3_03, alb.Q0_04, alb.Q1_04, alb.Q2_04, alb.Q3_04, alb.Q0_05, alb.Q1_05, alb.Q2_05, alb.Q3_05, alb.Q0_06, alb.Q1_06, alb.Q2_06, alb.Q3_06, alb.Q0_07, alb.Q1_07, alb.Q2_07, alb.Q3_07, alb.Q0_08, alb.Q1_08, alb.Q2_08, alb.Q3_08, alb.Q0_09, alb.Q1_09, alb.Q2_09, alb.Q3_09, alb.Q0_10, alb.Q1_10, alb.Q2_10, alb.Q3_10, alb.Q0_11, alb.Q1_11, alb.Q2_11, alb.Q3_11, alb.Q0_12, alb.Q1_12, alb.Q2_12, alb.Q3_12, alb.Q0_13, alb.Q1_13, alb.Q2_13, alb.Q3_13, alb.Q0_14, alb.Q1_14, alb.Q2_14, alb.Q3_14, alb.Q0_15, alb.Q1_15, alb.Q2_15, alb.Q3_15)

	return dbf, nil
}

//##############################################
//从map数组中获取年份数据
func Data(y int) (data []float64) {
	for k, v := range mapOfYears {
		if k == y {
			data = v
			break
		}
	}
	return
}

//从上年冬至到下年立春的节气名称
func NewJQArr(y int) *JQArr {
	jqArrT := JqT(y)
	var jmc []string
	var jqt []time.Time
	for i := 0; i < len(jqArrT); i++ {
		jmc = append(jmc, Jmc[i])
		jqt = append(jqt, jqArrT[i])
	}
	jqarr := &JQArr{
		Name: jmc,
		Time: jqt,
	}
	return jqarr
}

//节气相关计算
//阳历年份节气数组 冬至--冬至
func JqT(y int) []time.Time {
	var j24t []time.Time
	data := Data(y)
	for i := 1; i <= 25; i++ {
		hs := data[0] + data[i]
		hs8 := JdToLocalTime(hs)
		j24t = append(j24t, hs8)
	}
	return j24t
}

//立春和指定时间的比较值 true在立春之后或者等于立春 false在立春之前 同时返回立春时间戳(精确到日)
func (jqarr *JQArr) LinChuT(cust time.Time) (bool, time.Time) {
	m1t := jqarr.Time[3]
	m1t = time.Date(m1t.Year(), time.Month(int(m1t.Month())), m1t.Day(), 0, 0, 0, 0, time.Local)
	cust = time.Date(cust.Year(), cust.Month(), cust.Day(), 0, 0, 0, 0, time.Local)

	var b bool
	if cust.Equal(m1t) || cust.After(m1t) {
		b = true
	} else {
		b = false
	}
	return b, m1t
}

//节气数组 0大雪 1小寒 2立春 3惊蛰 4清明 5立夏 6芒种 7小暑 8立秋 9白露 10寒露 11立东
//十一月大雪节 十二月小寒节 正月立春节 二月惊蛰节 三月清明节 四月立夏节
//五月芒种节 六月小暑节 七月立秋节 八月白露节 九月寒露节 十月立冬节 (遇到闰月顺延)
func (jqarr *JQArr) Jie12ArrT() []time.Time {
	var jqArrT []time.Time
	//i: (0冬至 1小寒 3立春、5惊蛰、7清明、立夏、芒种、小暑、立秋、白露、寒露、立冬、23大雪、)
	for i := 1; i < len(jqarr.Time); i += 2 {
		jqt := jqarr.Time[i]
		jqArrT = append(jqArrT, jqt)
	}
	//jqArrT: 0小寒 1立春 2惊蛰 3清明 4立夏 5芒种 6小暑 7立秋 8白露 9 寒露 10立东 11大雪
	daxuet := jqArrT[11]
	times := jqArrT[:11]

	//0大雪 1小寒 2立春 3惊蛰 4清明 5立夏 6芒种 7小暑 8立秋 9白露 10寒露 11立东
	var j12T []time.Time
	j12T = append(j12T, daxuet)
	j12T = append(j12T, times...)
	return j12T
}

//各月中气 十一月冬至 十二月大寒 正月雨水 二月春分 三月谷雨 四月小满 五月夏至
//六月大暑 七月处暑 八月秋分 九月霜降 十月小雪
//中气 从上年冬至开始 13个中气
//0:冬至 大寒 雨水 春分 谷雨 小满 夏至 大暑 处暑 秋分 霜降 小雪 冬至
func (jqarr *JQArr) ZhongQiArrT() []time.Time {
	var zqArrT []time.Time
	//O:冬至
	for i := 0; i < len(jqarr.Time); i += 2 {
		zqt := jqarr.Time[i]
		zqArrT = append(zqArrT, zqt)
	}
	return zqArrT
}

//y,m,d时间对比12节时间 true在节之后 false在节之前
func DiffJieT(y, m, d int, moonShuotj12arrt []*MoonShuoTJ12T) (bool, int, int) {
	t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local) //精确到日
	var b bool
	var indexMGZ int //干支索引
	var indexswN int //朔望索引值
	for i := 0; i < len(moonShuotj12arrt); i++ {
		//最接近t的节
		jiet := moonShuotj12arrt[i].MoonJieT
		jiet = time.Date(jiet.Year(), jiet.Month(), jiet.Day(), 0, 0, 0, 0, time.Local)
		difft := t.Sub(jiet).Hours() / 24 //精确到日 这个计算方法在2033年失效 MonthGZ()的case false:做了强制性修正
		m := i + 11
		if m > 12 {
			m -= 12
		}
		if (int(difft) >= 0 && int(difft) < 30) && (t.Equal(jiet) || t.Before(jiet)) {
			indexswN = i //如果这里没数据 默认就是1
			indexMGZ = m //如果年份含有闰月 值m在闰月之后比实际月份大1 所以不能把m计算为农历月份
			b = true
		} else if (int(difft) >= 0 && int(difft) < 30) && (t.After(jiet)) {
			indexswN = i
			indexMGZ = m
			b = true
		} else if int(difft) < 0 && int(math.Abs(difft)) < 30 {
			indexswN = i - 1
			indexMGZ = m - 1
			if indexMGZ == 0 {
				indexMGZ = 12
			}
			b = true
		}
	}
	return b, indexMGZ, indexswN
}

//农历朔和本月节  0:上年农历十一月(这里没计算农历闰月 所以农历月份要待计算闰月之后才能最终确定)
//这里的T都是阳历时间戳 精确到秒
func MoonShuoTJ12ArrT(j12arrt []time.Time, shuowangt []*ShuoWangT) []*MoonShuoTJ12T {
	var moonsj12T *MoonShuoTJ12T
	var msjArrT []*MoonShuoTJ12T
	for i := 0; i < len(j12arrt); i++ { //0:农历上年11月的节
		for j := 0; j < len(shuowangt); j++ {
			if i == j {
				moonshuot := shuowangt[j].ShuoT
				moonj12t := j12arrt[j]
				moonsj12T = &MoonShuoTJ12T{
					MoonShuoT: moonshuot,
					MoonJieT:  moonj12t,
				}
				msjArrT = append(msjArrT, moonsj12T)
				break
			}

		}
	}
	return msjArrT
}

//0:上年十一月 ... (含闰月的年份 14:本年十二月) (不含润月的年份 14:下一年正月)
//十四个月的朔 上弦 望 下弦 时间精确到分钟
func MoonShuoWangT(shuoWObj []*ShuoWangF) []*ShuoWangT {
	var swTArrT []*ShuoWangT
	for i := 0; i < len(shuoWObj); i++ {
		shuojd := shuoWObj[i].ShuoF
		shangjd := shuoWObj[i].ShangXianF
		wangjd := shuoWObj[i].WangF
		xiajd := shuoWObj[i].XiaXianF
		shuot := JdToLocalTime(shuojd)
		shangt := JdToLocalTime(shangjd)
		wangt := JdToLocalTime(wangjd)
		xiat := JdToLocalTime(xiajd)
		swT := &ShuoWangT{
			ShuoT:      shuot,
			ShangXianT: shangt,
			WangT:      wangt,
			XiaXianT:   xiat,
		}
		swTArrT = append(swTArrT, swT)
	}
	return swTArrT
}

//0:阴历上年十一月  1:阴历上年十二月 3:阴历本年正月
//十五个月的朔望数据
func MoonShuoWangF(shuoWangObjArr []*ShuoWangF) []*ShuoWangF {
	var moonArr []*ShuoWangF
	for i := 3; i < len(shuoWangObjArr); i += 4 { //从索引3开始+=4取最后这组结构体值
		moonArr = append(moonArr, shuoWangObjArr[i])
	}
	return moonArr
}

//JD数据 从阴历上一年十一月开始到本年十月结束
func NewShuoWangF(index int, data []float64) []*ShuoWangF {
	var indexShuo, indexShangXian, indexWang, indexXiaXian float64
	var swArr []*ShuoWangF
	for i := index; i < len(data); i++ {
		switch (i - index) % 4 {
		case 0:
			indexShuo = data[0] + data[i]
		case 1:
			indexShangXian = data[0] + data[i]
		case 2:
			indexWang = data[0] + data[i]
		case 3:
			indexXiaXian = data[0] + data[i]
		}
		sw := &ShuoWangF{
			ShuoF:      indexShuo,
			ShangXianF: indexShangXian,
			WangF:      indexWang,
			XiaXianF:   indexXiaXian,
		}
		swArr = append(swArr, sw)
	}
	return swArr
}

func LeapmB(y int) bool {
	data := Data(y)
	_, m1t, _ := m1(data) //第一个冬至前的朔
	m11t, _ := m11(data)  //第二个冬至前的朔
	d := m11t.Sub(m1t).Hours() / 24
	l := d / 29.53
	x := math.Dim(l, 12.0) > 0.9 //true两冬至间有不含中气的闰月
	return x
}

func ZhongQiArrT(y int) []time.Time {
	data1 := Data(y)
	jd10 := data1[0]
	var zq1Arr []time.Time
	for i := 1; i < 25; i += 2 { //上一年冬至到下一年小雪
		jdx := jd10 + data1[i]
		jdxt := JdToLocalTime(jdx)
		zq1Arr = append(zq1Arr, jdxt)
	}
	return zq1Arr
}

//两个年份的中气
func AllZhongQi(y int) []time.Time {
	data1 := Data(y)
	jd10 := data1[0]
	var zq1Arr, zq11Arr []time.Time
	for i := 1; i < 25; i += 2 { //上一年冬至到下一年小雪
		jdx := jd10 + data1[i]
		jdxt := JdToLocalTime(jdx)
		zq1Arr = append(zq1Arr, jdxt)
	}

	data11 := Data(y + 1)
	jd110 := data11[0]
	for j := 1; j < 25; j += 2 {
		jdx := jd110 + data11[j]
		jdxt := JdToLocalTime(jdx)
		zq11Arr = append(zq11Arr, jdxt)
	}
	var all []time.Time
	all = append(all, zq1Arr...)
	all = append(all, zq11Arr...)
	return all
}

//农历全年合朔(2个冬至间的朔数据)
func AllShuo(y int) []float64 {
	data1 := Data(y)
	data11 := Data(y + 1)
	index1, _, m1jd0 := m1(data1)
	_, m1jd1 := m11(data1)
	index11, _, m11jd0 := m1(data11)
	_, m11jd1 := m11(data11)
	shuo1 := Shuox(index1, m1jd0, m1jd1, data1)
	shuo11 := Shuox(index11, m11jd0, m11jd1, data11)

	var allshuo []float64
	//重复的JD数据
	shuo1b := EqualT(shuo1[len(shuo1)-1], shuo11[0])
	shuo11b := EqualT(shuo1[len(shuo1)-2], shuo11[0])
	if shuo1b == true || shuo11b == true { //过滤掉重复的数据
		allshuo = append(shuo1, shuo11[1:]...)
	}
	return allshuo
}

//朔的时间戳比较 时间精确到小时
func EqualT(jd1, jd11 float64) bool {
	t1 := JdToLocalTime(jd1)
	t11 := JdToLocalTime(jd11)
	//时间精确到小时
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), 0, 0, 0, time.Local)
	t11 = time.Date(t11.Year(), t11.Month(), t11.Day(), t11.Hour(), 0, 0, 0, time.Local)

	return t1.Equal(t11)
}

//两冬至间朔 冬至前距离最近的朔 11月到11月
func Shuox(index int, m1jd, m11jd float64, datax []float64) (lunarShuox []float64) {
	var hsjdarr []float64
	for i := 26; i < len(datax); i += 4 {
		_hsjd := datax[0] + datax[i]
		hsjdarr = append(hsjdarr, _hsjd)
		jd11 := datax[0] + datax[i]
		if math.Dim(jd11, m11jd) < 0.9 && jd11 == m11jd {
			break
		}
	}
	lunarShuox = hsjdarr
	return
}

//阳历时间戳m1t通常是上一年农历十一月的初一
//index 即是JqHs数组对应的索引值
//第一个冬至前的朔 时间精确到日
func m1(data []float64) (index int, m1t time.Time, m1jd float64) {
	jd0 := data[0]
	dz0 := jd0 + data[1]       //第一个冬至(出现在上一年)
	dz0t := JdToLocalTime(dz0) //第一个冬至的时间戳
	//26第一个朔的时间 这里要确定这个数字是不是冬至前 距离冬至最近的一个朔
	for i := 26; i < len(data); i += 4 {
		jdx := jd0 + data[i]       //第i个朔的jd数据
		jdxt := JdToLocalTime(jdx) //第i个朔的时间戳
		//精确到日对比 确定两冬至间隔是否有不含中气的月份
		dz0t = time.Date(dz0t.Year(), dz0t.Month(), dz0t.Day(), 0, 0, 0, 0, time.UTC)
		jdxt = time.Date(jdxt.Year(), jdxt.Month(), jdxt.Day(), 0, 0, 0, 0, time.UTC)

		if jdxt.Equal(dz0t) == true { //朔和冬至同一天(精确到日)
			index = i
			m1jd = jdx
			m1t = JdToLocalTime(jdx)
		} else if jdxt.Equal(dz0t) == false && jdxt.Before(dz0t) { //朔在冬至前
			index = i
			m1jd = jdx
			m1t = JdToLocalTime(jdx)
		}
	}
	return
}

//它的值小于等于第一个冬至 特殊年份(1984 2014)等于冬至
//第二个冬至前的朔 时间精确到日
func m11(data []float64) (m11t time.Time, m11jd float64) {
	jd0 := data[0]
	dz1 := jd0 + data[25]
	dz1t := JdToLocalTime(dz1)
	//精确到时的计算方法
	for i := 74; i < len(data); i += 4 {
		jdx := jd0 + data[i]
		jdxt := JdToLocalTime(jdx)
		//精确到日对比 确定两冬至间隔是否有不含中气的月份
		dz1t = time.Date(dz1t.Year(), dz1t.Month(), dz1t.Day(), 0, 0, 0, 0, time.UTC)
		jdxt = time.Date(jdxt.Year(), jdxt.Month(), jdxt.Day(), 0, 0, 0, 0, time.UTC)

		//朔和冬至同一天 1984,2014,2166,2204,2223,2318,2386,2405,2500,2557,2576,2595,2758,2853
		if jdxt.Equal(dz1t) == true {
			m11jd = jdx
			m11t = JdToLocalTime(jdx)
		} else if jdxt.Equal(dz1t) == false && jdxt.Before(dz1t) { //朔在冬至前
			m11jd = jdx
			m11t = JdToLocalTime(jdx) //这里保持原始转换时间
		}
	}
	return
}

//儒略日转time.Time 这里精确到秒
func JdToLocalTime(jd float64) time.Time {
	FormatT := "2006-01-02 15:04:05"
	cst := localTime(jd)
	t, err := time.Parse(FormatT, cst)
	if err != nil {
		fmt.Println("JDToLocalTime解析异常:", err)
		os.Exit(0)
	}
	return t
}

//转为为本地时间(Asia/Shanghai)
func localTime(jd float64) string {
	utc := jl.JDToTime(jd) //UTC时间
	local := utc
	location, err := tz.LoadLocation("Asia/Shanghai") //windows系统的兼容方法
	if err == nil {
		local = local.In(location)
	}
	s := local.Format("2006-01-02 15:04:05")
	return s
}

//阳历年月日转换为ind类型的JD数据
//用来计算日干支
func SolarYmdToJD(y, m, d int) (solarJd int) {
	ra := (14 - m) / 12
	ry := y + 4800 - ra
	rm := m + 12*ra - 3
	solarJd = (d + (153*rm+2)/5 + 365*ry + (ry / 4) - (ry / 100) + (ry / 400) - 32045)
	return
}
func covnH24Toh12(h int) int {
	var h12 int
	switch h {
	case 23, 00:
		h12 = 1
	case 1, 2:
		h12 = 2
	case 3, 4:
		h12 = 3
	case 5, 6:
		h12 = 4
	case 7, 8:
		h12 = 5
	case 9, 10:
		h12 = 6
	case 11, 12:
		h12 = 7
	case 13, 14:
		h12 = 8
	case 15, 16:
		h12 = 9
	case 17, 18:
		h12 = 10
	case 19, 20:
		h12 = 11
	case 21, 22:
		h12 = 12
	}
	return h12
}

type MoonSW float64

//月初jd
func (msw MoonSW) MoonJD(y int) (jd float64) {
	index, _, _ := m1(Data(y))
	shuoWangFArr := NewShuoWangF(index, Data(y))
	moonSWFarr := MoonShuoWangF(shuoWangFArr)
	for i := 0; i < len(moonSWFarr); i++ {
		//判断当前日期
		jd = moonSWFarr[i].ShuoF
	}
	return
}
