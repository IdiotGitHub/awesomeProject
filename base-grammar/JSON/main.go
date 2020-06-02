package main

import (
	. "encoding/json"
	"fmt"
)

/*
JSON
	JSON是一种轻量级的数据交换格式，已于人阅读和编写。同时也易于及其解析和生成
	JSON易于及其解析和生成。并有效的提升网络传输效率，通成程序在网络传输时会先将数据（结构体、map等）序列化成json字符串，到接收方得到json字符串时，再反序列化恢复成原来的数据类型（结构体、map等）。
JSON数据格式：
	在JS语言中，一切都是对象，因此，任何的数据类型都可以通过JSON来表示
JSON键值对是用来保存数据的一种方式：
JSON序列化：

*/
type Monster struct {
	Name     string  `json:"name"` //使用的是反射机制
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func JsonTest() {
	var monster = Monster{
		"xiaoxu",
		24,
		"1995-11-11",
		100000,
		"awesome",
	}
	data, err := Marshal(monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func sliceTest() {
	var slice []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 23
	m1["address"] = "beijing"
	slice = append(slice, m1)
	m2 := make(map[string]interface{})
	m2["name"] = "jack"
	m2["age"] = 23
	m2["address"] = "beijing"
	slice = append(slice, m2)
	data, err := Marshal(slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
func unmarshalStruct() {
	str := "{\"name\":\"xiaoxu\",\"age\":24,\"birthda\":\"1995-11-11\",\"sal\":100000,\"skill\":\"awesome\"}"
	var monster Monster
	err := Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(monster)
}
func unmarshalMap() {
	str := "{\"address\":\"beijing\",\"age\":23,\"name\":\"jack\"}"
	var m map[string]interface{}
	//反序列化map时，在底层已经进行make了
	err := Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}
func main() {
	JsonTest()
	sliceTest()
	unmarshalStruct()
	unmarshalMap()
}
