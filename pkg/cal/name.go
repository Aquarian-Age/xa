package cal

import "time"

var (
	//二十四节气名称
	Jmc = []string{
		"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
		"春分", "清明", "谷雨", "立夏", "小满", "芒种",
		"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
		"秋分", "寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
	}
	//十天干 甲1 ...癸10
	Gan = [11]string{"err", "甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	//十二地支 子1 ...亥12
	Zhi = [13]string{"err", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	//月名称
	Ymc = []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二"}
	//农历日名称
	Rmc = []string{
		"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十",
		"十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十",
		"廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十"}
	//二十八宿 东-->北-->西-->南逆时针走
	XingSu28 = [28]string{
		"角木蛟", "亢金龙", "氐土貉", "房日兔", "心月狐", "尾火虎", "箕水豹", //东方青龙
		"斗木獬", "牛金牛", "女土蝠", "虚日鼠", "危月燕", "室火猪", "壁水貐", //北方玄武
		"奎木狼", "娄金狗", "胃土雉", "昴日鸡", "毕月乌", "觜火猴", "参水猿", //西方白虎
		"井木犴", "鬼金羊", "柳土獐", "星日马", "张月鹿", "翼火蛇", "轸水蚓", //南方朱雀
	}
)

//干支信息
type Cal struct {
	YearGZ  string `json:"year_gz"`
	MonthGZ string `json:"month_gz"`
	DayGZ   string `json:"day_gz"`
	HourGZ  string `json:"hour_gz"`
}

//不含闰月年份的数据
type Lunar struct {
	LY   int    `json:"y"`
	LM   int    `json:"m"`
	LRmc string `json:"rmc"` //农历日名称
	Ydx  string `json:"ydx"` //月大小

	LeapM   int    `json:"leap_m"`
	LeapRmc string `json:"leap_rmc"` //闰月日名称
	LeapYdx string `json:"leap_ydx"`

	Sday     int    `json:"sday"`
	LeapSday int    `json:"leap_sday"`
	Week     string `json:"week"`
}

//月历 这里是农历 初始是农历初一到农历本月结束
type YueLi struct {
	LunarD []string `json:"ld"`
	SolarD []string `json:"sd"`
	GzD    []string `json:"gzd"`
	RiQin  []string `json:"ri_qin"`
	QiShaB []bool   `json:"qi_sha_b"`
}

//闰月月历
type LeapYl struct {
	Leapl  []string `json:"leapl"`
	Leaps  []string `json:"leaps"`
	Leapgz []string `json:"leapgz"`
}

//从上年冬至到下年立春的节气名称　节气时间戳
// 0:冬至(上一年) 1:小寒 2:大寒 3:立春(本年) 4:雨水 惊蛰 春分 清明 谷雨 立夏 小满 11芒种
//夏至 小暑 大暑 立秋 处暑 白露 秋分 寒露 霜降 立冬 小雪 23大雪
//冬至 小寒 大寒 27:立春(下一年)
type JQArr struct {
	Name []string    //节气名称
	Time []time.Time //时间精确到秒钟
}

//朔望JD数据
type ShuoWangF struct {
	ShuoF      float64 `json:"shuof"`
	ShangXianF float64 `json:"shang_xianf"`
	WangF      float64 `json:"wangf"`
	XiaXianF   float64 `json:"xia_xianf"`
}

//朔 上弦 望 下弦 时间戳
type ShuoWangT struct {
	ShuoT      time.Time `json:"shuo_t"`
	ShangXianT time.Time `json:"shang_xian_t"`
	WangT      time.Time `json:"wang_t"`
	XiaXianT   time.Time `json:"xia_xian_t"`
}

//朔 上弦 望 下弦 时间戳 string格式
type ShuoWantTS struct {
	ShuoTS      string `json:"shuo_ts"`
	ShangXianTS string `json:"shang_xian_ts"`
	WangTS      string `json:"wang_ts"`
	XiaXianTS   string `json:"xia_xian_ts"`
}

//农历本月朔(初一日)和本月节
type MoonShuoTJ12T struct {
	MoonShuoT time.Time `json:"moon_shuo"` //农历月朔时间戳
	MoonJieT  time.Time `json:"moon_jie"`  //本月节时间戳
}
