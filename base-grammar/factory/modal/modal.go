package modal

//创建一个首字母小写的结构体
type student struct {
	Name  string
	Score float64
}

//要使得别的包也可以使用这个结构体，可以通过工厂模式来解决
func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		Score: score,
	}
}

//如果结构体的字段首字母为小写，则也可以使用这种方法来解决
//这时候就跟Java的getter和setter相同了
