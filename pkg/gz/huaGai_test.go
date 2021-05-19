package gz

import (
	"strings"
	"testing"
)

// go test --run TestGetHuaGai
func TestGetHuaGai(t *testing.T) {
	want := []string{
		"辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未",
	}
	arr := GetJzArr()
	for i := 0; i < len(arr); i++ {
		s := huaGai(arr[i])
		for j := 0; i < len(want); j++ {
			if i == j {
				if !strings.EqualFold(s, want[j]) {
					t.Errorf("func (%s)=%s want:%s", arr[i], s, want[j])
				}
				break
			}
		}
	}
}
