package gz

import (
	"math"
	"strings"
)

var (
	shiErZhiYiXiang = map[string]string{
		"子": `代表军人 盗贼 流动性职业 好淫乱 人无主见 也主机密文件
子水为忌克主腿痛 脚痛 破财 家住近水 有水灾之人 有外伤骨折之人
身体方面主 膀胱`,
		"丑": `代表贤士 官人 忠厚老实 性情倔强 丑女 冤仇诅咒 印信文书 金融 包含金融相关的职业
身体方面主 脾 丑也主女性的性器官
室内类像 柜斛`,
		"寅": `文书 财帛 官贵 主清高 主服装 文凭 才华 代表有文化的老人 
身体方面主胸 腰 尾椎 胆 手
室内类像 箱柜`,
		"卯": `代表车船 买卖交易 信息 婚姻媒介 公吏 司机 四肢
卯为用生主利文书 易与女人口舌 
身体方面主 颈椎 肝 指
室内类像 床`,
		"辰": `医生 医巫卜相 又主药物 倔强好斗 貌凶(处死地) 眼小 嗓门大
辰丑:热毒 糖尿病 发烧 泥塘 坑
辰为忌克土有小人 官司 手术 疾病 官灾 车祸 姐妹二婚
身体方面主 胃 肩 胸
室内类像 盆瓮衣包厢`,
		"巳": `主文书信息 惊恐之事 多疑 不讲道理 爱生闷气 
也主轻薄 轻浮 淫乱 生主女人私情
身体方面主 心 面 咽喉
室内类像 炉灶 火炉`,
		"午": `主文书 荣誉 四肢 婚姻信息 书籍 主惊恐灾祸 
又主有血光之事 为主口食  酒食 主皮肤 又代表毒药(或副作用大)
身体方面主 小肠 眼睛
室内类像 衣架笼皮箱`,
		"未": `主人能言善辩 婚姻媒介 代表厨师或与之相关的职业 古人云 未不吃药 免毒气入肠
末为用也主女情 婚外恋 关节炎 末为忌克主斗 婚姻不顺
身体方面主 脾 脊梁
室内类像 中堂 外院`,
		"申": `传送之神 又主军 警 政 法之人 
又主奔波流动 又主凶恶 代表喜欢武术 刀枪棍棒之人或很喜欢现代武器等
工具方面主 尖刀 小刀 锐器
申为忌克主道路不顺 伤灾车祸 丧事 路见武界之人
身体方面主 大肠 经络
室内类像 神词 佛堂`,
		"酉": `主 四肢 论女人主沉静高雅 见水主放荡 淫乱 轻浮(水有制不淫)
金水相逢又见土素质高 做事诡秘 又主金融 镜子 庭院 饰品 
酉:工具方面主 宽刀 大刀 钝器 酉为忌用生主桃花
身体方面主 肺 皮肤 精血
室内类像 凳子 刀剑`,
		"戌": `医巫卜相 僧道 孤寡之人 信佛之人 佛教用品等
主恶人 黑社会 执法者 课内戌多主虚诈 诈骗 引申意为表面上好 实际上比谁都坏
未戌:风湿性病
身体方面主 胃 也主男性的性器官
室内类像 瓮 粮食`,
		"亥": `赏赐 旺相有制主人善良 旺无制主人轻薄 
暗昧之事 暗昧之人 不守本份之人 妓女 暗昧之地 厕所 排水沟等
亥主妄想 坐那什么也不想干 也不知在想什么 有"海"的意思
亥子:身体方面主 血液 肾 淋巴 脑 泪 唾液
亥为忌生主家人有怪异附身 有病人
室内类像 登台 帐布`,
	}
	zhiWuXingMap = map[string]string{
		"亥": "水", "子": "水",
		"寅": "木", "卯": "木",
		"辰": "土", "未": "土", "戌": "土", "丑": "土",
		"巳": "火", "午": "火",
		"申": "金", "酉": "金",
	}
	zhiHuaHeMap = map[string]string{
		"午": "未", "未": "午",
		"子": "丑", "丑": "子",
		"寅": "亥", "亥": "寅",
		"卯": "戌", "戌": "卯",
		"辰": "酉", "酉": "辰",
		"巳": "申", "申": "巳",
	}
)

// AliasZhi 日支 1+mod(JD正午+1,12) 这里传入的JD是时间精确到日计算
func AliasZhi(jd float64) string {
	n := int(math.Ceil(math.Mod(jd+1, 12)))
	zhis := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	if n > 11 {
		n -= 12
	}
	return zhis[n]
}

// ZHI 地支
type ZHI string

