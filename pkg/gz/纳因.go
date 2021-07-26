/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package gz

import "strings"

//年月日时 的干支纳音
func GetNaYin(gzx ...string) string {
	var arr []string
	for i := 0; i < len(gzx); i++ {
		nyx := NaYin(gzx[i])
		s := nyx[gzx[i]]
		arr = append(arr, s)
	}
	return strings.Join(arr, "-")
}

//干支纳音
func NaYin(gz string) (nywx map[string]string) {
	arr := []string{
		"海中金", "炉中火", "大林木", "路旁土",
		"剑锋金", "山头火", "涧下水", "城头土",
		"白蜡金", "杨柳木", "泉中水", "屋上土",
		"霹雳火", "松柏木", "长流水", //14
		"沙中金", "山下火", "平地木", "壁上土",
		"金箔金", "覆灯火", "天河水", "大驿土",
		"钗钏金", "桑柘木", "大溪水", "沙中土",
		"天上火", "石榴木", "大海水", //29
	}
	//六十甲子
	jz := GetJzArr()
	//纳音
	nywx = make(map[string]string)

	//六十甲子数组索引号除以2求商　商为当前干支的纳音索引数字
	for i := 0; i < len(jz); i++ {
		if strings.EqualFold(gz, jz[i]) {
			index := i / 2
			nywx[jz[i]] = arr[index]
		}
	}
	return
}
