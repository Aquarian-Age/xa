package gz

import (
	"fmt"
	"github.com/starainrt/astro/calendar"
	"testing"
	"time"
)

func TestAliasStarName(t *testing.T) {
	tx := time.Now().Local()
	for i := 1; i <= 31; i++ {
		tx = time.Date(tx.Year(), tx.Month(), i, 0, 0, 0, 0, time.Local)
		jd := calendar.Date2JDE(tx)
		weekn := WeekNumber(jd)
		aliaszhi := AliasZhi(jd)
		name := AliasStarName(weekn, aliaszhi)
		fmt.Println(tx.String()[:10], weekn, aliaszhi, name)
	}
}

/*
2022-02-01 2 酉 觜火猴
2022-02-02 3 戌 参水猿
2022-02-03 4 亥 井木犴
2022-02-04 5 子 鬼金羊
2022-02-05 6 丑 柳土獐
2022-02-06 7 寅 星日马
2022-02-07 1 卯 张月鹿
2022-02-08 2 辰 翼火蛇
2022-02-09 3 巳 轸水蚓
2022-02-10 4 午 角木蛟
2022-02-11 5 未 亢金龙
2022-02-12 6 申 氐土貉
2022-02-13 7 酉 房日兔
2022-02-14 1 戌 心月狐
2022-02-15 2 亥 尾火虎
2022-02-16 3 子 箕水豹
2022-02-17 4 丑 斗木獬
2022-02-18 5 寅 牛金牛
2022-02-19 6 卯 女土蝠
2022-02-20 7 辰 虚日鼠
2022-02-21 1 巳 危月燕
2022-02-22 2 午 室火猪
2022-02-23 3 未 壁水貐
2022-02-24 4 申 奎木狼
2022-02-25 5 酉 娄金狗
2022-02-26 6 戌 胃土雉
2022-02-27 7 亥 昴日鸡
2022-02-28 1 子 毕月乌
2022-03-01 2 丑 觜火猴
*/
