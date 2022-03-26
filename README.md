### 干支计算

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
fmt.Printf("%s年 %s月 %s日 %s时\n", gzo.YGZ, gzo.MGZ, gzo.DGZ, gzo.HGZ)
nys := gzo.GetNaYinString()
fmt.Println(nys) //纳音
nyd := gzo.NaYin(gzo.DGZ)
fmt.Printf("日干支纳音:%s\n", nyd)

lunars, moons := gzo.GetLunar()
fmt.Printf("%s %s\n", lunars, moons)
djc := gzo.JianChuDay()
fmt.Printf("日建除:%s\n", djc)
t := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
wn := int(t.Weekday())
riqin := gzo.RiQin(wn)
fmt.Printf("日禽:%s\n", riqin)
dhh := gzo.ShiHuangHei1()
fmt.Printf("日黄黑:%s\n", dhh)
hhh := gzo.RiHuangHei1()
fmt.Printf("时辰黄黑:%s\n", hhh)
jqs := gzo.JieQi()
fmt.Printf("%s\n", jqs) //当前节气
yjo := gzo.YueJiangStruct()
fmt.Printf("月将:%s(%s)\n", yjo.Zhi, yjo.Name)
fmt.Printf("%s(%s)\n", yjo.ZhongQiName, yjo.ZhongQiT) //中气名称(时间)
fmt.Printf("天马:%s\n", yjo.TaiChongTianMa(gzo.HGZ))
}
```

