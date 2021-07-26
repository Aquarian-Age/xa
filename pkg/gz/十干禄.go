/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package gz

import "strings"

/*
   甲禄在寅　乙禄在卯　甲辰旬寅卯空　甲辰乙巳为无禄
   庚禄在申　辛禄在酉　甲戌旬申酉空　庚辰辛巳为无禄
   丙戊禄在巳　　　　　甲午旬巳空　　丙申戊戌为无禄
   丁己禄在午        甲申旬午空　　丁亥己丑为无禄
   壬禄在亥　　　　　　甲子旬亥空　　壬申为无禄
   癸禄在子          甲寅旬子空　　癸亥为无禄
*/
//十干禄位
func Lu(gz string) string {
	lmap := map[string]string{"甲": "寅", "乙": "卯", "丙": "巳", "丁": "午", "戊": "巳", "己": "午", "庚": "申", "辛": "酉", "壬": "亥", "癸": "子"}
	arr := []string{"甲辰", "乙巳", "庚辰", "辛巳", "丙申", "戊戌", "丁亥", "己丑", "壬申", "癸亥"} //无禄干支
	var s string

	for i := 0; i < len(arr); i++ {
		if strings.EqualFold(gz, arr[i]) {
			s = gz + "无禄"
			break
		} else {
			for k, v := range lmap {
				if strings.ContainsAny(gz, k) {
					s = gz + "-->" + k + "禄在" + v
					break
				}
			}
		}
	}
	return s
}
