package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Character struct {
	Id      int
	Charid  string `orm:"unique"`
	Class   string
	explain string
}

func (this *Character) GetChar(charid string) *Character {
	o := orm.NewOrm()
	o.Using("default")
	char := Character{Charid: charid}
	err := o.Read(&char)

	if err == orm.ErrNoRows {
		fmt.Println("GetChar:", charid, "- No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("GetChar:", charid, "- No primary key found.")
	}

	return &char
}
