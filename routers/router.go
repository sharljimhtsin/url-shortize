package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"website/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/", &controllers.MainController{}, "post:Add")
	beego.Router("/dump", &controllers.MainController{}, "get:List")
	beego.Router("/welcome", &controllers.MainController{}, "get:Welcome")
	beego.Router("/:hash", &controllers.MainController{}, "get:Query")
}
