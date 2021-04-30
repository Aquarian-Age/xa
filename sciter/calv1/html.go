package main

const html = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>阳历转干支历</title>
</head>
<body>
<form>
    <select id="yid" name="year">
        <script type="text/tiscript">
            var t = new Date();
            var selecty = $(select#yid);
            for(var i=1600;i<3500;++i){
                selecty.options.append(<option value={i}>{i}年</option>);
                var y = selecty.text;
                selecty.value = t.year;
            }
        </script>
    </select>
    <select id="mid" name="month">
        <script type="text/tiscript">
            var selectm = $(select#mid);
            for(var i=1;i<13;++i){
                selectm.options.append(<option value={i}>{i}月</option>);
                var m = selectm.text;
                selectm.value =t.month;
            }
        </script>
    </select>
    <select id="did" name="day">
        <script type="text/tiscript">
            var selectd = $(select#did);
            for(var i=0;i<=31;++i){
                selectd.options.append(<option value={i}>{i}日</option>);
                var d = selectd.text;
                selectd.value = t.day;
            }
        </script>
    </select>
    <select id="hid" name="hour">
        <script type="text/tiscript">
            var selecth = $(select#hid);
            for(var i=0;i<=24;++i){
                selecth.options.append(<option value={i}>{i}时</option>);
                var h = selecth.text;
                selecth.value = t.hour;
            }
        </script>
    </select>
    <button id="btn1">查看</button>
</form>
<p id="p1" style="font-size: 16px"></p>

<script type="text/tiscript">
    function get(){
        var opt = $(select[name = 'year']).$$(option);
        var ly = 0;
        for (var child in opt) {
            if (child.getState(Element.STATE_CHECKED)) {
                ly = child.value;
            }
        }
        var optm = $(select[name = 'month']).$$(option);
        var lm = 0;
        for (var child in optm) {
            if (child.getState(Element.STATE_CHECKED)) {
                lm = child.value;
            }
        }
        var optd = $(select[name = 'day']).$$(option);
        var ld = 0;
        for (var child in optd) {
            if (child.getState(Element.STATE_CHECKED)) {
                ld = child.value;
            }
        }
        var opth = $(select[name = 'hour']).$$(option);
        var lh = 0;
        for (var child in opth) {
            if (child.getState(Element.STATE_CHECKED)) {
                lh = child.value;
            }
        }
        var formData = {
            "year": ly,
            "month": lm,
            "day": ld,
            "hour": lh
        };
        return formData;
    }

    $(#btn1).on("click", function(){
        var data = get();
        var gz = view.ymdToGZ(data.year,data.month,data.day,data.hour);
        var ganZhi= parseData(gz);
        $(#p1).html = ganZhi.year_gz+"年-"+ganZhi.month_gz+"月-"+ganZhi.day_gz+"日-"+ganZhi.hour_gz+"时"+"<br>";
    });
</script>
</body>
</html>`
