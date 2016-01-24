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

	fmt.Println("now time is : ", timestamp)

	// 格式化为字符串,tm为Time类型

	// 从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	// 2006-01-02 15:04:05，这个日期是固定的，其他日期会出错

	tm1, _ := time.Parse("2006-01-02 15:04:05", "2016-01-30 09:00:00")
	fmt.Println(tm1.Unix())
	tm2, _ := time.Parse("2006-01-02 15:04:05", "2016-01-30 10:00:00")
	fmt.Println(tm2.Unix())
	tm3, _ := time.Parse("2006-01-02 15:04:05", "2016-01-30 11:00:00")
	fmt.Println(tm3.Unix())
	tm4, _ := time.Parse("2006-01-02 15:04:05", "2016-01-30 12:20:00")
	fmt.Println(tm4.Unix())
	tm5, _ := time.Parse("2006-01-02 15:04:05", "2016-01-30 11:00:00")
	fmt.Println(tm5.Unix())
	tm6, _ := time.Parse("2006-01-02 15:04:05", "2016-01-30 12:45:50")
	fmt.Println(tm6.Unix())

}