package main

const html = `
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style type="text/css">
        #form {
           position: absolute;
            top: 0.5px;
            text-align: center;
        }

        #btn {
           position: absolute;
            top: 50px;
            text-align: center;
        }

        #show {
            position: absolute;
            top: 90px;
            font-size: 150%;
            white-space: pre
        }
        /*右下角显示*/
        .ridgown{
            position: fixed;
            bottom: 3px;
            right: 3px;
            padding: 8px;
            font-size: 150%;
            white-space: pre;
        }
        #formjx {
            position: absolute;
            top: 50px;
            right: 10px;
            text-align: right;
        }
        #form1{
          position: absolute;
            top: 0.5px;
            right: 85px;
            text-align: center;
        }
        #btngzs{
            position: absolute;
            top: 0.5px;
            right: 10px;
        }
        .img{
            position:absolute;
            top:0.1px;
            left: 820px;
        }
        #qimen{
            position: absolute;
            top: 260px;
            left: 230px;
            font-size: 150%;
            white-space: pre
        }
        #tables{
            position: absolute;
            top: 260px;
            text-align: center;
            width: 390px;
        }
    </style>
    <title>农历择日</title>
</head>

<body>

<form id="form">
    <!-- 下拉年 -->
    <select id="yearid" name="year">
        <script type="text/tiscript">
		var t = new Date();
        var select = $(select#yearid);
        for( var i = 1601; i <= 3498; ++i ){
          select.options.append(<option value={i}>{i}年</option>);
            var s = select.text;
            //设置默认显示
            select.value=t.year;
        }
        </script>
    </select>
    <!-- 下拉月份 -->
    <select id="monthid" name="month">
        <option value=" ">农历月</option>
        <script type="text/tiscript">
				var marr=["正月(寅)", "二月(卯)", "三月(辰)", "四月(巳)", "五月(午)", "六月(未)", "七月(申)", "八月(酉)", "九月(戌)", "十月(亥)", "十一月(子)", "十二月(丑)"]
				var selectm= $(select#monthid);
				for (var i =0;i<12;++i){
				    selectm.options.append(<option value={i+1}>{marr[i]}</option>);
					var m = select.text;
					selectm.value = t.month;
				}
        </script>
    </select>
    <!-- 下拉日期 -->
    <select id="dayid" name="day">
        <option value=" ">农历日</option>
        <script type="text/tiscript">
				var selectd= $(select#dayid);
				var darr = ["初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十",
				"十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十",
				"廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十"];
				for (var i =1;i<31;++i){
					selectd.options.append(<option value={i}>{darr[i-1]}</option>);
					var d = select.text;
                        if (t.day ==31){
                            selectd.value = t.day-1;
                        }else{
                            selectd.value = t.day;
                        }
				}
        </script>
    </select>
    <!-- 时辰 子时1 丑时2 寅时3... -->
    <select id="hourid" name="hour">
        <option value=" ">农历时辰</option>
        <script type="text/tiscript">
				var selecth =$(select#hourid);
				var harr=["子时(23~1)", "丑时(1~3)", "寅时(3~5)", "卯时(5~7)", "辰时(7~9)", "巳时(9~11)", "午时(11~13))", "未时(13~15)", "申时(15~17)", "酉时(17~19)", "戌时(19~21)","亥时(21~23)"]
				for (var i =0;i<12;++i){
				  // stdout.println(harr[i]);
				  selecth.options.append(<option value={i+1}>{harr[i]}</option>);
				  var h = select.text;
					if (t.hour/2==0){
                        selecth.value = 1;
					}else{
                        selecth.value = t.hour/2;
					}
				}
        </script>
    </select>
    <!--生肖选择-->
    <select id="sxid" name="shengxiao">
        <option value=" ">生肖</option>
        <script type="text/tiscript">
				var selectsx =$(select#sxid);
				var sxarr=["鼠","牛","虎","兔","龙","蛇","马","羊","猴","鸡","狗","猪"]
				for (var i =0;i<12;++i){
				  //stdout.println(sxarr[i]);
				  selectsx.options.append(<option value={sxarr[i]}>{sxarr[i]}</option>);
				  var sx = select.text;
				  selectsx.value = "虎";
				}
        </script>
    </select>
    <!--闰月选择-->
    <select id="leapmbid" name="lmb" title="否:农历不闰月 是:农历闰月">
        <option value=" ">闰月</option>
        <script type="text/tiscript">
			var selectmb=$(select#leapmbid);
			var larr = ["否","是"]
			var mbarr =["false","true"];
			for (var i =0;i<2;++i){
				selectmb.options.append(<option value={mbarr[i]}>{larr[i]}</option>);
				var mb = select.text;
				selectmb.value = 0;
			}
		 </script>
    </select>
</form>
<!--hr-->
<!--按钮:基本纪年 择吉 择日 奇门 禽星  月历-->
<div id="btn">
    <button id="btn1">纪年信息</button>
    <button id="btn2">小六壬择吉</button>
    <button id="btn3" title="岁吉凶 日吉凶依据干支计算">协纪辩方书</button>
    <button id="btn7" title="九星 八门 八神 天盘奇仪 地盘奇仪 暗干支">奇门</button>
    <button id="btn8">禽星</button>
    <button id="btn4">月历</button>
    <button id="btn5">24节气</button>
    <button id="btn6">地母经</button>
</div>
<!--显示基本纪年 择吉 择日 奇门 禽星  月历-->
<div id="show">
    <p id="p1"></p>
    <p id="p2"></p>
    <p id="p4"></p>
    <p id="p5"></p>
    <p id="p6"></p>
</div>
<!--奇门九宫格-->
<p id="qimen"></p>
<!--奇门九宫格结束-->

<!--月历表-->
<table id="tables"  rules="all" width="100%">
    <tr class="menu">
        <td #1 style="font-size:20px"></td>
        <td #2 style="font-size:20px"></td>
        <td #3 style="font-size:20px"></td>
        <td #4 style="font-size:20px"></td>
        <td #5 style="font-size:20px"></td>
        <td #6 style="font-size:20px"></td>
    </tr>
    <td #line1></td>
    <tr class="menu">
        <td #7 style="font-size:20px"></td>
        <td #8 style="font-size:20px"></td>
        <td #9 style="font-size:20px"></td>
        <td #10 style="font-size:20px"></td>
        <td #11 style="font-size:20px"></td>
        <td #12 style="font-size:20px"></td>
    </tr>
     <td #line2></td>
    <tr class="menu">
        <td #13 style="font-size:20px"></td>
        <td #14 style="font-size:20px"></td>
        <td #15 style="font-size:20px"></td>
        <td #16 style="font-size:20px"></td>
        <td #17 style="font-size:20px"></td>
        <td #18 style="font-size:20px"></td>
    </tr>
    <td #line3></td>
    <tr class="menu">
        <td #19 style="font-size:20px"></td>
        <td #20 style="font-size:20px"></td>
        <td #21 style="font-size:20px"></td>
        <td #22 style="font-size:20px"></td>
        <td #23 style="font-size:20px"></td>
        <td #24 style="font-size:20px"></td>
    </tr>
    <td #line4></td>
    <tr class="menu">
        <td #25 style="font-size:20px"></td>
        <td #26 style="font-size:20px"></td>
        <td #27 style="font-size:20px"></td>
        <td #28 style="font-size:20px"></td>
        <td #29 style="font-size:20px"></td>
        <td #30 style="font-size:20px"></td>
    </tr>
    <td #line5></td>
</table>

<!-- info 吉凶 -->
<div id="formjx">
    <form action="">
        <select id="yji" name="suiji">
            <option value=" ">岁吉</option>
            <option value="岁德">岁德</option>
            <option value="岁德合">岁德合</option>
            <option value="岁枝德">岁枝德</option>
            <option value="天乙贵人">天乙贵人</option>
            <option value="岁禄">岁禄</option>
            <option value="灾退">灾退</option>
        </select>
        <select id="yxiong" name="suixiong">
            <option value=" ">岁凶</option>
            <script type="text/tiscript">
				var sxArr = ["力士","丧门","太阴弔克","管符","白虎","黄幡","豹尾","病符","死符小耗","飞廉","金神","太岁","五鬼","破败五鬼"];
                var selectx = $(select#yxiong);
                var sx = 0;
                for (var i=0;i<sxArr.length;++i){
                   selectx.options.append(<option value={sxArr[i]}>{sxArr[i]}</option>);
				   var sx = select.text;
				}
            </script>
        </select>
        <select id="ysha" name="suisha">
            <option value=" ">岁煞</option>
            <script type="text/tiscript">
					var ssarr = ["三煞","岁破大耗","岁刑","太岁大煞","太岁劫煞","灾煞","岁煞"];
                    var selectss = $(select#ysha);
                    for (var i=0;i<7;++i){
                        selectss.options.append(<option value={ssarr[i]}>{ssarr[i]}</option>);
                        var ss = select.text;
                    }
            </script>
        </select>
        <br>
        <!--月吉凶-->
        <select id="mji" name="yueji">
            <option value=" ">月吉</option>
            <script type="text/tiscript">
            var mjarr = ["天道","天赦","天德", "天德合", "天恩", "天愿", "月德", "月德合", "月恩",
            "月空", "母仓", "时德", "阴德", "阳德", "时阳生气", "益后", "续世", "四相",
            "天仓", "要安", "敬安", "三合", "五合", "六合", "天医天喜", "五富", "玉宇",
            "福德天巫", "六仪", "金堂", "天马", "时阴", "驿马", "普护", "福生", "解神",
            "除神", "聖心", "吠鸣日", "吠鸣对日", "临日", "王官守相民日", "建禄",
            "宝", "养日", "专日", "制日"];
                var selectmj =$(select#mji);
                for (var i=0;i<mjarr.length;++i){
                    selectmj.options.append(<option value={mjarr[i]}>{mjarr[i]}</option>);
                    var monthj = select.text;
                }
            </script>
        </select>
        <select id="mxiong" name="yuexiong">
            <option value=" ">月凶</option>
            <script type="text/tiscript">
                var mxarr = ["天罡河魁死神", "河魁天罡", "月建", "月刑", "月害", "月破",
                "月厌地火", "月煞月虚", "天火灾煞", "小时", "血支", "天贼", "往亡",
                "咸池大时大败", "厌对招摇", "九空", "九坎", "游祸", "劫煞", "重日",
                "管符死气", "大耗", "小耗", "复日", "天吏致死", "大煞", "八龙",
                "四穷", "四耗", "四废", "四忌", "五虚", "五离", "八风", "土符",
                "归忌", "血忌", "五墓", "八专", "触水龙", "兵禁", "月忌日",
                "无禄日", "上朔", "伐日"];
                var selectmx =$(select#mxiong);
                for (var i=0;i<mxarr.length;++i){
                    selectmx.options.append(<option value={mxarr[i]}>{mxarr[i]}</option>);
                    var monthx = select.text;
                }
            </script>
        </select>
        <!--建除 月日论-->
        <select id="jianChu" name="jianchu">
            <option value="">日建除</option>
            <script type="text/tiscript">
                var jcarr = ["建","除","满","平","定","执","破","危","成","收","开","闭"];
                var selectjc = $(select#jianChu);
                for (var i=0;i<jcarr.length;++i){
                    selectjc.options.append(<option value={jcarr[i]}>{jcarr[i]}</option>);
                    var jcname = select.text;
                }
            </script>
        </select>
    </form>

    <button id="btninfo">查看吉凶信息</button>
</div>
<p class="ridgown" id="yjx"></p>
<!--info end-->

<!--干支八宫查询-->
<div id="form1">
    <form>
        <select id="zhiid" name="zhi">
            <option value=" ">地支信息</option>
            <script type="text/tiscript">
				var selectz =$(select#zhiid);
				var zhiarr=["子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"]
				for (var i =0;i<12;++i){
				  selectz.options.append(<option value={zhiarr[i]}>{zhiarr[i]}</option>);
				  var dz = select.text;
				}
            </script>
        </select>
        <select id="ganid" name="gan">
            <option value=" ">天干信息</option>
            <script type="text/tiscript">
				var selectg =$(select#ganid);
				var ganarr=["甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"]
				for (var i =0;i<10;++i){
				  selectg.options.append(<option value={ganarr[i]}>{ganarr[i]}</option>);
				  var tg = select.text;
				}
            </script>
        </select>
        <select id="guaid" name="gua">
            <option value=" ">八卦信息</option>
            <script type="text/tiscript">
				var selectgua =$(select#guaid);
				var garr=["坎", "艮", "震", "巽", "离", "坤", "兑", "乾"]
				for (var i =0;i<8;++i){
				  selectgua.options.append(<option value={garr[i]}>{garr[i]}</option>);
				  var gua = select.text;
				}
            </script>
        </select>
    </form>
</div>
<button id="btngzs">查看</button>
<!--干支八宫查询 end-->

<script type="text/tiscript">

    //清除月历表内容
    var clean = function(){
        $(#1).html ="";
        $(#2).html ="";
        $(#3).html ="";
        $(#4).html ="";
        $(#5).html ="";
        $(#6).html ="";
        $(#7).html ="";
        $(#8).html ="";
        $(#9).html ="";
        $(#10).html="";
        $(#11).html="";
        $(#12).html="";
        $(#13).html ="";
        $(#14).html ="";
        $(#15).html ="";
        $(#16).html ="";
        $(#17).html ="";
        $(#18).html ="";
        $(#19).html ="";
        $(#20).html ="";
        $(#21).html ="";
        $(#22).html ="";
        $(#23).html ="";
        $(#24).html ="";
        $(#25).html ="";
        $(#26).html ="";
        $(#27).html ="";
        $(#28).html ="";
        $(#29).html ="";
        $(#30).html ="";
    };
    //下拉列表数据
    function ymdhsb(){
          //年
          var opt = $(select[name='year']).$$(option);
          var ly= 0;
          for(var child in opt) {
              //判断元素是否选中
              if(child.getState(Element.STATE_CHECKED)) {
                  ly = child.value;
                //  stdout.println("选中y:",ly);
              }
          }
          //月
          var optm = $(select[name='month']).$$(option);
          var lm= 0;
          for(var child in optm) {
              //判断元素是否选中
              if(child.getState(Element.STATE_CHECKED)) {
                  lm = child.value;
                //  stdout.println("选中m:",lm);
              }
          }
          //日
          var optd = $(select[name='day']).$$(option);
          var ld= 0;
          for(var child in optd) {
              //判断元素是否选中
              if(child.getState(Element.STATE_CHECKED)) {
                  ld = child.value;
                //  stdout.println("选中d:",ld);
              }
          }
           //时辰
          var opth = $(select[name='hour']).$$(option);
          var lh= 0;
          for(var child in opth) {
              //判断元素是否选中
              if(child.getState(Element.STATE_CHECKED)) {
                  lh = child.value;
                //  stdout.println("选中h:",lh);
              }
          }
          //生肖
          var optsx = $(select[name='shengxiao']).$$(option);
          var sx =0;
          for(var child in optsx){
          //判断元素是否选中
              if(child.getState(Element.STATE_CHECKED)){
                  sx = child.value;
                 // stdout.println("选中sx:",sx);
              }
          }
          //闰月
          var optmb=$(select[name='lmb']).$$(option);
          var lb=0;
          for(var child in optmb){
          //判断元素是否选中
              if(child.getState(Element.STATE_CHECKED)){
                  lb = child.value;
                  //stdout.println("选中lb:",lb);
              }
          }
         var formData = {
         "year":ly,
         "month":lm,
         "day":ld,
         "hour":lh,
         "zodiac":sx,
         "leapmb":lb
         };
         //stdout.println("func ymdh()-->",JSON.stringify(formData));
         return formData;
    }

    //基本纪年信息
	$(#btn1).on("click",function(){
		var infoymdh = ymdhsb();
		//stdout.println(infoymdh.year,infoymdh.month, infoymdh.day, infoymdh.hour,infoymdh.zodiac, infoymdh.leapmb)

         //定一个window方法
          var info = view.ymdinfo(infoymdh.year,infoymdh.month, infoymdh.day, infoymdh.hour,infoymdh.zodiac, infoymdh.leapmb);
         //stdout.println("ymdinfo:",info);//返回go处理之后的结果
          var jsonData = parseData(info);//解析json
          if(jsonData){
           // stdout.printf("JSON data is : %v\n", jsonData);
            $(#p1).value = jsonData.sjn+"\n"+jsonData.ljn+"\n"+jsonData.gjn+"\n"+jsonData.ny+"\n"+jsonData.lmb;
          }else{
            stdout.println("No data");
          }
    });

    //小六壬
    $(#btn2).on("click",function(){
        var xlrymdh = ymdhsb();
        var xlrinfo = view.xlrzjinfo(xlrymdh.year,xlrymdh.month, xlrymdh.day, xlrymdh.hour,xlrymdh.zodiac, xlrymdh.leapmb);
        var jsonData = parseData(xlrinfo);
        if(jsonData){
            $(#qimen).value = "";
            clean();
            //stdout.printf("JSON data is : %v\n", jsonData);
            $(#p2).value = jsonData.xstar_name+"\n"+jsonData.xzeji+"\n"+jsonData.x_ji_gan_arr+"\n"+jsonData.x_qi_sha_arr;
        }else{
            stdout.println("No data");
        }
    });
    //协纪辩方书-
    $(#btn3).on("click",function(){
        var xjymdh = ymdhsb();
        var xjinfo = view.xjbfsinfo(xjymdh.year,xjymdh.month,xjymdh.day,xjymdh.hour,xjymdh.zodiac,xjymdh.leapmb);
        var xjbfs = parseData(xjinfo);
        if (xjbfs){
            $(#qimen).value = "";
            clean();
            $(#p2).value = xjbfs.nb+"\n"+xjbfs.yb+"\n"+xjbfs.rb+"\n"+xjbfs.bw
        }else{
            stdout.println("err no data");
        }
    });

    //今日农历日期-->(月历表)
	$(#btn4).on("click",function(){
		var ylb = ymdhsb();
        var ylbs = view.todaytinfo(ylb.year,ylb.month,ylb.day,ylb.hour,ylb.zodiac,ylb.leapmb);
        $(#qimen).value = "";
        var obj = parseData(ylbs);
        var s = obj.sdays;//阳历数组
        var l = obj.ldays;
        var g = obj.gzhis;

        if (obj.sdays.length==30){
            $(#p2).value = "";//清空避免重复
            $(#qimen).value ="";
            $(#1).html = s[0]+"<br>"+l[0]+"<br>"+g[0]+"<br>"
            $(#2).html = s[1]+"<br>"+l[1]+"<br>"+g[1]+"<br>"
            $(#3).html = s[2]+"<br>"+l[2]+"<br>"+g[2]+"<br>"
            $(#4).html = s[3]+"<br>"+l[3]+"<br>"+g[3]+"<br>"
            $(#5).html = s[4]+"<br>"+l[4]+"<br>"+g[4]+"<br>"
            $(#6).html = s[5]+"<br>"+l[5]+"<br>"+g[5]+"<br>"
			$(#line1).html= "<br>"
            $(#7).html = s[6]+"<br>"+l[6]+"<br>"+g[6]+"<br>"
            $(#8).html = s[7]+"<br>"+l[7]+"<br>"+g[7]+"<br>"
            $(#9).html = s[8]+"<br>"+l[8]+"<br>"+g[8]+"<br>"
            $(#10).html =s[9]+"<br>"+l[9]+"<br>"+g[9]+"<br>"
            $(#11).html=s[10]+"<br>"+l[10]+"<br>"+g[10]+"<br>"
            $(#12).html=s[11]+"<br>"+l[11]+"<br>"+g[11]+"<br>"
            $(#line2).html= "<br>"
            $(#13).html = s[12]+"<br>"+l[12]+"<br>"+g[12]+"<br>"
            $(#14).html = s[13]+"<br>"+l[13]+"<br>"+g[13]+"<br>"
            $(#15).html = s[14]+"<br>"+l[14]+"<br>"+g[14]+"<br>"
            $(#16).html = s[15]+"<br>"+l[15]+"<br>"+g[15]+"<br>"
            $(#17).html = s[16]+"<br>"+l[16]+"<br>"+g[16]+"<br>"
            $(#18).html = s[17]+"<br>"+l[17]+"<br>"+g[17]+"<br>"
            $(#line3).html= "<br>"
            $(#19).html = s[18]+"<br>"+l[18]+"<br>"+g[18]+"<br>"
            $(#20).html = s[19]+"<br>"+l[19]+"<br>"+g[19]+"<br>"
            $(#21).html = s[20]+"<br>"+l[20]+"<br>"+g[20]+"<br>"
            $(#22).html = s[21]+"<br>"+l[21]+"<br>"+g[21]+"<br>"
            $(#23).html = s[22]+"<br>"+l[22]+"<br>"+g[22]+"<br>"
            $(#24).html = s[23]+"<br>"+l[23]+"<br>"+g[23]+"<br>"
            $(#line4).html= "<br>"
            $(#25).html = s[24]+"<br>"+l[24]+"<br>"+g[24]+"<br>"
            $(#26).html = s[25]+"<br>"+l[25]+"<br>"+g[25]+"<br>"
            $(#27).html = s[26]+"<br>"+l[26]+"<br>"+g[26]+"<br>"
            $(#28).html = s[27]+"<br>"+l[27]+"<br>"+g[27]+"<br>"
            $(#29).html = s[28]+"<br>"+l[28]+"<br>"+g[28]+"<br>"
            $(#30).html = s[29]+"<br>"+l[29]+"<br>"+g[29]+"<br>"
            $(#line5).html= "<br>"
          }else if (obj.sdays.length==29){
            $(#p2).value = "";//清空避免重复
            $(#qimen).value ="";

            $(#1).html = s[0]+"<br>"+l[0]+"<br>"+g[0]+"<br>"
            $(#2).html = s[1]+"<br>"+l[1]+"<br>"+g[1]+"<br>"
            $(#3).html = s[2]+"<br>"+l[2]+"<br>"+g[2]+"<br>"
            $(#4).html = s[3]+"<br>"+l[3]+"<br>"+g[3]+"<br>"
            $(#5).html = s[4]+"<br>"+l[4]+"<br>"+g[4]+"<br>"
            $(#6).html = s[5]+"<br>"+l[5]+"<br>"+g[5]+"<br>"
            $(#line1).html= "<br>"
            $(#7).html = s[6]+"<br>"+l[6]+"<br>"+g[6]+"<br>"
            $(#8).html = s[7]+"<br>"+l[7]+"<br>"+g[7]+"<br>"
            $(#9).html = s[8]+"<br>"+l[8]+"<br>"+g[8]+"<br>"
            $(#10).html =s[9]+"<br>"+l[9]+"<br>"+g[9]+"<br>"
            $(#11).html=s[10]+"<br>"+l[10]+"<br>"+g[10]+"<br>"
            $(#12).html=s[11]+"<br>"+l[11]+"<br>"+g[11]+"<br>"
            $(#line2).html= "<br>"
            $(#13).html = s[12]+"<br>"+l[12]+"<br>"+g[12]+"<br>"
            $(#14).html = s[13]+"<br>"+l[13]+"<br>"+g[13]+"<br>"
            $(#15).html = s[14]+"<br>"+l[14]+"<br>"+g[14]+"<br>"
            $(#16).html = s[15]+"<br>"+l[15]+"<br>"+g[15]+"<br>"
            $(#17).html = s[16]+"<br>"+l[16]+"<br>"+g[15]+"<br>"
            $(#18).html = s[17]+"<br>"+l[17]+"<br>"+g[17]+"<br>"
            $(#line3).html= "<br>"
            $(#19).html = s[18]+"<br>"+l[18]+"<br>"+g[18]+"<br>"
            $(#20).html = s[19]+"<br>"+l[19]+"<br>"+g[19]+"<br>"
            $(#21).html = s[20]+"<br>"+l[20]+"<br>"+g[20]+"<br>"
            $(#22).html = s[21]+"<br>"+l[21]+"<br>"+g[21]+"<br>"
            $(#23).html = s[22]+"<br>"+l[22]+"<br>"+g[22]+"<br>"
            $(#24).html = s[23]+"<br>"+l[23]+"<br>"+g[23]+"<br>"
            $(#line4).html= "<br>"
            $(#25).html = s[24]+"<br>"+l[24]+"<br>"+g[24]+"<br>"
            $(#26).html = s[25]+"<br>"+l[25]+"<br>"+g[25]+"<br>"
            $(#27).html = s[26]+"<br>"+l[26]+"<br>"+g[26]+"<br>"
            $(#28).html = s[27]+"<br>"+l[27]+"<br>"+g[27]+"<br>"
            $(#29).html = s[28]+"<br>"+l[28]+"<br>"+g[28]+"<br>"
            $(#line5).html= "<br>"
           // $(#30).html = s[29]+"<br>"+l[29]+"<br>"+g[29]+"<br>"+"<br>"
        }
       // $(#p2).value= tinfo;
    });

    //24节气
	$(#btn5).on("click",function(){
         var opt = $(select[name='year']).$$(option);
         var ly= 0;
            for(var child in opt) {
              if(child.getState(Element.STATE_CHECKED)) {
                  ly = child.value;
            }
         }
        var jq24 = view.jieqiinfo(ly);
        $(#yjx).value = jq24;//显示到右下角布局内
    });

    //btn6 aboout-->地母经
     var root = view.window;
     $(#btn6).on("click",function(){
		var dm = ymdhsb();
        var dimus= view.aboutinfo(dm.year,dm.month, dm.day, dm.hour,dm.zodiac, dm.leapmb);
        $(#qimen).value ="";
        clean();
        $(#p2).value = dimus;
        //var ab = parseData(about);
       // view.msgbox("https://github.com/Aquarian-Age/ccal");//弹窗显示
    });

    //奇门
	$(#btn7).on("click",function(){
        var qmymdh = ymdhsb();
        var qms = view.qimeninfo(qmymdh.year,qmymdh.month, qmymdh.day, qmymdh.hour,qmymdh.zodiac, qmymdh.leapmb);
        var qmdJs = view.qmMethod(qmymdh.year,qmymdh.month, qmymdh.day, qmymdh.hour,qmymdh.zodiac, qmymdh.leapmb);
       //0:九星 1:八门 2:暗干支 3:天盘奇仪 4:八神 5:地盘奇仪
       var qmjs = parseData(qms);
       clean();
       $(#p2).value = "";
       $(#qimen).html=qmjs.jie_qi+" "+qmjs.yin_yang+qmjs.n+"局"+" "+qmjs.yuan+"<br>"+
       "值符:"+qmjs.zhifu+" "+"值使:"+qmjs.zhishi+"<br>"+
       "----------------------------"+"<br>"+
        qmjs.g_4[0]+"   "+"|"+qmjs.g_9[0]+"   ""|"+qmjs.g_2[0]+"   "+"<br>"+
        qmjs.g_4[1]+"   "+"|"+qmjs.g_9[1]+"   "+"|"+qmjs.g_2[1]+"   "+"<br>"+
        qmjs.g_4[4]+"   "+"|"+qmjs.g_9[4]+"   "+"|"+qmjs.g_2[4]+"   "+"<br>"+
        qmjs.g_4[3]+"     "+"|"+qmjs.g_9[3]+"     "+"|"+qmjs.g_2[3]+"     "+"<br>"+
        qmjs.g_4[5]+"     "+"|"+qmjs.g_9[5]+"     "+"|"+qmjs.g_2[5]+"     "+"<br>"+
        qmjs.g_4[2]+"   "+"|"+qmjs.g_9[2]+"   "+"|"+qmjs.g_2[2]+"<br>"+
        "----------------------------"+"<br>"+
        qmjs.g_3[0]+"   "+"|"+qmjs.g_5[0]+"   ""|"+qmjs.g_7[0]+"   "+"<br>"+
        qmjs.g_3[1]+"   "+"|"+qmjs.g_5[1]+"       "+"|"+qmjs.g_7[1]+"   "+"<br>"+
        qmjs.g_3[4]+"   "+"|"+qmjs.g_5[4]+"       "+"|"+qmjs.g_7[4]+"   "+"<br>"+
        qmjs.g_3[3]+"     "+"|"+qmjs.g_5[3]+"     "+"|"+qmjs.g_7[3]+"     "+"<br>"+
        qmjs.g_3[5]+"     "+"|"+qmjs.g_5[5]+"     "+"|"+qmjs.g_7[5]+"     "+"<br>"+
        qmjs.g_3[2]+"   "+"|"+qmjs.g_5[2]+"   "+"|"+qmjs.g_7[2]+"<br>"+
       "----------------------------"+"<br>"+
        qmjs.g_8[0]+"   "+"|"+qmjs.g_1[0]+"   ""|"+qmjs.g_6[0]+"   "+"<br>"+
        qmjs.g_8[1]+"   "+"|"+qmjs.g_1[1]+"   "+"|"+qmjs.g_6[1]+"   "+"<br>"+
        qmjs.g_8[4]+"   "+"|"+qmjs.g_1[4]+"   "+"|"+qmjs.g_6[4]+"   "+"<br>"+
        qmjs.g_8[3]+"     "+"|"+qmjs.g_1[3]+"     "+"|"+qmjs.g_6[3]+"     "+"<br>"+
        qmjs.g_8[5]+"     "+"|"+qmjs.g_1[5]+"     "+"|"+qmjs.g_6[5]+"     "+"<br>"+
        qmjs.g_8[2]+"   "+"|"+qmjs.g_1[2]+"   "+"|"+qmjs.g_6[2]+"<br>"+
        "----------------------------"+"<br>"
        //奇门方法 显示到右下角
         //stdout.println(qmd);
         var qmd = parseData(qmdJs);
         $(#yjx).html="[地四户] "+qmd.di_si_hu + "<br>"+
         "[地私门] "+qmd.di_si_men + "<br>"+
         "[太冲天马] "+qmd.tian_ma + "<br>"+
         "[天三门] "+qmd.tian_san_men +"<br>"+
         "[五符] "+qmd.wu_fus[0]+" "+qmd.wu_fus[1]+" "+qmd.wu_fus[2]+" "+qmd.wu_fus[3]+" "+qmd.wu_fus[4]+" "+qmd.wu_fus[5]+"<br>"+
          qmd.wu_fus[6]+" "+qmd.wu_fus[7]+" "+qmd.wu_fus[8]+" "+qmd.wu_fus[9]+" "+qmd.wu_fus[10]+" "+qmd.wu_fus[11]+"<br>"+
         "[时孤虚] "+qmd.gu_xu + "<br>"
    });

    //禽星
	$(#btn8).on("click",function(){
	    $(#qimen).value = "";
	    clean();
        var qinymdh = ymdhsb();
        var qxs = view.qinxinginfo(qinymdh.year,qinymdh.month, qinymdh.day, qinymdh.hour,qinymdh.zodiac, qinymdh.leapmb);
        $(#p2).value = qxs;
    });

    //info 吉凶下拉列表
	$(#btninfo).on("click",function(){
      var opt = $(select[name='suiji']).$$(option);
      var yj= 0;
         for(var child in opt) {
            if(child.getState(Element.STATE_CHECKED)) {
                yj = child.value;
            }
         }
         var opt1 = $(select[name='suixiong']).$$(option);
         var yj1= 0;
         for(var child in opt1) {
            if(child.getState(Element.STATE_CHECKED)) {
                yj1 = child.value;
            }
         }
         var opt2 = $(select[name='suisha']).$$(option);
         var yj2= 0;
         for(var child in opt2) {
            if(child.getState(Element.STATE_CHECKED)) {
                yj2 = child.value;
            }
         }
         //月吉凶
         var optmj = $(select[name='yueji']).$$(option);
         var mj = 0;
         for (var child in optmj){
            if (child.getState(Element.STATE_CHECKED)){
                mj = child.value;
            }
         }
         //月吉凶
         var optmx = $(select[name='yuexiong']).$$(option);
         var mx = 0;
         for (var child in optmx){
            if (child.getState(Element.STATE_CHECKED)){
                mx = child.value;
            }
         }
         //日建除
         var optjc = $(select[name='jianchu']).$$(option);
         var jc = 0;
         for (var child in optjc){
            if (child.getState(Element.STATE_CHECKED)){
                jc = child.value;
            }
         }

         //返回数据到后端
         var formData = {
         "suiji":yj,
         "suixiong":yj1,
         "suisha":yj2,
          "mjs":mj,
          "mxs":mx,
          "jc":jc
         };

         var jxs = view.jiXiongInfo(JSON.stringify(formData));
         $(#yjx).value = jxs;
    });

    //查询 天干 地支 八卦
    $(#btngzs).on("click",function(){
         var optg = $(select[name='gan']).$$(option);
            var gan = 0;
            for (var child in optg){
              if (child.getState(Element.STATE_CHECKED)){
                gan = child.value;
            }
         }

        var optz = $(select[name='zhi']).$$(option);
            var zhi = 0;
            for (var child in optz){
                if (child.getState(Element.STATE_CHECKED)){
                zhi = child.value;
            }
        }

        var optgua = $(select[name='gua']).$$(option);
            var gua = 0;
            for (var child in optgua){
                if (child.getState(Element.STATE_CHECKED)){
                gua = child.value;
            }
        }

        var gzgData ={
            "gan":gan,
            "zhi":zhi,
            "gua":gua
        };
        var gzg = view.ganZhiGua(JSON.stringify(gzgData));
        view.msgbox(gzg);//弹窗显示
    });
</script>

</body>

</html>
`
