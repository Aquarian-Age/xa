### 干支计算

- func NewGanZhi(year, month, day, hour int) *GanZhi{} 月干支计算精确到日
```text
例如： 2022年2月4日 4:50分立春.
传入: 2022, 2, 4, 0 显示: 壬寅 壬寅 戊子 壬子
```

- func NewTGanZhi(year, month, day, hour int) *GanZhi{} 月干支计算精确到时
```text
2022 2 4 4H 壬寅年 月干支:壬寅
2022 2 4 3H 辛丑年 月干支:辛丑
```

- [简单农历](https://github.com/Aquarian-Age/ccal/releases/tag/chineseLunar)

- [28宿日历](https://github.com/Aquarian-Age/ccal/releases/tag/28%E5%AE%BF%E6%97%A5%E5%8E%86)
