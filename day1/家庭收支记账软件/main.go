package main

import (
	"fmt"
	"os"
	"strings"
)

type Account struct {
	inOrOut string
	//balance     float64
	money       float64
	description string
}

func menu() {
	fmt.Println("---------------家庭收支记账软件---------------")
	fmt.Println("                 1.收支明细                  ")
	fmt.Println("                 2.登记收入                  ")
	fmt.Println("                 3.登记支出                  ")
	fmt.Println("                 4.退出                  ")
	fmt.Print("请选择(1-4)")

}

func function(option string, account *[]Account, balance *float64) {
	switch option {
	case "1":
		fmt.Println("--------------当前收支明细记录--------------")
		if len(*account) == 0 {
			fmt.Println("当前没有收支记录，请先收入一笔吧：")
		} else {
			fmt.Println("收支	账户金额	收支金额	说  明")
			for i := 0; i < len(*account); i++ {
				fmt.Printf("%s\t%.2f\t%.2f\t%s\t\n", (*account)[i].inOrOut,
					*balance, (*account)[i].money, (*account)[i].description)
			}
		}
	case "2":
		var acc Account
		fmt.Println("本次收入金额：")
		fmt.Scanln(&acc.money)
		fmt.Println("本次收入说明：")
		fmt.Scanln(&acc.description)
		acc.inOrOut = "收入"
		*balance += acc.money
		*account = append(*account, acc)
	case "3":
		var acc Account
		fmt.Println("本次支出金额：")
		fmt.Scanln(&acc.money)
		fmt.Println("本次支出说明：")
		fmt.Scanln(&acc.description)
		acc.inOrOut = "支出"
		*balance -= acc.money
		*account = append(*account, acc)
	case "4":
		var exit string
		fmt.Println("你真的要退出码？Y/N")
		fmt.Scanln(&exit)
		if strings.EqualFold(exit, "Y") {
			fmt.Println("感谢使用")
			os.Exit(0)
		}
	default:
		fmt.Println("你的输入有误")
	}
}
func main() {
	//slice必须分配空间
	var account = make([]Account, 0)
	var option string
	var balance float64
	for true {
		menu()
		fmt.Scanln(&option)
		function(option, &account, &balance)
	}
}