func Zhi(zhi string) ZHI {
	return ZHI(zhi)
}

// YiXiang 地支意象
func (z ZHI) YiXiang() string {
	return shiErZhiYiXiang[string(z)]
}

// Xing 刑 寅申巳三刑 丑戌未三刑 子卯相刑 辰午酉亥自刑
func (z ZHI) Xing() string {
	xmap := map[string]string{
		"子": "卯", "卯": "子", //子卯相刑多做殃，钱财难近病卧床。

		"寅": "巳", //寅巳相刑最凶恶，见申必定出夭折
		"巳": "申", //巳申相刑大不祥，中年必见伤一场。
		"申": "寅",

		"丑": "戌", //丑戍相刑不为强，房房男子见刑伤
		"戌": "未", //戍未相刑煞气来，三口之家一个亡。
		"未": "丑",

		"酉": "酉",
		"午": "午",
		"辰": "辰",
		"亥": "亥", //辰午酉亥自相刑，长男少女多离娘。
	}
	return xmap[string(z)]
}

// XingHide 刑
func (z ZHI) XingHide() string {
	xmap := map[string]string{
		"子": `无礼之刑 子卯相刑多做殃，钱财难近病卧床。
没气质 眼光高 讲话没礼貌 说话不客气 不随便与人交谈 自视清高 看到不喜欢的人 他就不会去理会对方，
脾气不好 没礼貌 刑中最凶兆 不孝不悌，相妒不睦 剋损母亲 妇人有此刑 翁姑不合 且易损孕。
`,
		"卯": `无礼之刑 子卯相刑多做殃，钱财难近病卧床。
没气质 眼光高 讲话没礼貌 说话不客气 不随便与人交谈 自视清高 看到不喜欢的人 他就不会去理会对方，
脾气不好 没礼貌 刑中最凶兆 不孝不悌，相妒不睦 剋损母亲 妇人有此刑 翁姑不合 且易损孕。
`,

		"寅": `恃势之刑 寅巳相刑最凶恶，见申必定出夭折
做到累死 别人也不会感激你 替别人打江山 没有恩惠 你有十分力 劝你留三分可收尾 
你会嫌人，别人也会嫌你 做事可以做的成功 但只要有一小点做不好 就会被修理 
性情冷酷薄义 易遭陷害及恶事发生 好有此刑 易损孕`,
		"巳": `恃势之刑 巳申相刑大不祥，中年必见伤一场
做到累死 别人也不会感激你 替别人打江山 没有恩惠 你有十分力 劝你留三分可收尾 
你会嫌人，别人也会嫌你 做事可以做的成功 但只要有一小点做不好 就会被修理 
性情冷酷薄义 易遭陷害及恶事发生 好有此刑 易损孕`,
		"申": `恃势之刑 巳申相刑大不祥，中年必见伤一场
做到累死 别人也不会感激你 替别人打江山 没有恩惠 你有十分力 劝你留三分可收尾 
你会嫌人，别人也会嫌你 做事可以做的成功 但只要有一小点做不好 就会被修理 
性情冷酷薄义 易遭陷害及恶事发生 好有此刑 易损孕`,

		"丑": `无恩之刑 丑戍相刑不为强，房房男子见刑伤
太过有自信 过于猛进 易遭挫折 无恻隐之心 刚毅且易罹灾 妇人有此刑 易孤独
自刑之刑 心里的憋闷 不知道要向谁说 找不到对象说 
有话不想跟别人说 有话说不出口 在心里一点一点的累积(尤其是亥月生人者)
明知不可为而为之 长拿石头砸自己的脚 
会想不开 内心憋闷 不知找谁诉说 自卑(亥) 
若在月令更严重 酉午较亥来得轻微 辰的自刑最轻微 `,
		"戌": `无恩之刑 戍未相刑煞气来，三口之家一个亡
太过有自信 过于猛进 易遭挫折 无恻隐之心 刚毅且易罹灾 妇人有此刑 易孤独
自刑之刑 心里的憋闷 不知道要向谁说 找不到对象说 
有话不想跟别人说 有话说不出口 在心里一点一点的累积(尤其是亥月生人者)
明知不可为而为之 长拿石头砸自己的脚 
会想不开 内心憋闷 不知找谁诉说 自卑(亥) 
若在月令更严重 酉午较亥来得轻微 辰的自刑最轻微 `,
		"未": `无恩之刑 戍未相刑煞气来，三口之家一个亡
太过有自信 过于猛进 易遭挫折 无恻隐之心 刚毅且易罹灾 妇人有此刑 易孤独
自刑之刑 心里的憋闷 不知道要向谁说 找不到对象说 
有话不想跟别人说 有话说不出口 在心里一点一点的累积(尤其是亥月生人者)
明知不可为而为之 长拿石头砸自己的脚 
会想不开 内心憋闷 不知找谁诉说 自卑(亥) 
若在月令更严重 酉午较亥来得轻微 辰的自刑最轻微 `,

		"酉": `辰午酉亥自相刑，长男少女多离娘
讲义气 较鸡婆 遇到懒散的人或者不讲义气不讲理的人 他会生气  太过热情变成忧虑 `,
		"午": `辰午酉亥自相刑，长男少女多离娘
好胜心强 不喜欢别人用话刺激他 个性极端 没耐心 健忘 
属马者 喜欢别人拍他马屁 所以要说好听的 要好好沟通 `,
		"辰": `辰午酉亥自相刑，长男少女多离娘
固执 有原则 不喜欢别人左右他 喜欢独立 作老大 怀才不遇 做事有头无尾`,
		"亥": `辰午酉亥自相刑，长男少女多离娘
聪明 智慧 明理 有事不说 已有自杀倾向`, //。
	}
	return xmap[string(z)]
}

