package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)
var db  *sql.DB

var _usql = "select id,name from test where id = ?"

type User struct{
	Id int64
	Name string
}

func main() {
	if user,err := daoFunc(1);err != nil {
		fmt.Println("query user id = 1 ,errs: %v",err.Error())
	}else if(user == nil){
		fmt.Println("query user id = 1 row is empty")
	}else{
		fmt.Printf("query user id = 1 ,name = %s \n",user.Name)
	}
}

//模拟dao层业务，sql.ErrNoRows应该属于正常业务，查询无结果，不应该往上抛
func daoFunc(id int) (user *User,err error) {
	user = new(User)
	if err = db.QueryRow(_usql,id).Scan(&user.Id,&user.Name);err != nil {
		if err == sql.ErrNoRows {
			return nil,nil
		}
		err = errors.Wrap(err,"query user by id err")
	}
	return user,err
}

//初始化db连接，不关闭连接
func init() {
	if db != nil {
		return
	}
	var err error
	// 初始化连接池
	if db, err = sql.Open("DbDriverName","xxxx"); err != nil {
		//初始化db异常，直接panic
		panic(fmt.Sprintf("main:init_db_err\t%v", err))
	}
	// 设置连接池信息
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(200)
}
