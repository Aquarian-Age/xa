package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("now: ", time.Now().Local())
	robotgo.TypeStr("Automatically press the space bar")
	for {
		select {
		case <-time.After(19 * time.Minute):
			robotgo.KeyTap("space")
			//_, _ := robotgo.ReadAll()
			t := time.Now().Local()
			fmt.Printf("%v\n", t)
		}
		//robotgo.KeyTap("k","command")//win+k按键
	}
}
