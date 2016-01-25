//2016-01-30 09:00:00	2016-01-30 10:00:00
//2016-01-30 11:00:00 	2016-01-30 12:20:00
//2016-01-30 10:10:30 	2016-01-30 12:45:50

package main

import (
	"fmt"
	"time"
)

func main() {
	//获取时间戳

	timestamp := time.Now()
	fmt.Println("time is : ", timestamp)
	fmt.Println("now time is :", timestamp.Format("2006-01-02 15:04:05"))
	fmt.Println("now unix time is : ", timestamp.Unix())
	zone, _ := timestamp.Zone()
	fmt.Println("now zone : ", zone)

	// 格式化为字符串,tm为Time类型

	// 从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	// 2006-01-02 15:04:05，这个日期是固定的，其他日期会出错

	loc, _ := time.LoadLocation("Europe/Berlin")
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println("loc : ", loc)
	fmt.Println("------------")
	tm1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-01-30 09:00:00", loc)
	fmt.Println(tm1)
	fmt.Println(tm1)
	fmt.Println(tm1.Unix())
	fmt.Println("------------")
	tm2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-01-30 10:00:00", loc)
	fmt.Println(tm2)
	fmt.Println(tm2.Unix())
	fmt.Println("------------")
	tm3, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-01-30 11:00:00", loc)
	fmt.Println(tm3)
	fmt.Println(tm3.Unix())
	fmt.Println("------------")
	tm4, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-01-30 12:20:00", loc)
	fmt.Println(tm4)
	fmt.Println(tm4.Unix())
	fmt.Println("------------")
	tm5, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-01-30 11:00:00", loc)
	fmt.Println(tm5)
	fmt.Println(tm5.Unix())
	fmt.Println("------------")
	tm6, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-01-30 12:45:50", loc)
	fmt.Println(tm6)
	fmt.Println(tm6.Unix())
	fmt.Println("------------")
}