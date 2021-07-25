package cal

import "strings"

//地支属性
type ZHI struct {
	Name      string `json:"name"`       //名称
	FangXiang string `json:"fang_xiang"` //方向(后天八卦)
	WuXing    string `json:"wu_xing"`    //五行属性
	Gua       string `json:"gua"`        //八卦名称
	YinYang   bool   `json:"yin_yang"`   //阴阳 true:阳　false:阴
	ShengXiao string `json:"sheng_xiao"` //十二生肖
}

//地支属性信息
func NewZHI(zhi string) (z *ZHI) {
	z = new(ZHI)
	//地支
	zhiArr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	//生肖
	sx := []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	//八方
	bf := []string{"北", "东北", "东", "东南", "南", "西南", "西", "西北"}
	//八卦
	bg := []string{"坎", "艮", "震", "巽", "离", "坤", "兑", "乾"}
	//五行
	wx := []string{"木", "火", "土", "金", "水"}

	switch zhi {
	case zhiArr[0]: //子
		z = &ZHI{
			Name:      Zhi[0],
			FangXiang: bf[0],
			WuXing:    wx[4],
			Gua:       bg[0],
			YinYang:   true,
			ShengXiao: sx[0],
		}
	case zhiArr[1]: //丑
		z = &ZHI{
			Name:      Zhi[1],
			FangXiang: bf[1],
			WuXing:    wx[2],
			Gua:       bg[1],
			YinYang:   false,
			ShengXiao: sx[1],
		}
	case zhiArr[2]: //寅
		z = &ZHI{
			Name:      Zhi[2],
			FangXiang: bf[1],
			WuXing:    wx[0],
			Gua:       bg[1],
			YinYang:   true,
			ShengXiao: sx[2],
		}
	case zhiArr[3]: //卯
		z = &ZHI{
			Name:      Zhi[3],
			FangXiang: bf[2],
			WuXing:    wx[0],
			Gua:       bg[2],
			YinYang:   false,
			ShengXiao: sx[3],
		}
	case zhiArr[4]: //辰土龙
		z = &ZHI{
			Name:      Zhi[4],
			FangXiang: bf[3],
			WuXing:    wx[2],
			Gua:       bg[3],
			YinYang:   true,
			ShengXiao: sx[4],
		}
	case zhiArr[5]: //巳火蛇
		z = &ZHI{
			Name:      Zhi[5],
			FangXiang: bf[3],
			WuXing:    wx[1],
			Gua:       bg[3],
			YinYang:   false,
			ShengXiao: sx[5],
		}
	case zhiArr[6]: //午
		z = &ZHI{
			Name:      Zhi[6],
			FangXiang: bf[4],
			WuXing:    wx[1],
			Gua:       bg[4],
			YinYang:   true,
			ShengXiao: sx[6],
		}
	case zhiArr[7]: //未土羊
		z = &ZHI{
			Name:      Zhi[7],
			FangXiang: bf[5],
			WuXing:    wx[2],
			Gua:       bg[5],
			YinYang:   false,
			ShengXiao: sx[7],
		}
	case zhiArr[8]: //申金猴
		z = &ZHI{
			Name:      Zhi[8],
			FangXiang: bf[5],
			WuXing:    wx[3],
			Gua:       bg[5],
			YinYang:   true,
			ShengXiao: sx[8],
		}
	case zhiArr[9]: //酉金鸡
		z = &ZHI{
			Name:      Zhi[9],
			FangXiang: bf[6],
			WuXing:    wx[3],
			Gua:       bg[6],
			YinYang:   false,
			ShengXiao: sx[9],
		}
	case zhiArr[10]: //戌土狗
		z = &ZHI{
			Name:      Zhi[10],
			FangXiang: bf[7],
			WuXing:    wx[2],
			Gua:       bg[7],
			YinYang:   true,
			ShengXiao: sx[10],
		}
	case zhiArr[11]: //亥水猪
		z = &ZHI{
			Name:      Zhi[11],
			FangXiang: bf[7],
			WuXing:    wx[4],
			Gua:       bg[7],
			YinYang:   false,
			ShengXiao: sx[11],
		}
	}
	return
}

//禄
func Lu(gz string) (lus string) {

	gan := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

	var i int
	for i = 0; i < len(gan); i++ {
		if strings.ContainsAny(gz, gan[i]) {
			break
		}
	}
	switch i {
	case 0: //甲
		lus = zhi[2]
	case 1: //乙
		lus = zhi[3]
	case 2: //丙
		lus = zhi[5]
	case 4: //戊
		lus = zhi[5]
	case 3: //丁
		lus = zhi[6]
	case 5: //己
		lus = zhi[6]
	case 6: //庚
		lus = zhi[8]
	case 7: //辛
		lus = zhi[9]
	case 8: //壬
		lus = zhi[11]
	case 9: //癸
		lus = zhi[0]
	}
	return
}

//五行生克 木火土金水 临位相生隔位相克
//比和n=0 前者生后者n=1 前者克后者n=-1 后者生前者n=2 后者克前者n=-2
func Wxsk(wx1, wx2 string) (n int) {
	wxs := []string{"木", "火", "土", "金", "水", "木", "火"}
	var w1 int
	for i := 0; i < 5; i++ {
		if strings.EqualFold(wx1, wxs[i]) {
			w1 = i
		}
	}
	var w2 int
	for j := 0; j < 5; j++ {
		if strings.EqualFold(wx2, wxs[j]) {
			w2 = j
			break
		}
	}
	for x := 0; x < len(wxs); x++ {
		//前者生/克后者
		k := w1 + 2
		s := w1 + 1
		if strings.EqualFold(wx1, wx2) {
			n = 0
			break
		}
		if strings.EqualFold(wx2, wxs[k]) {
			n = -1
			break
		}
		if strings.EqualFold(wx2, wxs[s]) {
			n = 1
			break
		}
		//后者生/克前者
		rk := w2 + 2
		rs := w2 + 1
		if strings.EqualFold(wx1, wxs[rk]) {
			n = -2
			break
		}
		if strings.EqualFold(wx1, wxs[rs]) {
			n = 2
			break
		}

	}
	return
}

//阴阳贵人诀
func GuiRenJue(dgz string) (yang, yin string) {
	乙己 := []string{"申", "子"} //乙申阳　乙子阴　己申阴　己子阳
	甲戊 := []string{"未", "丑"} //甲未阳　甲丑阴　戊未阴　戊丑阳
	丙丁 := []string{"酉", "亥"}
	壬癸 := []string{"卯", "巳"}
	庚 := []string{"丑", "未"}
	辛 := []string{"寅", "午"}
	gan := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}

	var i int
	for i = 0; i < len(gan); i++ {
		if strings.ContainsAny(dgz, gan[i]) {
			break
		}
	}
	switch i {
	case 0: //甲
		yang = 甲戊[0]
		yin = 甲戊[1]
	case 4: //戊
		yang = 甲戊[1]
		yin = 甲戊[0]
	case 1: //乙
		yang = 乙己[0]
		yin = 乙己[1]
	case 5: //己
		yang = 乙己[1]
		yin = 乙己[0]
	case 2: //丙
		yang = 丙丁[0]
		yin = 丙丁[1]
	case 3: //丁
		yang = 丙丁[1]
		yin = 丙丁[0]
	case 6: //庚
		yang = 庚[0]
		yin = 庚[1]
	case 7: //辛
		yang = 辛[0]
		yin = 辛[1]
	case 8: //壬
		yang = 壬癸[0]
		yin = 壬癸[1]
	case 9: //癸
		yang = 壬癸[1]
		yin = 壬癸[0]
	}
	return
}