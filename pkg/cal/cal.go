package cal

import (
	"errors"
	"fmt"
	"log"
	"time"
)

//阳历年月日时 转换为干支
func NewCal(y, m, d, h int) *Cal {
	jqarrOBj := NewJQArr(y)
	cust := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	b, _ := jqarrOBj.LinChuT(cust)
	//节气数组 0小寒 1立春 2惊蛰 3清明 4立夏 5芒种 6小暑 7立秋 8白露 9 寒露 10立东 11大雪
	jie12arrT := jqarrOBj.Jie12ArrT()

	yearg, yearz := YearGZ(y, b)
	ygz := fmt.Sprintf("%s%s", yearg, yearz)

	index, _, _ := m1(Data(y))
	shuoWangObjArr := NewShuoWangF(index, Data(y))
	moonSW := MoonShuoWangF(shuoWangObjArr)
	moonSwArrT := MoonShuoWangT(moonSW)
	moonswj12arrT := MoonShuoTJ12ArrT(jie12arrT, moonSwArrT)
	diffjietB, indexMGZ, _ := DiffJieT(y, m, d, moonswj12arrT)
	mgz, err := MonthGZ(yearg, indexMGZ, diffjietB)
	if err != nil {
		log.Fatal(err)
	}
	dgz, daygN := DayGZ(y, m, d)
	hgz := HourGZ(daygN, h)
	gzinfo := &Cal{
		YearGZ:  ygz,
		MonthGZ: mgz,
		DayGZ:   dgz,
		HourGZ:  hgz,
	}
	return gzinfo
}

//干支纳因
func (c *Cal) NaYin() string {
	return NaYin(c.YearGZ, c.MonthGZ, c.DayGZ, c.HourGZ)
}

//年干支
func YearGZ(solary int, b bool) (aliasGan, aliasZhi string) {
	switch b {
	case false: //false:日期在立春之前
		//干
		g := 1 + (solary+6)%10
		if g -= 1; g < 1 {
			g += 10
		}
		aliasGan = Gan[g]
		//支
		z := 1 + (solary+8)%12
		if z -= 1; z < 1 {
			z += 12
		}
		aliasZhi = Zhi[z]
	case true: //日期在立春之后
		yearg := 1 + (solary+6)%10
		yearz := 1 + (solary+8)%12
		aliasGan = Gan[yearg]
		aliasZhi = Zhi[yearz]
	}
	return
}

//月干支
func MonthGZ(yearg string, indexMGZ int, diffJietB bool) (string, error) {
	switch diffJietB {
	case true: //正常五虎盾元
		return WuHuDun(yearg, indexMGZ), nil
	case false:
		indexMGZ = 11 //2033-1-1这是一个简单粗暴的方法
		return WuHuDun(yearg, indexMGZ), nil
	default:
		err := errors.New("月干支异常")
		return "", err
	}
}

//五虎盾元
func WuHuDun(yearg string, indexMGZ int) (mgz string) {
	switch yearg {
	case Gan[1], Gan[6]:
		gzarr := []string{"err", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉", "甲戌", "乙亥", "丙子", "丁丑"}
		mgz = gzarr[indexMGZ]
	case Gan[2], Gan[7]:
		gzarr := []string{"err", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未", "甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑"}
		mgz = gzarr[indexMGZ]
	case Gan[3], Gan[8]:
		gzarr := []string{"err", "庚寅", "辛卯", "壬辰", "癸巳", "甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑"}
		mgz = gzarr[indexMGZ]
	case Gan[4], Gan[9]:
		gzarr := []string{"err", "壬寅", "癸卯", "甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑"}
		mgz = gzarr[indexMGZ]
	case Gan[5], Gan[10]:
		gzarr := []string{"err", "甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥", "甲子", "乙丑"}
		mgz = gzarr[indexMGZ]
	}
	return
}

//日干支
//农历日干支 日干数字
func DayGZ(y, m, d int) (string, int) {
	solarDayJd := SolarYmdToJD(y, m, d)
	//如果干数字=0则为10
	g := 1 + (solarDayJd%60-1)%10 //干
	if g == 0 {
		g += 10
	}
	z := 1 + +(solarDayJd%60+1)%12 //支

	daygN := g //日干数字
	daygM := Gan[g]
	dayzM := Zhi[z]
	dgz := daygM + dayzM
	return dgz, daygN
}

//阳历时间对应的时辰干支
func HourGZ(daygN, h int) string {
	h = covnH24Toh12(h)
	var gz []string
	switch daygN {
	case 1, 6: //甲己还加甲
		gz = []string{"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉", "甲戌", "乙亥"}
	case 2, 7: //乙庚丙作初
		gz = []string{"丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未", "甲申", "乙酉", "丙戌", "丁亥"}
	case 3, 8: //丙辛从戊起
		gz = []string{"戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳", "甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥"}
	case 4, 9: //丁壬庚子居
		gz = []string{"庚子", "辛丑", "壬寅", "癸卯", "甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥"}
	case 5, 10: //戊癸何方发?壬子是真途
		gz = []string{"壬子", "癸丑", "甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"}
	}
	aliasHGZ := gz[h-1]
	return aliasHGZ
}

// 五虎盾元 甲己之年丙作初，乙庚之歲戊為頭，丙辛歲首從庚起，丁壬壬位順流行，若問戊癸何方法，甲寅之上好推求
// 传入年干 阴历月份数字 返回月干支数组(0正月 11腊月)
//  月干支数组
func mgzArr(yg string) []string {
	gan := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	zhi := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"} //正月建寅
	var arr []string                                                            //月干支数组
	switch yg {
	case gan[0], gan[5]: //甲己 丙寅
		end := gan[:2]
		head := gan[2:]
		arr = whd(gan, zhi, head, end)
	case gan[1], gan[6]: //乙庚 戊寅
		end := gan[:4]
		head := gan[4:]
		arr = whd(gan, zhi, head, end)
	case gan[2], gan[7]: //丙辛 庚寅
		end := gan[:6]
		head := gan[6:]
		arr = whd(gan, zhi, head, end)
	case gan[3], gan[8]: //丁壬 壬寅
		end := gan[:8]
		head := gan[8:]
		arr = whd(gan, zhi, head, end)
	case gan[4], gan[9]: //戊癸 甲寅
		end := gan
		head := gan
		arr = whd(gan, zhi, head, end)
	}
	return arr
}

//五虎遁
func whd(gan, zhi, head, end []string) []string {
	var arr []string
	gan = append(head, end...)
	gan = append(gan, gan...)
	for i := 0; i < len(gan); i++ {
		for j := i; j < len(zhi); j++ {
			s := gan[j] + zhi[j]
			arr = append(arr, s)
			if j == i {
				break
			}

		}
	}
	return arr
}
