package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"test/controllers"
)
import _ "github.com/go-sql-driver/mysql"

func init() {
	beego.Router("/", &controllers.MainController{})
	/*//给请求指定自定义方法,一个请求指定一个方法
	beego.Router("/login",
		&controllers.LoginController{},
		"get:ShowLogin;post:PostFunc")
	//给多个请求指定一个方法
	beego.Router("/index",
		&controllers.IndexController{},
		"get,post:HandleFunc")
	//给所有请求指定一个方法
	beego.Router("index", &controllers.IndexController{}, "*:HandleFunc")
	//当两种指定方法冲突的时候，范围越小优先级越高
	*/
}
