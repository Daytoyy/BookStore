package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

//根据用户名和密码从数据库中查询一条记录
func CheckUsernameAndPssword(username string, password string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username, password) //QueryRow执行一次返回一行结果
	user := &model.User{}
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	return user, nil
}

//根据用户名从数据库中查询一条记录
func CheckUserName(username string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users where username = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username) //QueryRow执行一次返回一行结果
	user := &model.User{}
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	return user, nil
}
func SaveUser(username string, password string, email string) error {
	//写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		fmt.Println("预编译和执行出错:", err)
		return err
	}
	return nil
}
