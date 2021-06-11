package main

import (
	"fmt"

	"github.com/Aquarian-Age/xa/pkg/gz"
	"github.com/Aquarian-Age/xa/pkg/x"
)

func main() {

	jzArr := gz.GetJzArr() //生成六十甲子数组
	fmt.Println(jzArr)

	csmap := gz.ChangSheng("甲子") //干对应的长生map
	fmt.Println(csmap)

	wxsk := gz.GetWXSKS("癸酉") //五行生克
	wxn := gz.Wxsk("癸酉")
	fmt.Println(wxsk, wxn)

	lu := gz.Lu("癸酉") //干支禄
	lu += gz.Lu("庚辰")
	fmt.Println(lu)

	nyall := gz.GetNaYin("辛丑", "壬辰", "庚戌", "丙子")
	fmt.Println(nyall)

	gzs := gz.NewGanZhi(2021, 5, 8, 8)
	fmt.Println(gzs.YGZ, gzs.MGZ, gzs.DGZ, gzs.HGZ) //辛丑 癸巳 丙辰 壬辰

	xcs := gz.GetXianChi(gzs.YGZ, gzs.MGZ, gzs.DGZ, gzs.HGZ) //咸池
	fmt.Println(xcs)

	lns, lns1 := x.LiuNianBiao(1990)
	fmt.Printf("流年:%s\n%s\n", lns, lns1)
}
