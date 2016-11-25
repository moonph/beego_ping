package routers

import (
	"github.com/moonph/beego_ping/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
