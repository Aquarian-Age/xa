package main

import (
	"fmt"
	"time"

	"github.com/Aquarian-Age/xa/pkg/gz"
	"github.com/Aquarian-Age/xa/pkg/x"
)

func main() {

	jzArr := gz.GetJzArr() //生成六十甲子数组
	fmt.Println(jzArr)

	csmap := gz.ChangShengMap("甲子") //干对应的长生map
	fmt.Println(csmap)

	wxsk := gz.GetWXSKS("癸酉") //五行生克
	wxn := gz.Wxsk("癸酉")
	fmt.Println(wxsk, wxn)

	lu := gz.Lu("癸酉") //干支禄
	lu += gz.Lu("庚辰")
	fmt.Println(lu)

	nyall := gz.GetNaYin("辛丑", "壬辰", "庚戌", "丙子")
	fmt.Println(nyall)
	t := time.Now().Local()
	obj := gz.NewGanZhi(t.Year(), int(t.Month()), t.Day(), t.Hour())
	fmt.Println(obj.YGZ, obj.MGZ, obj.DGZ, obj.HGZ)

	xcs := gz.GetXianChi(obj.YGZ, obj.MGZ, obj.DGZ, obj.HGZ) //咸池
	fmt.Println(xcs)

	lns, lns1 := x.LiuNianBiao(t.Year())
	fmt.Printf("流年:%s\n%s\n", lns, lns1)
	//
	jcDay := obj.RiJianChu()
	hhDay := obj.RiHuangHei()
	fmt.Printf("%s月%s日　建除:%s 黄黑:%s\n", obj.MGZ, obj.DGZ, jcDay, hhDay)

	//

	wn := int(t.Weekday())
	if wn == 0 {
		wn = 7
	}
	riQin := obj.RiQin(wn)
	fmt.Printf("%s 日禽:%s\n", t.Format("2006/1/2 15:01:02"), riQin)

}
