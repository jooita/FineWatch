package routers

import (
	"FineWatch/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/maps", &controllers.MapController{}, "get:Index")
	beego.Router("/characters", &controllers.CharController{}, "get:Index")

	beego.Router("/maps/charsformap", &controllers.MapController{}, "get:CharsForMap")
	beego.Router("/characters/char", &controllers.CharController{}, "get:Char")

	beego.Router("/maps/map", &controllers.MapController{}, "get:Map")
}
