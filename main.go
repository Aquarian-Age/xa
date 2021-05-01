package main

import (
	"fmt"

	"github.com/Aquarian-Age/xa/pkg/gz"
)

func main() {

	jzArr := gz.GetJzArr() //生成六十甲子数组
	fmt.Println(jzArr)
}
