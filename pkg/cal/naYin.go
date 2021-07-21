package cal

import (
	"fmt"
	"strings"
)

func NaYin(YGZ, MGZ, DGZ, HGZ string) string {
	ygzny := ny(YGZ)
	mgzny := ny(MGZ)
	dgzny := ny(DGZ)
	hgzny := ny(HGZ)
	nys := fmt.Sprintf("%s %s %s %s",
		ygzny[YGZ], mgzny[MGZ], dgzny[DGZ], hgzny[HGZ])
	return nys
}

//干支纳音
func ny(gz string) (nywx map[string]string) {
	wx := []string{
		"海中金", "炉中火", "大林木", "路旁土",
		"剑锋金", "山头火", "涧下水", "城头土",
		"白蜡金", "杨柳木", "泉中水", "屋上土",
		"霹雳火", "松柏木", "长流水", //14
		"沙中金", "山下火", "平地木", "壁上土",
		"金箔金", "覆灯火", "天河水", "大驿土",
		"钗钏金", "桑柘木", "大溪水", "沙中土",
		"天上火", "石榴木", "大海水", //29
	}
	//六十甲子
	jz := MakeJZ60()
	//纳音
	nywx = make(map[string]string)
	//六十甲子数组索引号除以2求商　商为当前干支的纳音索引数字
	for i := 0; i < len(jz); i++ {
		if strings.EqualFold(gz, jz[i]) {
			index := i / 2
			nywx[jz[i]] = wx[index]
		}
	}
	return
}

//生成六十甲子
func MakeJZ60() []string {
	gan := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

	var gzs = []string{}
	//阳干配阳支　甲阳数１　子阳数１　甲配子－甲子...
	//阴干配阴之　乙阴数２　丑阴数２　乙配丑－乙丑...
	for i := 0; i < len(gan); i++ {
		for j := 0; j < len(zhi); j++ {
			g := i % 2
			z := j % 2
			if g == z { //过滤掉阴阳不相配的干支
				gs := gan[i]
				zs := zhi[j]
				gz := fmt.Sprintf("%s%s", gs, zs)
				gzs = append(gzs, gz)
			}

		}
	}

	//分别列出十天干数组
	jia := gzs[:6]
	yi := gzs[6:12]
	bing := gzs[12:18]
	ding := gzs[18:24]
	wu := gzs[24:30]
	ji := gzs[30:36]
	geng := gzs[36:42]
	xin := gzs[42:48]
	ren := gzs[48:54]
	gui := gzs[54:60]

	var jzs, jxs, jss, jws, jcs, jys []string //六甲旬
	for x := 0; x < len(gzs); x++ {
		n := x % 6
		switch n {
		case 0: //甲子旬
			jzs = append(jzs, jia[0], yi[0], bing[1], ding[1], wu[2], ji[2], geng[3], xin[3], ren[4], gui[4])
		case 5: //甲戌旬
			jxs = append(jxs, jia[5], yi[5], bing[0], ding[0], wu[1], ji[1], geng[2], xin[2], ren[3], gui[3])
		case 4: //甲申旬
			jss = append(jss, jia[4], yi[4], bing[5], ding[5], wu[0], ji[0], geng[1], xin[1], ren[2], gui[2])
		case 3: //甲午旬
			jws = append(jws, jia[3], yi[3], bing[4], ding[4], wu[5], ji[5], geng[0], xin[0], ren[1], gui[1])
		case 2: //甲辰旬
			jcs = append(jcs, jia[2], yi[2], bing[3], ding[3], wu[4], ji[4], geng[5], xin[5], ren[0], gui[0])
		case 1: //甲寅旬
			jys = append(jys, jia[1], yi[1], bing[2], ding[2], wu[3], ji[3], geng[4], xin[4], ren[5], gui[5])
		}

	}
	var jz60 []string
	//按六甲旬顺序(甲子　甲戌　甲申　甲午　甲辰　甲寅)添加元素
	jz60 = append(jz60, jzs[:10]...)
	jz60 = append(jz60, jxs[:10]...)
	jz60 = append(jz60, jss[:10]...)
	jz60 = append(jz60, jws[:10]...)
	jz60 = append(jz60, jcs[:10]...)
	jz60 = append(jz60, jys[:10]...)

	return jz60
}
