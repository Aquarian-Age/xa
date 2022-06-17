### 干支计算

- 传入阳历年月日时
- func NewGanZhi(year, month, day, hour int) *GanZhi{} 月干支计算精确到日
```text
例如： 2022年2月4日 4:50分立春. 
传入: 2022, 2, 4, 0 显示: 壬寅 壬寅 戊子 壬子
```
- func NewTGanZhi(year, month, day, hour int) *GanZhi{} 月干支计算精确到时
```text
2033-2-3 19H  壬年 月干支:癸丑
2033-2-3 20H  癸年 月干支:甲寅

2022 2 4 4H 壬寅年 月干支:壬寅
2022 2 4 3H 辛丑年 月干支:辛丑
```

- func NewTMGanZhi(year, month, day, hour, min int) *GanZhi{} 干支精确到分钟
```text
y, m, d, h, min := 2022, 5, 21, 9, 21 //立夏: 2022-05-05 20:25:46
y, m, d, h, min = 2022, 5, 21, 9, 23  //小满: 2022-05-21 09:22:24
```

### 示例
```go
package main

import (
	"fmt"
	"time"

	"github.com/Aquarian-Age/xa/pkg/gz"
)

func main() {
	y, m, d, h, min := 2022, 3, 26, 12, 59
	gzo := gz.NewTMGanZhi(y, m, d, h, min)
	fmt.Printf("%s年 %s月 %s日 %s时\n", gzo.Ygz, gzo.Mgz, gzo.Dgz, gzo.Hgz)
	nys := gzo.GetNaYinString()
	fmt.Println(nys) //纳音
	nyd := gzo.NaYin(gzo.Dgz)
	fmt.Printf("日干支纳音:%s\n", nyd)

	lunars, moons := gzo.GetLunar()
	fmt.Printf("%s %s\n", lunars, moons)
	djc := gzo.JianChuDay()
	fmt.Printf("日建除:%s\n", djc)
	fmt.Println(gzo.DiSiHu().DiSiHuString()) //地四户

	t := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	wn := int(t.Weekday())
	riqin := gzo.RiQin(wn)
	fmt.Printf("日禽:%s\n", riqin)
	dhh := gzo.RiHuangHei1()
	fmt.Printf("日黄黑:%s\n", dhh)
	hhh := gzo.ShiHuangHei1()
	fmt.Printf("时辰黄黑:%s\n", hhh)
	jqs := gzo.JieQi()
	fmt.Printf("%s\n", jqs) //当前节气
	yjo := gzo.YueJiangStruct()
	fmt.Printf("月将:%s(%s)\n", yjo.Zhi, yjo.Name)
	fmt.Printf("%s(%s)\n", yjo.ZhongQiName, yjo.ZhongQiT) //中气名称(时间)
	fmt.Printf("天马:%s\n", yjo.TaiChongTianMa(gzo.Hgz))
	m1, m2, m3 := yjo.TianSanMen(gzo.Hgz)
	fmt.Printf("天三门:%s %s %s\n", m1, m2, m3)
	dan, mu := yjo.GuiDengTianMen(gzo.Dgz)
	fmt.Printf("登天门: %s %s\n", dan, mu)
	dsm := yjo.DiSiMen(gzo.Dgz, gzo.Hgz)
	dsms := dsm.DiSiMenString()
	fmt.Println("地私门:", dsms)

	jqarr := gzo.Jq24()
	for i := 0; i < len(jqarr); i++ {
		fmt.Println(jqarr[i])
	}

	// fmt.Println(gzo.Jq24T())
}
```

```text
壬寅年 癸卯月 戊寅日 戊午时
金箔金-金箔金-城头土-天上火
日干支纳音:城头土
阴历: 二月廿四 月相: 0.390827
日建除:闭
地四户: 除在:未 定在:戌 危在:丑 开在:辰
日禽:胃土雉
日黄黑:青龙
时辰黄黑:白虎
春分: 2022-03-20 23:33:14
月将:戌(河魁)
春分(2022-03-20)
天马:亥
天三门:太冲:亥 从魁:巳 小吉:卯
登天门: 申 寅
地私门: 六合:午 太阴:亥 太常:丑
冬至:2021-12-21 23:59:08
小寒:2022-01-05 17:13:53
大寒:2022-01-20 10:38:55
立春:2022-02-04 04:50:36
雨水:2022-02-19 00:42:50
惊蛰:2022-03-05 22:43:34
春分:2022-03-20 23:33:14
清明:2022-04-05 03:20:03
谷雨:2022-04-20 10:24:05
立夏:2022-05-05 20:25:46
小满:2022-05-21 09:22:24
芒种:2022-06-06 00:25:37
夏至:2022-06-21 17:13:40
小暑:2022-07-07 10:37:49
大暑:2022-07-23 04:06:48
立秋:2022-08-07 20:28:58
处暑:2022-08-23 11:15:58
白露:2022-09-07 23:32:07
秋分:2022-09-23 09:03:30
寒露:2022-10-08 15:22:16
霜降:2022-10-23 18:35:30
立冬:2022-11-07 18:45:18
小雪:2022-11-22 16:20:18
大雪:2022-12-07 11:46:04
冬至:2022-12-22 05:48:01
小寒:2023-01-05 23:04:39
大寒:2023-01-20 16:29:21
```