// Chong 冲
//子午相冲、丑未相冲、寅申相冲、卯酉相冲、辰戌相冲、巳亥相冲、
func (z ZHI) Chong() string {
	xmap := map[string]string{
		"子": `午`, "午": `子`,
		"丑": `未`, "未": `丑`,
		"寅": `申`, "申": `寅`,
		"卯": `酉`, "酉": `卯`,
		"辰": `戌`, "戌": `辰`,
		"巳": `亥`, "亥": `巳`}
	return xmap[string(z)]
}

// ChongHide 冲
func (z ZHI) ChongHide() string {
	xmap := map[string]string{
		"子": `子午相冲难消闲，奔波忙碌空一场。
身不安 水灾火光泪涟涟 问病身体总不好 主脑血管 血栓之类 水火不容 
情绪不稳定 脾气不好 个性不好 个性极端 人缘很好 异性缘佳 
较有发疯几率 脑神经衰弱 的现象 子午的人通常都很漂亮 
`,
		"午": `子午相冲难消闲，奔波忙碌空一场。
身不安 水灾火光泪涟涟 问病身体总不好 主脑血管 血栓之类 水火不容 
情绪不稳定 脾气不好 个性不好 个性极端 人缘很好 异性缘佳 
较有发疯几率 脑神经衰弱 的现象 子午的人通常都很漂亮 
`,
		"丑": `丑未相冲不看强，少男奔波老母伤。
冤仇诅咒 语言冲突 不打仗 爱追根究低 
主观意识强烈 易赔钱 财库冲开 开销大 
爱问 钻牛角尖 较会跟邻居吵架 女命易流产 
`,
		"未": `丑未相冲不看强，少男奔波老母伤。
冤仇诅咒 语言冲突 不打仗 爱追根究低 
主观意识强烈 易赔钱 财库冲开 开销大 
爱问 钻牛角尖 较会跟邻居吵架 女命易流产 
`,
		"寅": `寅申相冲两受伤，中男少妇不做良。
更改计划 路途 呼唤往来 忙碌 劳碌命 开车很快 较会走大马路 有车关 易生车祸 
六亲较无缘 一辈子靠自己 锐利 人缘好 异性缘佳 心性不定 （桃花冲）阴易近身
`,
		"申": `寅申相冲两受伤，中男少妇不做良。
更改计划 路途 呼唤往来 忙碌 劳碌命 开车很快 较会走大马路 有车关 易生车祸 
六亲较无缘 一辈子靠自己 锐利 人缘好 异性缘佳 心性不定 （桃花冲）阴易近身
`,
		"卯": `卯酉相冲多做殃，钱财难近病卧床。
`,
		"酉": `卯酉相冲多做殃，钱财难近病卧床。
`,
		"辰": `辰戍相冲大不祥，奴仆逃亡女做娼
争讼斗打到一起 辰与戌属现金 库冲库 撞到事业宫的话 投资会失败 不服输 自圆其说 自找台阶 
喜做老大 脾气不好 理由多 会将错就错 归错於别人 做事野心大 开销大 须注意婚姻问题
`,
		"戌": `辰戍相冲大不祥，奴仆逃亡女做娼
争讼斗打到一起 辰与戌属现金 库冲库 撞到事业宫的话 投资会失败 不服输 自圆其说 自找台阶 
喜做老大 脾气不好 理由多 会将错就错 归错於别人 做事野心大 开销大 须注意婚姻问题
`,
		"巳": `巳亥相冲不相当，忙忙碌碌奔它乡。
为气（生气）索（索要）辩才无碍 口才佳 很会辩 爱聊天 常常祸从口出追根究低 较会钻小卷 有车关
`,
		"亥": `巳亥相冲不相当，忙忙碌碌奔它乡。
为气（生气）索（索要）辩才无碍 口才佳 很会辩 爱聊天 常常祸从口出追根究低 较会钻小卷 有车关
`}
	return xmap[string(z)]
}

