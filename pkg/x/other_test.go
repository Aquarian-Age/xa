package x

import (
	"fmt"
	"testing"
)

func TestLiuNianBiao(t *testing.T) {
	by := 1990
	s, s1 := LiuNianBiao(by)
	fmt.Println(s, "\n", s1)
}
