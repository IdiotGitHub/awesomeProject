package modal

import "fmt"

//不能随便查看人的性别工资

type person struct {
	Name   string
	age    int
	salary float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确。。")
	}
}
func (p *person) GetAge() int {
	return p.age
}
func (p *person) SetSalary(salary float64) {
	if salary >= 3000 && salary <= 30000 {
		p.salary = salary
	} else {
		fmt.Println("工资范围错误")
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}

//Account结构体，要求有字段：账号（长度6-10位之间）、余额（必须大于20）、密码（必须是6位）
type Account struct {
	accountNo string
	password  string
	balance   float64
}

func NewAccount() *Account {
	return &Account{}
}
func (a *Account) SetAccountNo(num string) {
	if len(num) >= 6 && len(num) <= 10 {
		a.accountNo = num
	} else {
		fmt.Println("请输入6-10位账号")
	}
}
func (a *Account) GetAccount() string {
	return a.accountNo
}
func (a *Account) SetPassword(password string) {
	if len(password) == 6 {
		a.password = password
	} else {
		fmt.Println("请输入6位密码")
	}
}
func (a *Account) GetPassword() string {
	return a.password
}
func (a *Account) SetBalance(balance float64) {
	if balance > 20 {
		a.balance = balance
	} else {
		fmt.Println("请输入正确的余额")
	}
}
func (a *Account) GetBalance() float64 {
	return a.balance
}