// Po 破
func (z ZHI) Po() string {
	xmap := map[string]string{
		"子": `酉`,
		"丑": `辰`,
		"寅": `亥`,
		"卯": `午`,
		"辰": `丑`,
		"巳": `申`,
		"午": `卯`,
		"未": `戌`,
		"申": `巳`,
		"酉": `子`,
		"戌": `未`,
		"亥": `寅`}
	return xmap[string(z)]
}

// PoHide 破
func (z ZHI) PoHide() string {
	xmap := map[string]string{
		"子": `酉子相破大不良，少小卧床闹疾荒
经常面临经济困难 
`,
		"丑": `丑辰相破女占先，家中男女乱花钱
六亲难谐 男女多见不睦
`,
		"寅": `亥寅相破不吉祥，口舌事非多做殃
为人优柔寡断 耳根轻 简单被说服
`,
		"卯": `午卯相破家业败，男盗离乱女做娼
主门户破财
`,
		"辰": `丑辰相破女占先，家中男女乱花钱
六亲难谐 男女多见不睦
`,
		"巳": `巳申相破又带合，离离散散见孤单
仿照力强 办事拿不定主意
`,
		"午": `午卯相破家业败，男盗离乱女做娼
主门户破财
`,
		"未": `戍未相破主伤人，诸事难安伤六親
为人有城府
`,
		"申": `巳申相破又带合，离离散散见孤单
仿照力强 办事拿不定主意
`,
		"酉": `酉子相破大不良，少小卧床闹疾荒
经常面临经济困难 
`,
		"戌": `戍未相破主伤人，诸事难安伤六親
为人有城府
`,
		"亥": `亥寅相破不吉祥，口舌事非多做殃
为人优柔寡断 耳根轻 简单被说服
`}
	return xmap[string(z)]
}

// Hai 害
func (z ZHI) Hai() string {
	xmap := map[string]string{
		"子": `未`,
		"丑": `午`,
		"寅": `巳`,
		"卯": `辰`,
		"辰": `卯`,
		"巳": `寅`,
		"午": `丑`,
		"未": `子`,
		"申": `亥`,
		"酉": `戌`,
		"戌": `酉`,
		"亥": `申`,
	}
	return xmap[string(z)]
}

// HaiHide 害
func (z ZHI) HaiHide() string {
	xmap := map[string]string{
		"子": `子未相害口舌生，男人在外不顺情
个性极端 容易犯小人 易换工作 貌合神离 无话可说 会要求对方 （最严重的害又称天地害南北害）
`,
		"丑": `丑午相害不久长，吵吵闹闹去它乡
耐性差 容易生气 貌合神离 
`,
		"寅": `寅巳相害煞在行，小人口舌不安宁
是非多 无恩情 （人情）易犯小人 冷眼旁观的态度 属驿马害 辩才无碍 如果离婚 也可能同住一屋檐下 
`,
		"卯": `卯辰相害不相当，同包兄弟也骂娘
本身要注意 易遭周边亲人相害 杀伤力很大 
好朋友拖后腿 兄弟无缘 手足无助 要他好 反而害他 遇亲近的人反驳力越大 
`,
		"辰": `卯辰相害不相当，同包兄弟也骂娘
本身要注意 易遭周边亲人相害 杀伤力很大 
好朋友拖后腿 兄弟无缘 手足无助 要他好 反而害他 遇亲近的人反驳力越大 
`,
		"巳": `寅巳相害煞在行，小人口舌不安宁
是非多 无恩情 （人情）易犯小人 冷眼旁观的态度 属驿马害 辩才无碍 如果离婚 也可能同住一屋檐下 
`,
		"午": `丑午相害不久长，吵吵闹闹去它乡
耐性差 容易生气 貌合神离 
`,
		"未": `子未相害口舌生，男人在外不顺情
个性极端 容易犯小人 易换工作 貌合神离 无话可说 会要求对方 （最严重的害又称天地害南北害）
`,
		"申": `亥申相害闹殃殃，口舌事非大不祥
是非多 无恩情 易犯小人 （比喻相见不如还念 相见就吵 不见就念）属驿马害
`,
		"酉": `戍酉相害小人多，如若加卯主口舌
与卯辰害相似 容易被近亲戏弄 （鸡犬不宁 哭笑不得 离婚几率高 ）
`,
		"戌": `戍酉相害小人多，如若加卯主口舌
与卯辰害相似 容易被近亲戏弄 （鸡犬不宁 哭笑不得 离婚几率高 ）
`,
		"亥": `亥申相害闹殃殃，口舌事非大不祥
是非多 无恩情 易犯小人 （比喻相见不如还念 相见就吵 不见就念）属驿马害
`,
	}
	return xmap[string(z)]
}

