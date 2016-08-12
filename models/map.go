package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Map struct {
	Mapid   string `orm:"pk"`
	KMapid  string
	Class   string
	Explain string
}

func (this *Map) GetAll() []*Map {
	var maps []*Map
	o := orm.NewOrm()
	o.Using("default")
	o.QueryTable(this).All(&maps)

	return maps
}

func (this *Map) GetClassMap(class string) []*Map {
	var maps []*Map
	o := orm.NewOrm()
	o.Using("default")
	o.QueryTable(this).Filter("Class", class).All(&maps)

	return maps
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
	} else {
		fmt.Println("GetMap:", mapinfo.Mapid, mapinfo.Class, mapinfo.Explain)
	}

	return &mapinfo
}

func (this *Map) InsertMap() {

	fmt.Println("Insert Map!")

	o := orm.NewOrm()
	o.Using("default")
	for i := range mapids {
		maps := Map{Mapid: mapids[i], KMapid: kmapids[i]}
		switch {
		case i < 3:
			maps.Class = "assault"
			break
		case i < 6:
			maps.Class = "control"
			break
		case i < 9:
			maps.Class = "escort"
			break
		case i < 12:
			maps.Class = "assaultEscort"
		}
		id, err := o.Insert(&maps)
		if err == nil {
			fmt.Println("Map ID :", id)
		}
	}
}
