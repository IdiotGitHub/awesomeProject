package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/**
传统测试方法的缺点：
	1.不方便，我们需要在main函数中去调用，这样就需要去修改main函数，如果现在项目正在运行，就可能去停止项目
	2.不利于管理，当我们测试多个函数或者多个模块时，都需要在main函数中，不利于管理和清晰思路
Go语言中自带一个轻量级的测试框架testing和自带go test命令来实现单元测试和性能测试，testing框架和其它语言中的测试框架类似，
可以基于这个框架写针对响应函数的测试用例，也可以基于该框架写响应的压力测试。通过单元测试可以解决如下问题。
	1.确保每个函数是可运行，并且运行结果是正确的。
	2.确保写出来的代码性能是好的
	3.单元测试能及时的发现程序设计或实现的逻辑错误，是问题及早暴漏，便于问题的定位解决，而性能测试的重点在于发现程序设计上的一些问题，让程序能够在高并发的情况下海恩那个保持稳定
测试用例细节：
	1.测试用例文件名必须以_test.go结尾，
	2.测试用例函数必须以Test开头，一般来说就是Test——被测试的函数名，Test后面的第一个字母不能是小写的
	3.一个测试用例文件中，可以有多个测试用例函数，
	4.测试用例的参数类型必须是*test.T
	5.运行测试用例指令
		1.cmd>go test [如果运行成功，无日志，会有错误日志]
		2.cmd>go test -v [运行正确或是错误都会输出日志]
	6.当错误出现时，可以使用t.Fatalf来格式化输出错误信息，并退出程序
	7.t.Logf方法可以输出响应的日志
	8.测试用例函数，并没有放在main函数中，也执行了，这就是测试用例的方便之处
	9.PASS表示测试用例运行成功，FAIL表示测试用例运行失败
	10.测试单个文件，一定要带上被测试的源文件
		go test -v name_test.go name.go
	11.测试单个方法
		go test -v -test.run func
*/
func main() {
	/*	monster := Monster{
			"DaXu",
			24,
			"super power",
		}
		monster.Store()
	*/
	monster := Monster{}
	monster.ReStore()
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func (m *Monster) Store() {
	file, err := os.OpenFile("d:\\Monster_json.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("OpenFile Error", err)
		return
	}
	monster, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json format error", err)
	}
	writeString, err := file.WriteString(string(monster))
	if err != nil {
		fmt.Println("WriteString to Monster_json.txt error", err)
		return
	}
	fmt.Println("WriteString to Monster_json.txt success,the length of string is", writeString)

}
func (m *Monster) ReStore() {
	readFile, err := ioutil.ReadFile("d:\\Monster_json.txt")
	if err != nil {
		fmt.Println("Read file error", err)
		return
	} else {
		fmt.Println("Read file success ,the file  ", readFile)
	}
	err = json.Unmarshal(readFile, m)
	if err != nil {
		fmt.Println("ummarshal Monster error ", err)
		return
	}
	fmt.Println("ummarshal Monster success,monster is", *m)
}
