/*
 * The code is automatically generated by the Goland.
 * Copyright © Aquarian-Age. All Rights Reserved.
 * Licensed under MIT.
 */

package gz

import (
	"fmt"
	"github.com/starainrt/astro/calendar"
	"testing"
	"time"
)

func TestAliasGan(t *testing.T) {
	tx := time.Now().Local()
	for i := 0; i <= 31; i++ {
		xt := time.Date(tx.Year(), tx.Month()+1, i, 0, 0, 0, 0, time.Local)
		jd := calendar.Date2JDE(xt)
		n := AliasGan(jd)
		fmt.Println(xt.String()[:10], n)
	}
}

/*
2022-02-28 壬
2022-03-01 癸
2022-03-02 甲
2022-03-03 乙
2022-03-04 丙
2022-03-05 丁
2022-03-06 戊
2022-03-07 己
2022-03-08 庚
2022-03-09 辛
2022-03-10 壬
2022-03-11 癸
2022-03-12 甲
2022-03-13 乙
2022-03-14 丙
2022-03-15 丁
2022-03-16 戊
2022-03-17 己
2022-03-18 庚
2022-03-19 辛
2022-03-20 壬
2022-03-21 癸
2022-03-22 甲
2022-03-23 乙
2022-03-24 丙
2022-03-25 丁
2022-03-26 戊
2022-03-27 己
2022-03-28 庚
2022-03-29 辛
2022-03-30 壬
2022-03-31 癸
*/
func TestGan_ChangSheng(t *testing.T) {
	//gzo := NewGanZhi(2022, 2, 4, 4)
	//fmt.Println(gzo.YGZ, gzo.MGZ, gzo.DGZ)
	//gan := "庚"
	//arr := gzo.ChangShengArr(gan)
	//g := gzo.Gan(gan)
	//changsheng := g.ChangSheng(arr)
	//fmt.Println(gan, "长生", changsheng)
	//fmt.Println(gan, "死", g.Si(arr))
	//fmt.Println(gan, "绝", g.Jue(arr))
}

//	壬寅 壬寅 戊子
//	庚 长生 巳
//	庚 死 子
//	庚 绝 寅