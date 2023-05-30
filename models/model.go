package models

import (
	"database/sql"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

// 操作数据库代码
func init() {
	//第一个参数是数据库驱动
	//连接数据库字符串
	//"root:hjj156586044...@tcp(127.0.0.1:3306)"
	//用户名:密码@tcp(127.0.0.1:3306)/数据库名称?charset=utf8
	conn, err := sql.Open("mysql", "root:hjj156586044...@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic("连接失败" + err.Error())
		return
	}

	//创建表
	_, err2 := conn.Exec("create table IF NOT EXISTS itcast(name varchar(32),password varchar(50));")

	if err2 != nil {
		panic("连接失败" + err2.Error())
	}

	//增删改查
	//查询数据是否存在
	res, _ := conn.Query("SELECT * FROM itcast WHERE name = ?", "串子")
	if !res.Next() { //如果没有查询到数据，则插入
		conn.Exec("INSERT INTO itcast(name,password) VALUES (?,?)", "串子", "黑马")
	}
	ress, err := conn.Query("select name from itcast")
	if err != nil {
		panic(err)
	}
	var name string
	for ress.Next() {
		ress.Scan(&name)
		if err != nil {
			panic(err)
		}
		logs.Info(name)
	}
	//关闭数据库
	defer conn.Close()
}
