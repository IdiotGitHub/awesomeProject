package utils

import "fmt"

var(
	Age int
	Name string
)
func init() {
	fmt.Println("utils->init()")
	Age = 10
	Name = "jack"
}
