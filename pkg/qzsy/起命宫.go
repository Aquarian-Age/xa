/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package qzsy

import (
	"strings"

	"github.com/Aquarian-Age/xa/pkg/pub"
)

//传入月干支　时辰干支　返回命宫位
//起命宫　方法等同十八飞星策天紫微起命宫　这里月份使用干支历即正月建寅　寅月为正月　卯月为二月....
func MingGong(mgz, hgz string) string {
	//k月支 v太阳
	mz := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"} //月支
	sz := []string{"子", "亥", "戌", "酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑"} //太阳
	var s string
	for i := 0; i < len(mz); i++ {
		if strings.ContainsAny(mgz, mz[i]) {
			sun := sz[i]
			arr := pub.SortArr(pub.GetZhiS(hgz), mz)
			xArr := pub.SortArr(sun, mz)
			for j := 0; j < len(arr); j++ {
				if strings.EqualFold(arr[j], "卯") {
					s = xArr[j]
					break
				}
			}
			break
		}
	}
	return s
}
