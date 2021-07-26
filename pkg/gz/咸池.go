/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package gz

import (
	"fmt"
	"strings"
)

/*
《淮南子》曰：“日出扶桑，入于咸池。”故五行沐浴之地，名咸池。是取日人之义，万物暗昧之时。
寅午戌卯、已酉丑午、申子辰酉、亥卯未子即长生第二位沐浴之宫是也。
一名败神，一名桃花煞，其神之奸邪淫鄙，如生旺则美容仪，耽酒色，疏财好欢，破散家业，唯务贪淫；
如死绝，落魄不检，言行狡诈，游荡赌博，忘恩失信，私滥奸淫，靡所不为；
与元辰并，更临生旺者，多得匪人为妻；与贵人建禄并，多因油盐酒货得生，或因妇人暗昧之财起家，平生有水厄、痨瘵之疾，累遭遗失暗昧之灾。
此人入命，有破无成，非为吉兆，妇人尤忌之。

咸池非吉煞，日时与水命遇之尤凶。
*/

// GetXianChi  年月日时的咸池位置
func GetXianChi(arg ...string) string {
	var s string
	for i := 0; i < len(arg); i++ {
		sx := XianChi(arg[i])
		s += fmt.Sprintf("%s咸池在:%s ", arg[i], sx)
	}
	return s
}

// XianChi 咸池 传入干支
func XianChi(gz string) string {
	smap := map[string]string{
		"申": "酉", "子": "酉", "辰": "酉",
		"寅": "卯", "午": "卯", "戌": "卯",
		"巳": "午", "酉": "午", "丑": "午",
		"亥": "子", "卯": "子", "未": "子"}
	var s string
	for k, v := range smap {
		if strings.ContainsAny(gz, k) {
			s = v
			break
		}
	}
	return s
}
