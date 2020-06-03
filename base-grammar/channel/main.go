package main

import (
	"fmt"
	"sync"
)

/**
解决资源竞争问题，低水平程序员使用加锁的方式，高水平程序员使用channel
channel 管道
	1.channel本质上就是一个数据结构--队列
	2.数据是先进先出的[FIFO]
	3.线程安全，多goroutine访问时，不需要加锁
	4.channel有类型，
定义/声明
 	var varName chan type
	管道是引用类型
注意事项：
	1.channel中只能存放指定的数据类型
	2.channel的数据放满后，就不能再放入
	3.如果channel取出数据后，可以继续放入
	4.在没有使用协程的情况下，如果channel数据取完，再继续取数据，会发生deadlock
channel的关闭
	使用内置函数close可以关闭channel，channel一旦关闭就不可以再继续写入数据，但是仍然可以取数据
channel的遍历
	channel支持for-range的方式进行遍历，
		1.在遍历时，如果channel没有关闭，则会出现deadlock
		2.在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，正常退出
如果一直往一个channel中写入数据，那么在达到channel容量上限的时候会发生deadlock错误，但是如果边写边读的话就会发生阻塞，（即使读慢写快，也不会发生deadlock，此时编译器会知道有读的协程，写慢读快同理）
在使用管道的时候就是要有写有读，然后记得关就可以了
使用细节
	1.channel可以声明为只读或者只写（在默认情况下是双向的）
		1.声明为只写	var chan1 chan<- int
		1.声明为只读	var chan2 <-chan int
		这种情况可以写在形参中，防止其他的误操作（在外部声明为双向的，在形参使用单向的）
	2.使用select可以解决从管道取数据的阻塞问题
		可以在for{}中使用select语句，当取不到数据的时候也不会发生传统方式中的fatal，而是跳到下一个case中，都取不到最后跳转至default
	for {
		select {
			case v := <-channel1:
				...
			case v := <-channel2:
				...
			default :
				return或者break或其他处理逻辑
			}
		}
	3.如果在协程中发生了panic而没有进行捕获会导致整个程序崩溃，可以在可能发生panic的协程中使用defer-recover进行错误捕获，这样主线程和其他的协程就不会受到影响了
*/
type Cat struct {
	Name string
}

var myMap = make(map[int]uint64, 10)
var lock sync.Mutex
var interfaceChan chan interface{} = make(chan interface{}, 3)

func test(n int) {
	var res uint64 = 1
	for i := 1; i <= n; i++ {
		res *= uint64(i)
	}
	//会发生资源竞争,第一种解决方案是使用枷锁
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
func main() {
	/*for i := 0; i <= 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	for i, v := range myMap {
		fmt.Println("myMap[", i, "]=", v)
	}
	*/
	/*cat := Cat{"hello"}
	interfaceChan <- cat
	newCat := <-interfaceChan
	//类型断言
	cat = newCat.(Cat)
	fmt.Println(cat)
	*/
	/*	intChan := make(chan int, 50)
		exitChan := make(chan bool, 1)
		go writeData(intChan)
		go readData(intChan, exitChan)
		for {
			if _, ok := <-exitChan; !ok {
				break
			}
		}*/
	//开启第一个协程存8000数
	go putNum(numChan)
	//开启8个协程操作
	for i := 0; i < 8; i++ {
		go operation(numChan, resultChan, exitChan)
	}
	//开启一个匿名协程关闭resultChan
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
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
func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData", v)
	}
	exitChan <- true
	close(exitChan)
}

//求8000以内的素数，使用协程，管道的方式
//使用三个管道，一个管道用于存8000个数，一个管道用于存结果，一个管道用于存协程的完成标志
//8000个数存完后关闭第一个管道
//判断完成后将结果存入结果管道，在协程完成之后将运行完成标志存入标志管道，注意此处不能关闭任何管道，因为不知道其他协程是否运行完成
//主线程中再新开一个协程遍历标志管道，此处明确知道写成数量，所以不适用死循环来取，最后关闭结果管道-->疑问，此处没有关闭标志管道为啥不会发生deadlock？是因为发生deadlock的先决条件是，取完数据继续取。此处明确知道长度，所以不会发生deadlock。
//最后遍历结果管道将结果打印
var (
	//存放数的管道
	numChan = make(chan int, 1000)
	//存放结果的管道
	resultChan = make(chan int, 2000)
	//存放标志管道,退出管道，
	exitChan = make(chan bool, 8)
)

func putNum(numChan chan int) {
	for i := 1; i <= 8000; i++ {
		numChan <- i
	}
	close(numChan)
	fmt.Println("putNum function exit")
}
func operation(numChan chan int, resultChan chan int, exitChan chan bool) {
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
