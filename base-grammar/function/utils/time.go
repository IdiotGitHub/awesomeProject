package utils

import (
	"fmt"
	"time"
)

/**
时间和日期相关的函数
	time.Time类型，用来表示时间
时间常量
	const(
		Nanosecond Duration = 1 //纳秒
		Microsecond			= 1000 * Nanosecond
		Millisecond			= 1000 * Microsecond
		Second				= 1000 * Millisecond
		Minute				= 60 * Second
		Hour				= 60 * Minute
	)
	使用时间常量时不能使用除法
*/
func Demo1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("年=%v\n", now.Year())
	//可以使用int将月份进行强转
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())
}

//格式化日期和时间1
func Demo2() {
	fmt.Printf("当前时间日期为%d-%d-%d %d:%d:%d\n", time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}

//格式化日期第二种方式
//使用time.Now().Format("2006/01/02 15:04:05")
func Demo3() {
	now := time.Now()
	fmt.Printf("当前时间为：%v", now.Format("2006-01-02 15:04:05"))
}
