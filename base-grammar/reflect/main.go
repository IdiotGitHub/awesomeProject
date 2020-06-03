package main

import (
	"fmt"
	"reflect"
)

/*
反射
	1.反射可以在运行时动态获取变量的各种信息，比如变量的类型（type）、类别（kind）
	2.如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）
	3.通过反射，可以修改变量的值，可以调用关联的方法。
	4.使用反射，需要导入reflect包
反射的应用场景
	1.不知道接口调用哪个函数，根据传入参数在运行时确定调用的具体接口，这种需要对函数或方法反射，
	2.对结构体序列化时，如果结构体有指定Tag，也会使用到反射生成对应的字符串
反射中重要的函数
	1.reflect.TypeOf(变量名), 获取变量的类型，返回reflect.Type类型
	2.reflect.ValueOf(变量名), 获取变量的值，返回reflect.Value类型是一个结构体类型
	3.变量、interface{}、reflect.Value是可以相互转换的
反射注意事项和使用细节
	1.reflect.Value.Kind 获取变量的类别，返回的是一个常量
	2.Type是类型，Kind是类别，Type和Kind可能是相同的，也可能是不同的
	3.通过反射可以让变量在interface{}和reflect.Value之间相互转换，
	4.使用反射的方式来获取变量的值(并返回对应的类型),
	5.通过反射来修改变量，注意当使用SetXxx方法来设置需要通过对应的指针类型来完成，这样才能改变传入的变量的值，同时需要使用到reflect.Value.Elem()方法。
*/

func main() {
	/*	num := 100
		reflectDemo1(num)*/
	/*	stu := Student{"hello"}
		reflectDemo2(stu)*/
	/*	num := 100
		reflectDemo3(&num)
		fmt.Println(num)*/
	var v float64 = 1.2
	exercise1(v)
	monster := Monster{
		"xiaoxu",
		24,
		100,
		"male",
	}
	testStruct(monster)
}

//编写一个案例，
//演示对(基本数据类型、interface{}、reflect、Value)进行反射的基本操作
func reflectDemo1(b interface{}) {
	//通过反射获取传入的变量的type、kind、值
	//1.先获取reflect.Type
	reflectType := reflect.TypeOf(b)
	fmt.Println(reflectType)
	//2.获取reflect.Value,获取到的结果类型为reflect.Value类型，要想获取实际的具体类型可以使用断言
	reflectValue := reflect.ValueOf(b)
	fmt.Println(reflectValue)
	//2.将reflectValue转成interface{}
	iv := reflectValue.Interface()
	num := iv.(int)
	fmt.Println(num)
}

type Student struct {
	Name string
}

func reflectDemo2(b interface{}) {
	//通过反射获取传入的变量的type、kind、值
	//1.先获取reflect.Type
	reflectType := reflect.TypeOf(b)
	fmt.Println(reflectType)
	//2.获取reflect.Value,获取到的结果类型为reflect.Value类型，要想获取实际的具体类型可以使用断言
	reflectValue := reflect.ValueOf(b)
	fmt.Println(reflectValue)
	//3.将reflectValue转成interface{},此时iv的类型就是该结构体类型，但是在编译阶段，Go语言不知道他的具体类型，因为反射是作用在运行时，因此还是需要断言转换
	iv := reflectValue.Interface()
	stu := iv.(Student)
	fmt.Println(stu.Name)
}

//通过反射修改函数外部的值
func reflectDemo3(b interface{}) {
	//直接获取reflect.Value
	reflectValue := reflect.ValueOf(b)
	//因为外部传进来的是一个指针类型，所以需要使用.Elem().SetXxx()修改变量值
	reflectValue.Elem().SetInt(20)
}

//给一个变量 var v float64 = 1.2,请使用反射来得到它的reflect.Value,
//然后获取对应的Type,Kind和值， 并将reflect.Value转换成interface{},
//再将interface{}转换成float64
func exercise1(b interface{}) {
	refValue := reflect.ValueOf(b)
	refType := reflect.TypeOf(b)
	refKind := refValue.Kind()
	refKind2 := refType.Kind()
	iv := refValue.Interface()
	float := iv.(float64)
	fmt.Println("refValue=", refValue)
	fmt.Println("refType=", refType)
	fmt.Println("refKind=", refKind)
	fmt.Println("refKind2=", refKind2)
	fmt.Println("float64=", float)
}

//反射最佳实践
//1.使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
//问题，为什么使用反射的时候结构体的方法参数不能是该结构体的指针类型呢
type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
	Sex   string  `json:"sex"`
}

func (m Monster) Print() {
	fmt.Println(m)
}
func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (m Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}
func testStruct(b interface{}) {
	structType := reflect.TypeOf(b)
	structValue := reflect.ValueOf(b)
	structKind := structType.Kind()
	if structKind != reflect.Struct {
		fmt.Println("请传入一个结构体")
		return
	}
	fieldsCount := structValue.NumField()
	for i := 0; i < fieldsCount; i++ {
		tag := structType.Field(i).Tag.Get("json")
		fmt.Printf("%v = %v\n", tag, structValue.Field(i))
	}
	methodCount := structValue.NumMethod()
	fmt.Println("struct has", methodCount, " methods")
	//执行第二个方法，
	//反射中的结构体方法的排序是按照字母顺序进行排序
	//也可以通过MethodByName()来调具体的方法
	structValue.Method(1).Call(nil)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	result := structValue.Method(0).Call(params)
	fmt.Println("result[0]", result[0].Int())
}
