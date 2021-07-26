/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package gz

import (
	"fmt"
	"testing"
)

/* [子水生甲木 乙木克丑土 寅木生丙火 卯木生丁火 戊土比和辰土 巳火生己土 午火克庚金 未土生辛金 申金生壬水 酉金生癸水 甲木克戌土 亥水生乙木 子水克丙火 丁火生丑土 寅木克戊土 卯木克己土 辰土生庚金 巳火克辛金 壬水克午火 未土克癸水 申金克甲木 酉金克乙木 丙火生戌土 亥水克丁火 戊土克子水 己土比和丑土 庚金克寅木 辛金克卯木 辰土克壬水 癸水克巳火 甲木生午火 乙木克未土 丙火克申金 丁火克酉金 戊土比和戌土 己土克亥水 庚金生子水 丑土生辛金 壬水生寅木 癸水生卯木 甲木克辰土 乙木生巳火 丙火比和午火 丁火生未土 戊土生申金 己土生酉金 戌土生庚金 辛金生亥水 壬水比和子水 丑土克癸水 甲木比和寅木 乙木比和卯木 丙火生辰土 丁火比和巳火 午火生戊土 己土比和未土 庚金比和申金 辛金比和酉金 戌土克壬水 癸水比和亥水] */
func TestWuXingShengKe(t *testing.T) {
	gzarr := GetJzArr()
	var arrI []string
	for i := 0; i < len(gzarr); i++ {
		arrI = append(arrI, GetWXSKS(gzarr[i]))
	}
	fmt.Println(arrI)
}
