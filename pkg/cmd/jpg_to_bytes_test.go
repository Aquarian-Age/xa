package cmd

import (
	"fmt"
	"testing"
)

func TestJpgToBytes(t *testing.T) {
	name := "../calendar.ico"
	err, b := JpgToBytes(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", b)
}
