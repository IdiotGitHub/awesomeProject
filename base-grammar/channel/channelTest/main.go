package main

import "fmt"

func main() {
	//开启第一个协程存8000数
	go putNum(numChan)
	//开启8个协程操作
	for i := 0; i < 8; i++ {
		go operation(numChan, resultChan, exitChan)
	}
	//开启一个匿名协程关闭resultChan
	go func() {
		for i := 0; i < 8; i++ {
			v := <-exitChan
			fmt.Println("exitChan-->", v)
		}
		fmt.Println(len(resultChan))
		close(resultChan)
	}()
	for {
		num, ok := <-resultChan
		if !ok {
			break
		}
		fmt.Println(num)
	}
}

var (
	//存放数的管道
	numChan = make(chan int, 1000)
	//存放结果的管道
	resultChan = make(chan int, 2000)
	//存放标志管道,退出管道，
	exitChan = make(chan bool, 8)
)

func putNum(numChan chan int) {
	fmt.Println("putNum function start")

	for i := 1; i <= 8000; i++ {
		numChan <- i
	}
	fmt.Println("numChan length is ", len(numChan))
	close(numChan)
	fmt.Println("putNum function exit")
}
func operation(numChan chan int, resultChan chan int, exitChan chan bool) {
	fmt.Println("operation function start")
	for {
		v, ok := <-numChan
		if !ok {
			break
		}
		flag := true
		for i := 1; i <= v; i++ {
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
	fmt.Println("operation exit")
}
