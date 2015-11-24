package routers

import (
	"github.com/astaxie/beego"
	"music/controllers"
)

func init() {
	beego.Router("/", &controllers.WelcomeController{})
}
