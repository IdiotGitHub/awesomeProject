package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial error", err)
		return
	}
	fmt.Println("client dial success", dial)
	//使用bufio包中的reader来获取终端输入
	reader := bufio.NewReader(os.Stdin)
	for {
		//从终端获取一行输入
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("获取输入错误", err)
		}
		readString = strings.Trim(readString, "\r\n")
		n, err := dial.Write([]byte(readString))
		if err != nil {
			fmt.Println("发送数据错误", err)
		}
		fmt.Println("消息发送成功", n)
	}
}
