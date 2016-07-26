package routers

import (
	"FineWatch/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/maps", &controllers.MapController{}, "get:Index")
	beego.Router("/characters", &controllers.CharController{}, "get:Index")

	beego.Router("/maps/map", &controllers.MapController{}, "get:Map")
}
