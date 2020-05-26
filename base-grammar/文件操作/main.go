package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
文件
	文件在程序中以流的形式来操作的
流：
	数据在数据源（文件）和程序（内存）之间经理的路径
输入流：
	数据从数据源（文件）到程序（内存）的路径
输出流：
	数据从程序（内存）到数据源（文件）的路径
常用的文件操作函数和方法
	1.打开文件操作
		os.Open(path string) (file *File, err error)
	2.关闭文件
	func (f *File) Close() error
读文件操作实例
	1.带缓冲的文件读取
		bufio.NewReader(char)
	2.一次性读取一个文件
		ioutil.ReadFile(file_name) ([]byte, error)
		这个方法也不需要显示的Close这个文件，文件的Open和Close已经被封装到这个方法中
写文件操作实例
	1.func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
	说明，os.OpenFile是一个更一般性的文件打开函数，它会使用执行的选项（如，O_RDONLY)、指定的模式（如0666等）打开指定名称的文件。
	如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。
	第二个参数：文件打开模式（可以组合）
		os.O_RDWR	读写
		os.O_RDONLY	只读
		os.O_WRONLY	只写
		os.O_CREATE	创建文件
		os.O_TRUNC	清空文件内容
		os.O_APPEND	追加写
	第三个参数：权限控制（r->4,w->2,x->1)
判断文件是否存在
	os.Stat()
	func Stat(name string) (fi fileIO, err error)
	Golang判断文件或文件夹是否存在的方法为使用os.Stat()函数，利用函数返回的错误值进行判断：
		1.如果返回的错误为nil，说明文件或文件夹存在
		2.如果返回的错误类型使用os.IsNotExist()判断为true，说明文件或文件夹不存在
		3.如果返回的错误类型为其他类型，则不确定是否存在
拷贝文件
	func Copy(dst Writer, src Reader) (written int64, err error)
	io.Copy()

*/

func main() {
	//demo1()
	//demo2()
	//demo3()
	//demo4()
	CopyFile()
}

///demo1
func demo1() {
	file, err := os.Open("d:/linyidaxue.jpg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*file)
	fmt.Println(file)
	file.Close()
}

//demo2
func demo2() {
	file, err := os.Open("D:\\迅雷下载\\java介绍.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//创建一个文件指针
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
	for str, err := reader.ReadString('\n'); err != io.EOF; {
		fmt.Println(str)
	}
	fmt.Println("file read complete")

}

func demo3() {
	filePath := "D:\\迅雷下载\\java介绍.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println(string(content))
}

//创建一个新文件，写入内容为5行"hello world!"
func demo4() {
	//1.打开文件d:\\abc.txt
	filePath := "d:\\abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//准备写入5行
	str := "hello world!\n"
	//写入时使用带缓冲的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存，因此再调用WriterString方法时，是先将内容写入到缓存中，故需要flush刷新
	writer.Flush()
	/*	time.Sleep(time.Second * 20)
		err =os.Remove(filePath)
		if err != nil {
			fmt.Println(err)
		}*/
}

//copy file
func CopyFile() {
	srcFilePath := "d:\\linyidaxue.jpg"
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Println("err=", err)
	}
	reader := bufio.NewReader(srcFile)

	dstFilePath := "d:\\Download\\zhaopian.jpg"
	dstFile, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("err=", err)
	}
	defer func() {
		srcFile.Close()
		dstFile.Close()
	}()
	writer := bufio.NewWriter(dstFile)
	io.Copy(writer, reader)
	fmt.Println("copy file complete")
}
