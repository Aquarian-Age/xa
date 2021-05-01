package main

import (
	"fmt"

	"github.com/Aquarian-Age/xa/pkg/gz"
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

	nyall:= gz.GetNaYin("辛丑","壬辰","庚戌","丙子")
	fmt.Println(nyall)

}
