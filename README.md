### 干支计算

- 传入阳历年月日时

- func NewGanZhi(year, month, day, hour int) *GanZhi{} 月干支计算精确到日
```text
例如： 2022年2月4日 4:50分立春. 
传入: 2022, 2, 4, 0 显示: 壬寅 壬寅 戊子 壬子
```

- func NewTGanZhi(year, month, day, hour int) *GanZhi{} 月干支计算精确到时
```text
2033-2-3 19H  壬年 月干支:癸丑
2033-2-3 20H  癸年 月干支:甲寅

2022 2 4 4H 壬寅年 月干支:壬寅
2022 2 4 3H 辛丑年 月干支:辛丑
```

- func NewTMGanZhi(year, month, day, hour, min int) *GanZhi{} 干支精确到分钟
```text
y, m, d, h, min := 2022, 5, 21, 9, 21 //立夏: 2022-05-05 20:25:46
y, m, d, h, min = 2022, 5, 21, 9, 23  //小满: 2022-05-21 09:22:24
```

### 示例
![](./calendar-go.svg)

### License

[MIT](http://opensource.org/licenses/MIT)

Copyright (c) 2017-present, Aquarian-Age
