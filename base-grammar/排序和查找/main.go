package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
排序何查找
	排序：
		1.排序是将一组数据，按指定的顺序进行排列的过程
		2.排序的分类：
			1.内部排序：指将需要处理的所有数据都加载到内部存储器中进行排序。包括（交换式排序法、选择式排序法、插入式排序法）
			2.外部排序法：数据量过大，无法全部加载到内存中，需要借助外部存储进行排序。包括（合并排序法和直接合并排序法）
	1.交换式排序法：交换式排序属于内部排序法，是运用数据值比较后，按判断规则对数据位置进行交换，以达到排序的目的
		1，冒泡排序法
			冒泡排序的基本思想是：通过对待排序序列从后向前（从下标较大的元素开始），一次比较相邻元素的排序码，若发现逆序则交换，使排序码较小的元素逐渐从后向前部移动
			因为排序的过程中，个元素不断地接近自己的位置，如果一趟比较下来没有进行过交换，就说明序列有序，因此要在排序过程中设置一个标志判断元素是否进行过交换，从而减少不必要的比较
		2.快速排序法
	查找
		1.顺序查找
		2.二分查找
			推出条件：1.找到2.找不到:左边下标大于右边下标
*/
func main() {
	//var arr [9]int = [9]int{1, 21, 3, 8, 12, 34, 22, 44, 54}
	//maoPao(&arr)
	//fmt.Println(arr)
	//shunXuChaZhao()
	//erFen(&arr, 0, len(arr) - 1, 8)
	//excise1()
	//excise2()
	//excise3()
	excise4()
}
func maoPao(arr *[9]int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i] = arr[i] + arr[j]
				arr[j] = arr[i] - arr[j]
				arr[i] = arr[i] - arr[j]
			}
		}
	}
	fmt.Println(arr)
}
func shunXuChaZhao() {
	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName string
	fmt.Println("请输入查找的人名")
	fmt.Scanln(&heroName)
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			fmt.Println("查找成功，下标为:", i)
			break
		} else if i == len(names)-1 {
			fmt.Println("查找失败")
		}
	}
}

//erfen cha zhao
func erFen(arr *[10]int, leftIndex int, rightIndex int, desNum int) {
	fmt.Println(*arr)
	middleNum := (leftIndex + rightIndex) / 2
	if leftIndex > rightIndex {
		fmt.Println("未找到")
		return
	}
	if (*arr)[middleNum] > desNum {
		erFen(arr, middleNum+1, rightIndex, desNum)
	} else if (*arr)[middleNum] < desNum {
		erFen(arr, leftIndex, middleNum-1, desNum)
	} else {
		fmt.Println("已找到该数,下标为：", middleNum)
	}
}

//excise
//excise1	随机生成10个整数(1_100)保存到数组，并顺序打印以及求平均值、最大值和最小值的下标，最后查找是否包含55
func excise1() {
	rand.Seed(time.Now().UnixNano())
	var max, maxIndex, min, minIndex, sum int
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
		sum += arr[i]
	}
	fmt.Println("arr init:", arr)
	max = arr[0]
	min = max
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			maxIndex = i
		}
		if min > arr[i] {
			minIndex = i
		}
	}
	//排序
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i] += arr[j]
				arr[j] = arr[i] - arr[j]
				arr[i] = arr[i] - arr[j]
			}
		}
	}
	fmt.Println("maxIndex=", maxIndex, "minIndex=", minIndex)
	fmt.Println("arr:", arr)
	fmt.Println("avg=", float64(sum)/float64(len(arr)))
}

//excise2 已知有序数组（升序），要求插入一个元素，最后打印该数组，顺序依然是升序
func excise2() {
	var sortArr = [5]int{23, 34, 56, 67, 78}
	var num, index int
	fmt.Println("请输入要插入的元素：")
	fmt.Scanln(&num)
	for i := 0; i < len(sortArr); i++ {
		if num < sortArr[i] {
			index = i
			break
		}
	}
	sortArrSlice := sortArr[:]
	var slice = make([]int, index)
	copy(slice, sortArrSlice)
	slice = append(slice, num)
	slice = append(slice, sortArr[index:]...)
	fmt.Println("slice=", slice)
}

//excise3	已知数组arr[10]string，查找"AA"在其中是否存在，如果有多个"AA",找出下标
func excise3() {
	var arr = [10]string{"AA", "BB", "CC", "AA", "DD", "AA", "EE", "FF", "AA", "GG"}
	var slice = make([]int, 0)
	for index, value := range arr {
		if value == "AA" {
			slice = append(slice, index)
		}
	}
	if len(slice) == 0 {
		fmt.Println("未找到AA")
	} else {
		fmt.Println("找到AA，其下标分别为:", slice)
	}
}

//excise4 	随机生成10个数（1-100），使用冒泡排序法进行排序，然后查找是否有90这个数，并显示其下标，没有则提示没有
func excise4() {
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i] += arr[j]
				arr[j] = arr[i] - arr[j]
				arr[i] = arr[i] - arr[j]
			}
		}
	}
	erFen(&arr, 0, len(arr)-1, 90)
}
