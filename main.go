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

	wxsk:=gz.GetWXSKS("癸酉")//五行生克
	wxn:=gz.Wxsk("癸酉")
	fmt.Println(wxsk,wxn)
}
