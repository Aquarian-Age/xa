package gz

import (
	"fmt"
	"testing"
)

func TestGetRiJianChu(t *testing.T) {
	jz60 := GetJzArr()
	for i := 0; i < len(jz60); i++ {
		mgz := "甲寅"
		mgz = "乙卯"
		dgz := jz60[i]
		s := GetRiJianChu(mgz, dgz)
		fmt.Printf("%s-%s %s\n", mgz, dgz, s)
	}

}
