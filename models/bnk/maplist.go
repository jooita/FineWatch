package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type MapList struct {
	Id       int
	Map      *Map   `orm:"rel(fk)"`
	Position string `orm:"null"`
	State    string `orm:"null"`
	Explain  string `orm:"null"`
}

func (this *MapList) GetMapList(mapid string) {
	o := orm.NewOrm()
	o.Using("default")

	var maplists []orm.Params

	num, err := o.QueryTable(this).Filter("Map__Mapid", mapid).Values(&maplists)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maplists {
			fmt.Println("MapList.Get:", m["Map"], m["Position"], m["State"], m["Explain"])
		}
	}
}

func (this *MapList) InsertMapList() {
	fmt.Println("Insert MapList!")

	o := orm.NewOrm()
	o.Using("default")

	var maplist MapList
	var position string
	var state string

	for i := range mapids {

		m := Map{Mapid: mapids[i]}

		// OFFENSE - GOOD
		position = "offense"
		state = "good"
		maplist = MapList{Map: &m, Position: position, State: state}
		ErrHandle(o.Insert(&maplist))

		// OFFENSE - BAD
		state = "bad"
		maplist = MapList{Map: &m, Position: position, State: state}
		ErrHandle(o.Insert(&maplist))

		// OFFENSE - NORMAL
		state = "normal"
		maplist = MapList{Map: &m, Position: position, State: state}
		ErrHandle(o.Insert(&maplist))

		// DEFENSE - GOOD
		position = "defense"
		state = "good"
		maplist = MapList{Map: &m, Position: position, State: state}
		ErrHandle(o.Insert(&maplist))

		// DEFENSE - BAD
		state = "bad"
		maplist = MapList{Map: &m, Position: position, State: state}
		ErrHandle(o.Insert(&maplist))

		// DEFENSE - NORMAL
		state = "normal"
		maplist = MapList{Map: &m, Position: position, State: state}
		ErrHandle(o.Insert(&maplist))
	} //for mapids
}
