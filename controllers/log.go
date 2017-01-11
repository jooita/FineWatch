package controllers

import (
	"FineWatch/models"

	"github.com/astaxie/beego"
)

type LogController struct {
	beego.Controller
}

func (c *LogController) Index() {
	tag := c.GetString("battletag")
	region := c.GetString("region")

	profilestruct := new(models.Log)
	c.Data["Profile"] = profilestruct.GetLog(tag, region)

	c.TplName = "log/index.html"
}

func (c *LogController) Log() {

}
