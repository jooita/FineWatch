package controllers

import (
	"FineWatch/models"

	"github.com/astaxie/beego"
)

type CharController struct {
	beego.Controller
}

func (c *CharController) Index() {
	c.TplName = "characters/index.html"
}

func (c *CharController) Char() {
	charid := c.GetString("charid")
	charstruct := new(models.Character)
	c.Data["Char"] = charstruct.GetChar(charid)

	c.TplName = "characters/char.html"
}
