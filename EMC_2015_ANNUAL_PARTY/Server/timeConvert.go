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

	timestamp := time.Now().Unix()

	fmt.Println(timestamp)

 

	//格式化为字符串,tm为Time类型

	tm := time.Unix(timestamp, 0)

	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))

	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串

	tm1, _ := time.Parse("2006-01-02 03:04:05 PM", "2016-01-30 09:00:00 AM")
	fmt.Println(tm1.Unix())
	tm2, _ := time.Parse("2006-01-02 03:04:05 PM", "2016-01-30 10:00:00 AM")
	fmt.Println(tm2.Unix())
	tm3, _ := time.Parse("2006-01-02 03:04:05 PM", "2016-01-30 11:00:00 AM")
	fmt.Println(tm3.Unix())
	tm4, _ := time.Parse("2006-01-02 03:04:05 PM", "2016-01-30 12:20:00 AM")
	fmt.Println(tm4.Unix())
	tm5, _ := time.Parse("2006-01-02 03:04:05 PM", "2016-01-30 11:00:00 AM")
	fmt.Println(tm5.Unix())
	tm6, _ := time.Parse("2006-01-02 03:04:05 PM", "2016-01-30 12:45:50 AM")
	fmt.Println(tm6.Unix())

}