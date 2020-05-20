package main

import "fmt"

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
*/
func main() {
	var arr [9]int = [9]int{1, 21, 3, 8, 12, 34, 22, 44, 54}
	maoPao(&arr)
	fmt.Println(arr)
	shunXuChaZhao()
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
