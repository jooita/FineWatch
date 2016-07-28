package controllers

import (
	"FineWatch/models"

	"github.com/astaxie/beego"
)

type MapController struct {
	beego.Controller
}

func (c *MapController) Index() {
	//Inser Info

	//map
	//new(models.Map).InsertMap()

	//char
	//new(models.Character).InsertChar()

	//maplist
	//new(models.MapList).InsertMapList()

	//CharForMap
	new(models.CharsForMap).InsertCharsForMap()

	c.TplName = "maps/index.html"
}

func (c *MapController) Map() {

	mapid := c.GetString("mapid")
	mapstruct := new(models.Map)
	c.Data["Map"] = mapstruct.GetMap(mapid)
	c.TplName = "maps/map.html"
}
