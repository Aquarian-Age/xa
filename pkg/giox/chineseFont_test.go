/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package giox

import (
	"fmt"
	"testing"
)

func TestFontX(t *testing.T) {
	fontPath := "/opt/fonts/yahei.ttf"
	th := FontX(fontPath)
	fmt.Printf("TextSize:%v\n", th.TextSize)
	//Nnoto
	th = FontNoto()
	fmt.Printf("font Noto TextSize:%v\n", th.TextSize)
	//Roboto
	th = FontRoboto()
	fmt.Printf("Roboto TextSize:%v\n", th.TextSize)
	/*
		TextSize:16sp
		font Noto TextSize:16sp
		Roboto TextSize:16sp
	*/
}
