package main

import (
	"fmt"
)

/*
map
	map是key-value数据结构，又称为字段或者关联数组。类似其他编程语言的集合
	基本用法
		var map 变量名map[keytype]valuetype
		key可以是什么类型：
			bool、数字、string、指针、channel，还可以是只包含前面几个类型的接口、结构体、数组
			通常key为int、string
			注意，key不能是slice、map还有function，因为这几个没法用==来判断
		value可以是什么类型：
			跟key一样，通常为数组、string、map、struct
	**map只声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用。
	**key不能重复，重复赋值会覆盖，而且key是无序的
map的使用方式：
	1.先声明，再make
	2.声明直接make
	3.声明的时候直接赋值
		var m = map[string]string{"a":"b","b":"c",}最后一个逗号不能少
map增加和更新：
	map["key"] = value//如果key不存在，就是增加，如果存在就是更新
map删除
	delete(map, "key")//delete十一i个内置函数，如果key存在，就删除该key-value，如果key不存在，不操作但也不报错
	1.如果我们要删除map的所有的key，每有一个专门的方法一次删除，可以遍历一下key，诸葛删除
	2.或者map = make(...)，重新make一个
map查找
	val, flag := map1["key"]
	如果存在这个key，那么flag=true，val=value；否则flag= false，val=该类型的默认值
map的遍历：
	使用for-range的结构进行遍历
可以使用len(map)来获取map中的键值对个数
map切片
	就是map类型的数组，记得make分配内存，想要动态扩充切片容量要使用append（先创建一个map，再使用append）
map的排序
	将map中的key存入一个切片中，对切片进行排序，
map使用细节：
	1.map是引用类型，遵守引用类型传递的机制，再一个函数接收map，修改后会直接修改原来的map
	2.map的容量满了之后，在添加键值对，会自动扩容
	3.map的value也经常使用struct类型，跟适合管理复杂的数据（比value是一个map更好），
	4.当map的value也是一个map类型时，一定要make分配空间
make内置函数的说明：

*/
type Student struct {
	Id   int
	Age  int
	Name string
}

func main() {
	mapDeo()
	excise1()
}
func mapDeo() {
	var m map[string]string
	m = make(map[string]string)
	m["hello"] = "world"
	m["hello1"] = "world"
	m["hello2"] = "world"
	m["hello3"] = "world"
	m["hello4"] = "world"
	m["hello5"] = "world"
	m["hello6"] = "world"
	m["hello7"] = "world"
	m["hello8"] = "world"
	m["hello9"] = "world"
	m["hello11"] = "world"
	m["hello12"] = "world"
	fmt.Println(m)
	val, flag := m["hello33"]
	fmt.Println("val=", val, "flag", flag)
	for index, value := range m {
		fmt.Println("index=", index, "value=", value)
	}
}
func mapSlice() {
	//切片也可以使用make分配内存,而且必须要初始化长度
	var mapSlice = make([]map[string]string, 2)
	//map可以不初始化长度
	map1 := map[string]string{"hello": "world"}
	mapSlice = append(mapSlice, map1)
	fmt.Println(mapSlice)

}
func excise1() {
	//第二个map也必须要make
	map1 := make(map[string]map[string]string)
	map1["001"] = make(map[string]string)
	map1["001"]["name"] = "xiaoGou"
	map1["001"]["gender"] = "nan"
	map1["002"] = make(map[string]string)
	map1["002"]["name"] = "xiaoMao"
	fmt.Println(map1)
}

func structDemo() {
	student := Student{
		Id:   1,
		Name: "tom",
		Age:  23,
	}
	fmt.Println(student)
}
func excise2() {
	var users = make(map[string]map[string]string)

}
