package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type CharsForMap struct {
	Id      int
	Maplist *MapList   `orm:"rel(fk)"`
	Charid  *Character `orm:"rel(fk)"`
	Explain string     `orm:"null"`
}

func (this *CharsForMap) GetChars(mapid string) orm.ParamsList {
	var charids orm.ParamsList

	o := orm.NewOrm()
	o.Using("default")

	_, err := o.QueryTable(this).Filter("Maplist__Map__Mapid", mapid).ValuesFlat(&charids, "Charid")
	if err == nil {
		fmt.Printf("%s 's Characters:", mapid)
		fmt.Println(charids)
		return charids
	}
	return nil
}

func (this *CharsForMap) GetMaps() orm.ParamsList {
	var maplists orm.ParamsList

	o := orm.NewOrm()
	o.Using("default")

	_, err := o.QueryTable(this).ValuesFlat(&maplists, "Maplist")
	if err == nil {
		//fmt.Printf("%s", strings.Join(maplists, ", "))
		fmt.Println(maplists)
		return maplists
	}
	return nil
}

func (this *CharsForMap) InsertCharsForMap() {

	fmt.Println("Insert CharsForMap!")

	o := orm.NewOrm()
	o.Using("default")

	var cfm CharsForMap
	var listid int

	for i := range mapids { //order : offense:good,bad,normal / defense:good,bad,normal
		listid = (6 * i) + 1

		switch {
		case i < 3: //Assault
			// OFFENSE - GOOD
			cfm = CharsForMap{Maplist: &MapList{Id: listid}, Charid: &Character{Charid: "dva"}, Explain: ""}
			ErrHandle(o.Insert(&cfm))
			// OFFENSE - BAD
			listid++ //2,8,14

			// OFFENSE - NORMAL
			listid++ //3,9,15

			// DEFENSE - GOOD
			listid++ //4,10,16

			// DEFENSE - BAD
			listid++ //5,11,17

			// DEFENSE - NORMAL
			listid++ //6,12,18

			if mapids[i] == "anubis" {
				cfm = CharsForMap{Maplist: &MapList{Id: 3}, Charid: &Character{Charid: "ana"}, Explain: ""}
				ErrHandle(o.Insert(&cfm))
			}

			break
		case i < 6: //Control

			// OFFENSE - GOOD
			// OFFENSE - BAD
			listid++
			cfm = CharsForMap{Maplist: &MapList{Id: listid}, Charid: &Character{Charid: "lucio"}, Explain: ""}
			ErrHandle(o.Insert(&cfm))

			// OFFENSE - NORMAL
			listid++

			// DEFENSE - GOOD
			listid++

			// DEFENSE - BAD
			listid++

			// DEFENSE - NORMAL
			listid++

			if mapids[i] == "nepal" {
				cfm = CharsForMap{Maplist: &MapList{Id: 27}, Charid: &Character{Charid: "ana"}, Explain: ""}
				ErrHandle(o.Insert(&cfm))
			}

			break
		case i < 9: //Escort
			break
		case i < 12: //AssaultEscort
		} //switch

	}

}
