package controllers

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.Data["data"] = "caonima"
	c.Data["data"] = "xibaxiba"
	c.TplName = "testHtmk.html"
}
func (c *MainController) Post() {
	c.Data["data"] = "\n\n\n\n\n牛马上身！"
	c.TplName = "testHtmk.html"
}
