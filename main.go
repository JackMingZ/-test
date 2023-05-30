package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "test/models"
	_ "test/routers"
)

func main() {
	beego.Run()
}
