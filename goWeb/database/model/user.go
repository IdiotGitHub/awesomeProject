package model

import (
	"awesomeProject/goWeb/database/utils"
	"fmt"
)

//create User struct
type User struct {
	Id     int
	Name   string
	Gender int
	Email  string
}

//AddUser
func (u *User) AddUser() (err error) {
	sqlStr := "insert into user_info(name, gender, email) values(?, ?, ?);"
	//use prepare statement
	statement, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare sql error: ", err)
	}
	_, err = statement.Exec("admin", "1", "331538695@qq.com")
	if err != nil {
		fmt.Println("execute sql error :", err)
		return
	}
	return nil
}

//AddUser2
func (u *User) AddUser2() (err error) {
	sqlStr := "insert into user_info(name, gender, email) values(?, ?, ?);"
	//don't use prepare statement
	_, err = utils.Db.Exec(sqlStr, "admin2", "1", "helloWorld@hello.com")
	if err != nil {
		fmt.Println("execute sql error :", err)
		return
	}
	return nil
}

//select user from database
func (u *User) GetUserById() (user *User, err error) {
	sqlStr := "select * from user_info where id = ?;"
	statement, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return
	}
	row := statement.QueryRow(u.Id)
	//create vars to save result
	var id int
	var name string
	var gender int
	var email string
	err = row.Scan(&id, &name, &gender, &email)
	if err != nil {
		return
	}
	user = &User{
		Id:     id,
		Name:   name,
		Gender: gender,
		Email:  email,
	}
	return
}

//getUsers
func (u *User) GetUsers() (users []*User, err error) {
	sqlStr := "select * from user_info"
	statement, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return
	}
	rows, err := statement.Query()
	if err != nil {
		return
	}
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Gender, &user.Email)
		if err != nil {
			fmt.Println("scan data error", err)
		}
		users = append(users, user)
	}

	return
}
