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
 
```go
package main

import (
"fmt"
"time"

"github.com/Aquarian-Age/xa/pkg/gz"
)

func main() {
y, m, d, h := 2022, 3, 26, 12
gzo := gz.NewGanZhi(y, m, d, h)
fmt.Printf("%s年 %s月 %s日 %s时\n", gzo.Ygz, gzo.Mgz, gzo.Dgz, gzo.Hgz)
nys := gzo.GetNaYinString()
fmt.Println(nys) //纳音
nyd := gzo.NaYin(gzo.Dgz)
fmt.Printf("日干支纳音:%s\n", nyd)

lunars, moons := gzo.GetLunar()
fmt.Printf("%s %s\n", lunars, moons)
djc := gzo.JianChuDay()
fmt.Printf("日建除:%s\n", djc)
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
}
```

