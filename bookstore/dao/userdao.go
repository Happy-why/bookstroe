package dao

import (
	"fmt"
	"github.com/model"
	"github.com/utils"
)

//CheckUserNameAndPassword 验证用户名和密码 从数据中查询一条记录
func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		fmt.Println("CheckUserNameAndPassword row.Scan err:", err)
		return nil, err
	}
	return user, nil
}

//CheckUserName 验证用户名和密码 从数据中查询一条记录
func CheckUserName(username string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users where username = ? "
	//执行
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		fmt.Println("CheckUserName row.Scan err:", err)
		return user, err
	}
	return user, nil
}

//SaveUser 向数据库中插入用户信息
func SaveUser(username string, password string, email string) error {
	//sql
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		fmt.Println("utils.Db.Exec err:", err)
		return err
	}
	return nil
}
