/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package gz

import (
	"reflect"
	"testing"
)

func TestXianChi(t *testing.T) {
	want := map[string]string{
		"子": "酉", "丑": "午", "寅": "卯", "卯": "子", "辰": "酉", "巳": "午",
		"午": "卯", "未": "子", "申": "酉", "酉": "午", "戌": "卯", "亥": "子",
	}

	var smap = make(map[string]string)

	for i := 1; i < len(Zhi); i++ {
		s := XianChi(Zhi[i])
		smap[Zhi[i]] = s
	}

	if !reflect.DeepEqual(smap, want) {
		t.Errorf("func XianChi()=%v want:%v", smap, want)
	}
}
