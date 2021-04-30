package jdcal

import (
	"4d63.com/tz"
	jl "github.com/soniakeys/meeus/v3/julian"
	"log"
	"os"
	"time"
)

var formats = "2006-01-02 15:05:04"

func JDToCal(jd float64) time.Time {
	utcT := jl.JDToTime(jd) //utc这里决定了时间转换的准确度
	//fmt.Println("utcT1: ", utcT)
	////
	location, err := tz.LoadLocation("Asia/Shanghai") //兼容没有go time包的系统 1.15版本的go 也可以直接使用-tags timetzdata
	if err == nil {
		utcT = utcT.In(location)
	}
	//fmt.Println("utcT2: ", utcT)
	////
	utcs := utcT.Format(formats)
	t, err := time.Parse(formats, utcs)
	if err != nil {
		log.Println("ParseTime: ", err)
		os.Exit(0)
	}
	return t
}

/*
jd := 2743362.291817
utcT1:  2798-12-20 19:00:12.988811731 +0000 UTC
utcT2:  2798-12-21 03:00:12.988811731 +0800 CST
*/
