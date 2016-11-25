package routers

import (
	"github.com/astaxie/beego"
	"github.com/moonph/beego_ping/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ping", &controllers.PingController{})
}
