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
寅午戌人在戌 巳酉丑人在丑 申子辰人在辰 亥卯未人在未
凡命值一位者,近贵,多主僧道,术人,印禄同宫,三位俱定,食禄高强之人,三者全于年月日时者,乃妙。(林开五命)
华盖者,以三合之位入墓也,谓如寅午戌属火,戌为墓也,主人多能博识,立性聪明,
又云:华盖即墓杀也,所占颇异者,盖取三命旺相休囚凶神吉神多少,或身克杀,杀克身之类消息断之。(三命论)

华盖为庇荫之清神,主人旷颖神清,性灵恬淡,不较是非,好仙道伎巧事,一生不利财物,
惟与夹贵正印并,则为福清贵,特达,不利权握,日犯克妻,时犯克子,孤介之神也。(烛神经)

华盖要皆主福,若与岁干库同位为两重福,主大贵,三命提要华盖要见库墓冠带之乡,
假令癸卯金人,华盖在未,见丁未为冠带之位,若辛卯木人见之,为库墓之位,见癸未是木人之正印,最为贵格,余皆仿此。(珞琭贵贱格)
*/
func GetHuaGai(arg ...string) string {
	var s string
	for i := 0; i < len(arg); i++ {
		sx := huaGai(arg[i])
		s += fmt.Sprintf("%s化盖在:%s ", arg[i], sx)
	}
	return s
}

func huaGai(gz string) string {
	smap := map[string]string{
		"申": "辰", "子": "辰", "辰": "辰",
		"寅": "戌", "午": "戌", "戌": "戌",
		"巳": "丑", "酉": "丑", "丑": "丑",
		"亥": "未", "卯": "未", "未": "未"}
	var s string
	for k, v := range smap {
		if strings.ContainsAny(gz, k) {
			s = v
			break
		}
	}
	return s
}
