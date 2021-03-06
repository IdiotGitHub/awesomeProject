package main

import "fmt"

/*
原码、反码、补码
	1.二进制的最高位是符号位，0为正、1为负
	2.整数的原码、反码、补码都一样
	3.负数的反码等于它的原码符号位不变，其他位取反
	4.负数的补码等于它的补码加一
	5.0的反码，补码都是0
	6.在计算机运行的时候，都是以补码的方式来运算的
按位与&、按位或|、按位异或^
	1.两位全为1，结果为1，否则为0（有零则为零，一一才为一）
	2.两位有一个为1，结果为1，否则为0（有1则为一，零零才为零）
	3.两位有一个为0一个为1，结果为1，否则为零（不同则为1，相同则为零）
*/

func main() {
	fmt.Println(-2 & 2)
	fmt.Println(-3 | 2)
}
