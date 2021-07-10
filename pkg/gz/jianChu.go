package gz

import "strings"

//日建除
func GetRiJianChu(mgz, dgz string) string {
	mjc := jcMonthMap(mgz)
	return jcToday(dgz, mjc)
}

//日建除
func jcToday(dgz string, mjc map[string]string) string {
	var info string
	for dizhi, jc := range mjc {
		if strings.ContainsAny(dgz, dizhi) {
			info = jc
			break
		}
	}
	return info
}

//正月建寅 二月建卯 三月建辰 四月建巳 ...十二月建丑
//月建除 以月干支计算
func jcMonthMap(mgz string) (mjc map[string]string) {
	zhi := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"}
	var indexM int
	for i := 0; i < len(zhi); i++ {
		if strings.Contains(mgz, zhi[i]) {
			indexM = i
			break
		}
	}
	switch indexM {
	case 0:
		mjc = jcmMap(indexM)
	case 1:
		mjc = jcmMap(indexM)
	case 2:
		mjc = jcmMap(indexM)
	case 3:
		mjc = jcmMap(indexM)
	case 4:
		mjc = jcmMap(indexM)
	case 5:
		mjc = jcmMap(indexM)
	case 6:
		mjc = jcmMap(indexM)
	case 7:
		mjc = jcmMap(indexM)
	case 8:
		mjc = jcmMap(indexM)
	case 9:
		mjc = jcmMap(indexM)
	case 10:
		mjc = jcmMap(indexM)
	case 11:
		mjc = jcmMap(indexM)
	}
	return
}
func jcmMap(indexM int) (mjcmap map[string]string) {
	//十二建除
	jcArr := []string{"建", "除", "满", "平", "定", "执", "破", "危", "成", "收", "开", "闭"}
	//地支名称
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	var maps = make(map[string]string) //地支配十二建除　k:地支　v:十二建
	switch indexM {
	case 10: //十一月 建子
		for i := 0; i < len(zhi); i++ {
			for j := 0; j < len(jcArr); j++ {
				if i == j {
					maps[zhi[i]] = jcArr[j]
					break
				}
			}
		}
		mjcmap = maps
	case 11: //十二月 建丑
		sliceHead := zhi[:1]
		var newZhi []string
		newZhi = append(zhi[1:], sliceHead...)
		for i := 0; i < len(newZhi); i++ {
			for j := 0; j < len(jcArr); j++ {
				if i == j {
					maps[newZhi[i]] = jcArr[j]
					break
				}
			}
		}
		mjcmap = maps
	case 0: //正月 建寅
		sliceHead := zhi[:2]
		var newZhi []string
		newZhi = append(zhi[2:], sliceHead...)
		for i := 0; i < len(newZhi); i++ {
			for j := 0; j < len(jcArr); j++ {
				if i == j {
					maps[newZhi[i]] = jcArr[j]
					break
				}
			}
		}
		mjcmap = maps
	case 1: //二月 建卯
		sliceHead := zhi[:3]
		var newZhi []string
		newZhi = append(zhi[3:], sliceHead...)
		for i := 0; i < len(newZhi); i++ {
			for j := 0; j < len(jcArr); j++ {
				if i == j {
					maps[newZhi[i]] = jcArr[j]
					break
				}
			}
		}
		mjcmap = maps
	case 2: //三月 建辰
		sliceHead := zhi[:4]
		var newZhi []string
		newZhi = append(zhi[4:], sliceHead...)
		for i := 0; i < len(newZhi); i++ {
			for j := 0; j < len(jcArr); j++ {
				if i == j {
					maps[newZhi[i]] = jcArr[j]
					break
				}
			}
		}
		mjcmap = maps
	case 3: //四月
		sliceHead := zhi[:5]
		var newZhi []string
		newZhi = append(zhi[5:], sliceHead...)
		for i := 0; i < len(newZhi); i++ {
			for j := 0; j < len(jcArr); j++ {
				if i == j {
					maps[newZhi[i]] = jcArr[j]
					break
				}
			}
		}
		mjcmap = maps
	case 4: //五月
		sliceHead := zhi[:6]
		var newZhi []string
		newZhi = append(zhi[6:], sliceHead...)
		for i := 0; i < len(newZhi); i++ {
			for j := 0; j < len(jcArr); j++ {
				if i == j {
					maps[newZhi[i]] = jcArr[j]
					break
				}
			}
		}
		mjcmap = maps
	case 5: //六月
		sliceHead := zhi[:7]
		var newZhi []string
		newZhi = append(zhi[7:], sliceHead...)
		for i := 0; i < len(newZhi); i++ {
			for j := 0; j < len(jcArr); j++ {
				if i == j {
					maps[newZhi[i]] = jcArr[j]
					break
				}
			}
		}
		mjcmap = maps
	case 6: //七月
		sliceHead := zhi[:8]
		var newZhi []string
		newZhi = append(zhi[8:], sliceHead...)
		for i := 0; i < len(newZhi); i++ {
			for j := 0; j < len(jcArr); j++ {
				if i == j {
					maps[newZhi[i]] = jcArr[j]
					break
				}
			}
		}
		mjcmap = maps
	case 7: //八月
		sliceHead := zhi[:9]
		var newZhi []string
		newZhi = append(zhi[9:], sliceHead...)
		for i := 0; i < len(newZhi); i++ {
			for j := 0; j < len(jcArr); j++ {
				if i == j {
					maps[newZhi[i]] = jcArr[j]
					break
				}
			}
		}
		mjcmap = maps
	case 8: //九月
		sliceHead := zhi[:10]
		var newZhi []string
		newZhi = append(zhi[10:], sliceHead...)
		for i := 0; i < len(newZhi); i++ {
			for j := 0; j < len(jcArr); j++ {
				if i == j {
					maps[newZhi[i]] = jcArr[j]
					break
				}
			}
		}
		mjcmap = maps
	case 9: //十月
		sliceHead := zhi[:11]
		var newZhi []string
		newZhi = append(zhi[11:], sliceHead...)
		for i := 0; i < len(newZhi); i++ {
			for j := 0; j < len(jcArr); j++ {
				if i == j {
					maps[newZhi[i]] = jcArr[j]
					break
				}
			}
		}
		mjcmap = maps
	}

	return
}