// He 合
func (z ZHI) He() string {
	xmap := map[string]string{
		"子": `丑`,
		"丑": `子`,
		"寅": `亥`,
		"卯": `戌`,
		"辰": `酉`,
		"巳": `申`,
		"午": `未`,
		"未": `午`,
		"申": `巳`,
		"酉": `辰`,
		"戌": `卯`,
		"亥": `寅`}
	return xmap[string(z)]
}

// HeHide 合
func (z ZHI) HeHide() string {
	xmap := map[string]string{
		"子": `子丑相合不为灾，老老少少卧起来
逆合 凑合 问事失好后坏 问婚姻两人凑合过 不美满 也只违背意愿的遵从 不情愿 
子与丑合上方老人易患肾病 出服毒之人
`,
		"丑": `子丑相合不为灾，老老少少卧起来
逆合 凑合 问事失好后坏 问婚姻两人凑合过 不美满 也只违背意愿的遵从 不情愿 
子与丑合上方老人易患肾病 出服毒之人
`,
		"寅": `寅亥相合也不祥，大大小小懒洋洋
破合 寅把亥破了 问事好中有坏 午与寅合 上方老人有少亡 有残疾之人 
`,
		"卯": `卯戍相合不会伤，中男老父不做良
淫合 自焚之象 卯未草木 戌未火库 卯去合戌进了火库 岂不自焚 
卯木太冲 主无自控能力  意志不坚定 缺乏主见 
问婚主未婚同居 卯与戌合 上方老人或姐妹有二婚之人
`,
		"辰": `辰酉相合煞无用，少男懒堕要花钱
暗合 表私下定 不公开 
`,
		"巳": `巳申相合两相应，卷恋情郎主消磨
刑合 有矛盾 意见大 象夫妻 在一起就吵 分开还想 
巳与申合 六亲易有驼背之人 下代出兵士之人
`,
		"午": `午未相合为母女，家中贫穷不叫屈
明合 主公开 实心实意 毫不遮掩 午与未合 上方老人有重婚 下代必出文人 
`,
		"未": `午未相合为母女，家中贫穷不叫屈
明合 主公开 实心实意 毫不遮掩 午与未合 上方老人有重婚 下代必出文人 
`,
		"申": `巳申相合两相应，卷恋情郎主消磨
刑合 有矛盾 意见大 象夫妻 在一起就吵 分开还想 
巳与申合 六亲易有驼背之人 下代出兵士之人
`,
		"酉": `辰酉相合煞无用，少男懒堕要花钱
暗合 表私下定 不公开 
`,
		"戌": `卯戍相合不会伤，中男老父不做良
淫合 自焚之象 卯未草木 戌未火库 卯去合戌进了火库 岂不自焚 
卯木太冲 主无自控能力  意志不坚定 缺乏主见 
问婚主未婚同居 卯与戌合 上方老人或姐妹有二婚之人
`,
		"亥": `寅亥相合也不祥，大大小小懒洋洋
破合 寅把亥破了 问事好中有坏 午与寅合 上方老人有少亡 有残疾之人 
`}
	return xmap[string(z)]
}

// WuXing 五行
func (z ZHI) WuXing() string {
	return zhiWuXingMap[string(z)]
}

// WuXingShengKe 五行生克
func (z ZHI) WuXingShengKe() (string, string) {
	wx := zhiWuXingMap[string(z)]
	return WuXingShengKe(wx)
}

// HuaHe 六合 子丑合化土，寅亥合化木，卯戍合化火，辰酉合化金，巳申合化水，午与未合化土。
func (z ZHI) HuaHe(zhi string) (bool, string) {
	huaHeMap := map[string]string{
		"午": "土", "未": "土",
		"子": "土", "丑": "土",
		"寅": "木", "亥": "木",
		"卯": "火", "戌": "火",
		"辰": "金", "酉": "金",
		"巳": "水", "申": "水",
	}
	heZhi := zhiHuaHeMap[string(z)]
	if strings.EqualFold(zhi, heZhi) {
		return true, huaHeMap[string(z)]
	}
	return false, ""
}
