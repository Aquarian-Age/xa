package gz

import (
	"fmt"
	"testing"
)

func TestChangShengZhi(t *testing.T) {
	gan, name := "甲", "养"
	gan, name = "己", "临官"
	s := ChangShengZhi(gan, name)
	fmt.Printf("%s的%s地在%s\n", gan, name, s)
}
