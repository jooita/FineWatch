package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Map struct {
	Id      int
	Mapid   string `orm:"unique"`
	Class   string
	explain string
}

func (this *Map) GetMap(mapid string) *Map {
	o := orm.NewOrm()
	o.Using("default")
	mapinfo := Map{Mapid: mapid}
	err := o.Read(&mapinfo)

	if err == orm.ErrNoRows {
		fmt.Println("GetMap:", mapid, "- No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("GetMap:", mapid, "- No primary key found.")
	}

	fmt.Println("GetMap:", mapinfo.Mapid)

	return &mapinfo
}
