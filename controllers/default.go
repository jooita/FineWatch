package controllers

import (
	"FineWatch/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	logstruct := new(models.Log)
	c.Data["Log"] = logstruct.Get("베지밀-3657", "kr")
	c.TplName = "index.html"
}

func (c *MainController) Log() {
	tagid := c.GetString("tagid")
	region := c.GetString("regionid")

	logstruct := new(models.Log)
	c.Data["Log"] = logstruct.Get(tagid, region)

	c.TplName = "log.html"
}
