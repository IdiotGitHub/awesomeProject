package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//测试使用协程耗时
	start := time.Now().Nanosecond()
	//开启第一个协程存8000数
	go putNum(numChan)
	//开启8个协程操作
	for i := 0; i < numCpu; i++ {
		go operation(numChan, resultChan, exitChan)
	}
	//开启一个匿名协程关闭resultChan
	go func() {
		for i := 0; i < numCpu; i++ {
			<-exitChan
		}
		end := time.Now().Nanosecond()
		fmt.Println(start - end)
		close(resultChan)
	}()
	//打印结果
	for {
		_, ok := <-resultChan
		if !ok {
			break
		}
		//fmt.Println(num)
	}
	fmt.Println(numCpu)
}

var (
	//获取cpu数量
	numCpu = runtime.NumCPU()
	//存放数的管道
	numChan = make(chan int, 1000)
	//存放结果的管道
	resultChan = make(chan int, 2000)
	//存放标志管道,退出管道，
	exitChan = make(chan bool, numCpu)
)

func putNum(numChan chan int) {

	for i := 1; i <= 8000; i++ {
		numChan <- i
	}
	close(numChan)
}
func operation(numChan chan int, resultChan chan int, exitChan chan bool) {
	for {
		v, ok := <-numChan
		if !ok {
			break
		}
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resultChan <- v
		}
	}
	exitChan <- true
	//打印消息是很耗时的
	//fmt.Println("operation exit")
}
