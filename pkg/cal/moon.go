package cal

import "time"

//节气数组 0小寒 1立春 2惊蛰 3清明 4立夏 5芒种 6小暑 7立秋 8白露 9 寒露 10立东 11大雪
//阳历当月对应农历的朔望时间 string类型输出
func NewShuoWangTS(y, m, d int) *ShuoWantTS {
	indexs, _, _ := m1(Data(y))
	shuoWangObjArr := NewShuoWangF(indexs, Data(y))
	moonSW := MoonShuoWangF(shuoWangObjArr)
	moonSwArrT := MoonShuoWangT(moonSW)

	swts := new(ShuoWantTS)
	tf := "2006-01-02 15:04:05"
	start := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.Local)

	var index int
	for i := 0; i < len(moonSwArrT); i++ {
		shuot := moonSwArrT[i].ShuoT
		shangxiant := moonSwArrT[i].ShangXianT
		wangt := moonSwArrT[i].WangT
		xiaxiant := moonSwArrT[i].XiaXianT
		shuot = time.Date(shuot.Year(), shuot.Month(), shuot.Day(), 0, 0, 0, 0, time.Local)
		wangt = time.Date(wangt.Year(), wangt.Month(), wangt.Day(), 0, 0, 0, 0, time.Local)
		shangxiant = time.Date(shangxiant.Year(), shangxiant.Month(), shangxiant.Day(), 0, 0, 0, 0, time.Local)
		xiaxiant = time.Date(xiaxiant.Year(), xiaxiant.Month(), xiaxiant.Day(), 0, 0, 0, 0, time.Local)
		if shuot.Equal(start) || shuot.After(start) {
			index = i
			break
		} else if shangxiant.Equal(start) || shangxiant.After(start) {
			index = i
			break
		} else if wangt.Equal(start) || wangt.After(start) {
			index = i
			break
		} else if xiaxiant.Equal(start) || wangt.After(start) {
			index = i
			break
		}
	}

	shuo := moonSwArrT[index].ShuoT
	wang := moonSwArrT[index].WangT
	shang := moonSwArrT[index].ShangXianT
	xia := moonSwArrT[index].XiaXianT
	shuots := shuo.Format(tf)
	wangts := wang.Format(tf)
	xiaXiants := xia.Format(tf)
	shangXiants := shang.Format(tf)
	swts = &ShuoWantTS{
		ShuoTS:      shuots,
		ShangXianTS: shangXiants,
		WangTS:      wangts,
		XiaXianTS:   xiaXiants,
	}

	return swts
}

//这里是通过节来计算的 即当前阳历日期对应的节 如果阳历日期在节之前显示的是上个月(农历)的朔望 如果在节之后显示的是本月(农历)的朔望
//如 阳历2020年4月4日 清明节 清明为农历三月节 日期在4日之前显示的是农历二月的朔望 在4日之后显示的是农历三月的朔望
//节气数组 0小寒 1立春 2惊蛰 3清明 4立夏 5芒种 6小暑 7立秋 8白露 9 寒露 10立东 11大雪
//阳历年月日对应农历月份的 朔 望 上弦 下弦
func NewShuoWangT(y, m, d int) (*ShuoWangT, *ShuoWantTS) {
	jqarrOBj := NewJQArr(y)
	jie12arrT := jqarrOBj.Jie12ArrT()
	index, _, _ := m1(Data(y))
	shuoWangObjArr := NewShuoWangF(index, Data(y))
	moonSW := MoonShuoWangF(shuoWangObjArr)
	moonSwArrT := MoonShuoWangT(moonSW)
	moonswj12arrT := MoonShuoTJ12ArrT(jie12arrT, moonSwArrT)
	diffjietB, _, indexSWN := DiffJieT(y, m, d, moonswj12arrT)

	//根据节气找 碰到闰月会失效 显示的是非闰月之前和闰月相同月份的朔望
	var swts *ShuoWantTS
	var shuowt *ShuoWangT
	tf := "2006-01-02 15:04:05"
	switch diffjietB {
	case true:
		for j := 0; j < len(moonSwArrT); j++ {
			if j == indexSWN {
				shuo, wang, shangxian, xiaxian := moonSwArrT[j].ShuoT, moonSwArrT[j].WangT, moonSwArrT[j].ShangXianT, moonSwArrT[j].XiaXianT
				shuots := shuo.Format(tf)
				wangts := wang.Format(tf)
				shangxiants := shangxian.Format(tf)
				xiaxiants := xiaxian.Format(tf)

				shuowt = &ShuoWangT{
					ShuoT:      moonSwArrT[j].ShuoT,
					ShangXianT: moonSwArrT[j].ShangXianT,
					WangT:      moonSwArrT[j].WangT,
					XiaXianT:   moonSwArrT[j].XiaXianT,
				}
				swts = &ShuoWantTS{
					ShuoTS:      shuots,
					ShangXianTS: shangxiants,
					WangTS:      wangts,
					XiaXianTS:   xiaxiants,
				}
				break
			}
		}
	case false:
		indexSWN = 1
		for j := 0; j < len(moonSwArrT); j++ {
			if j == indexSWN {
				shuo, wang, shangxian, xiaxian := moonSwArrT[j].ShuoT, moonSwArrT[j].WangT, moonSwArrT[j].ShangXianT, moonSwArrT[j].XiaXianT
				shuots := shuo.Format(tf)
				wangts := wang.Format(tf)
				shangxiants := shangxian.Format(tf)
				xiaxiants := xiaxian.Format(tf)

				shuowt = &ShuoWangT{
					ShuoT:      moonSwArrT[j].ShuoT,
					ShangXianT: moonSwArrT[j].ShangXianT,
					WangT:      moonSwArrT[j].WangT,
					XiaXianT:   moonSwArrT[j].XiaXianT,
				}
				swts = &ShuoWantTS{
					ShuoTS:      shuots,
					ShangXianTS: shangxiants,
					WangTS:      wangts,
					XiaXianTS:   xiaxiants,
				}
				break
			}
		}
	}
	return shuowt, swts
}
