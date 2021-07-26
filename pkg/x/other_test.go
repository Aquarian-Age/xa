/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

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
