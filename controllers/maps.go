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

	//CharForMap
	//new(models.CharsForMap).InsertCharsForMap()

	class := []string{"assault", "control", "escort", "assaultEscort"}
	mapstruct := new(models.Map)
	for _, v := range class {
		c.Data[v] = mapstruct.GetClassMap(v)
	}
	c.TplName = "maps/index.html"
}
func (c *MapController) CharsForMap() {
	mapid := c.GetString("mapid")
	mapstruct := new(models.Map)
	cfmstruct := new(models.CharsForMap)
	c.Data["Map"] = mapstruct.GetMap(mapid)
	c.Data["Offense"] = cfmstruct.GetOf(mapid)
	c.Data["Defense"] = cfmstruct.GetDf(mapid)

	c.TplName = "maps/charsformap.html"
}

func (c *MapController) Map() {

	mapid := c.GetString("mapid")
	mapstruct := new(models.Map)
	c.Data["Map"] = mapstruct.GetMap(mapid)

	c.TplName = "maps/map.html"
}
