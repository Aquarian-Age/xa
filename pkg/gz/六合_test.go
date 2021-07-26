/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package gz

import (
	"strings"
	"testing"
)

func TestLiuHe(t *testing.T) {
	ygz := "甲子"
	lh := "丑"
	s := LiuHe(ygz)
	if !strings.EqualFold(s, lh) {
		t.Errorf("func LiuHe(%s)=%s want:%s", ygz, s, lh)
	}
}
