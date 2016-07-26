package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModelWithPrefix("tb_", new(Map))
	orm.RegisterModelWithPrefix("tb_", new(Character))
}
