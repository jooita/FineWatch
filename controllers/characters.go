package controllers

import (
	"github.com/astaxie/beego"
)

type CharController struct {
	beego.Controller
}

func (c *CharController) Index() {
	c.TplName = "characters/index.html"
}
