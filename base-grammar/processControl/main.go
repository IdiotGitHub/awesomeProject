package main

import (
	"fmt"
	"math"
)

/*
流程控制
	1.顺序控制
	2.分支控制
		1.单分支
			if 条件表达式{
				执行代码块
				}
			大括号必须有
			可以在条件表达式中进行变量声明并赋值
			if中的条件表达式，不能是别的语句（赋值表达式等）
		2.双分支
		3.多分支
		4.switch分支结构
			1.switch语句用于基于不同条件执行不同动作，每一个case分支都是唯一的，从上到下注意测试，直到匹配为止。
			2.匹配项后面也不需要再加break
			语法：
				switch 表达式 {
					case 表达式1、表达式2、……:
						语句块
					case 表达式3、表达式4、……:
						语句块
					default:
						语句块
				}
			跟其它语言的区别，case中不需要break，执行完成case自动退出整个switch；而且只有在所有的case不执行的时候采取执行default语句
			使用细节
				1.case后面可以是一个表达式（可以是常量（但是感觉并不实用）、变量、有返回值的函数）
				2.case后面的表达式数据类型必须和switch后面的表达式数据类型相同
				3.case后面的表达式如果是常量值，则不允许重复出现
				4.switch后面可以不带表达式类似if-else分支来使用
					var age int64 = 25
					switch{
					case age == 20:
						fmt.Println("age == 10")
					case age == 25:
						fmt.Println("age == 25")
					default:
						fmt.Println("没有相应的匹配项")
					}
				5.switch 后面也可以直接声明/定义一个变量分号结束
				6.switch穿透fallthrough，如果想要某一个case执行完成之后继续执行下一个case，那么可以在该case代码块最后使用fallthrough，这个功能有点鸡肋，可以使用case后面多表达式的形式代替
				7.Type Switch：switch语句还可以被用于type-switch来判断某个interface变量中实际指向的变量类型（这种情况可以在某些场景下进行使用，比如在web中在进行数据库查询的时候，有些时候是查询不到数据的，那么就可以使用这种方式进行判断，如果返回了正常类型，那么表示已经查询到相应的数据，如果是nil就是没有查询到，如果是别的，就可以使用default进行捕获了；有待验证）
				8.switch和if的比较：如果判断的具体数值不多，而且符合整数、浮点数、字符、字符串这几种类型。建议使用switch语句；其他情况，对区间判断和结果为bool类型的判断使用if
	3.循环控制
*/

func main() {
	//singleBench()
	//exercise1()
	//exercise2()
	//exercise3()
	//exercise4()
	switchControl()
}

func singleBench() {
	var age int32
	fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你已经成年了")
	} else {
		fmt.Println("你还未成年")
	}
	if age2 := 30; age2 > 33 {
		fmt.Println(age2)
	} else {
		fmt.Println(age2)
	}
}
func exercise1() {
	var num1 int32 = 23
	var num2 int32 = 32
	if num1+num2 >= 50 {
		fmt.Println("hello world")
	}
}
func exercise2() {
	var num1 float64 = 20.2
	var num2 float64 = 10.3
	if num1 > 10 && num2 < 20 {
		fmt.Println(num1 + num2)
	}
}

func exercise3() {
	var year int64 = 2020
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Printf("%d年是闰年", year)
	} else {
		fmt.Printf("%d年是平年", year)

	}
}

func exercise4() {
	var score float64
	fmt.Println("请输入岳云鹏的考试成绩：")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("恭喜岳云鹏获得一台BMW")
	} else if score > 80 && score < 100 {
		fmt.Println("恭喜岳云鹏获得一台iphone7p")
	} else if score >= 60 && score <= 80 {
		fmt.Println("恭喜岳云鹏获得一台ipad")
	} else {
		fmt.Println("恭喜岳云鹏获得一口西北风")
	}
}

func exercise5() {
	var a float64
	var b float64
	var c float64
	fmt.Println("请分别输入参数a、b、c的值，回车结束")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	if b*b-4*a*c > 0 {
		fmt.Println("方程有两个解：")
		fmt.Printf("解1 = %f, 解2 = %f\n", (-b+math.Sqrt(b*b-4*a*c))/2*a, (-b-math.Sqrt(b*b-4*a*c))/2*a)
	} else if b*b-4*a*c == 0 {
		fmt.Println("方程有一个解：")
		fmt.Printf("解 = %f\n", (-b+math.Sqrt(b*b-4*a*c))/2*a)
	} else {
		fmt.Println("方程无解")
	}
}

func switchControl() {
	//这个地方想用byte也可以，但是输入的时候就别用Scanln了，这个输入是输入字符串格式，
	//就需要使用Scanf来指定输入的格式即可
	var week string
	fmt.Println("请输入a、b、c、d、e、f、g")
	fmt.Scanln(&week)
	fmt.Println("你输入的是", week)
	switch week {
	case "a":
		fmt.Println("周一")
	case "b":
		fmt.Println("周二")
	case "c":
		fmt.Println("周三")
	case "d":
		fmt.Println("周四")
	case "e":
		fmt.Println("周五")
	case "f":
		fmt.Println("周六")
	case "g":
		fmt.Println("周日")
	default:
		fmt.Println("你的输入有误,请重新输入")
		switchControl()
	}

}
