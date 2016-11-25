package controllers

import (
	"github.com/astaxie/beego"
)

type PingController struct {
	beego.Controller
}

func (c *PingController) Get() {
	c.Data["json"] = map[string]string{"Message": "pong"}
	c.ServeJSON()
}
