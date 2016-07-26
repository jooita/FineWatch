package controllers

import (
	"FineWatch/models"
	"github.com/astaxie/beego"
)

type MapController struct {
	beego.Controller
}

func (c *MapController) Index() {
	c.TplName = "maps/index.html"
}

func (c *MapController) Map() {
	mapid := c.GetString("mapid")
	mapstruct := new(models.Map)
	mapstruct.GetMap(mapid)
	c.TplName = "maps/map.html"
}
