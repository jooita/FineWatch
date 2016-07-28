package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type HeroForMap struct {
	Id       int `orm:"unique;pk"`
	Map      *Map
	Char     *Character
	Position string
	State    string
	explain  string
}

/*
func (this *heroForMap) GetMaps(charid string) orm.ParamsList {
	var maps []*Map
	var mapids orm.ParamsList

	o := orm.NewOrm()
	o.Using("default")

	_, err := o.QueryTable(this).Filter("Char", charid).All(&maps, "Map")
	if err == nil {
		_, err := o.QueryTable(this).ValuesFlat(&mapids, "Mapid")

		if err == nil {
			fmt.Printf("%s 's Map Names: %s", charid, strings.Join(mapids, ", "))
		}
	}
	return nil
}

*/
func (this *HeroForMap) GetChars(mapid string) {

	o := orm.NewOrm()
	o.Using("default")

	mapinfo := Map{Mapid: mapid}
	var hfms []orm.Params

	num, err := o.QueryTable(this).Filter("Map", &mapinfo).Values(&hfms)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range hfms {
			fmt.Println(m["Map"], m["Char"])
		}
	}

	/*
		_, err := o.QueryTable(this).Filter("Map", mapid).ValuesFlat(&chars, "Char.Charid")
		if err == nil {
			return chars
		}

		return nil
	*/
}

func (this *HeroForMap) Insert() {
	fmt.Println("Insert HeroForMap!")
	o := orm.NewOrm()
	o.Using("default")

	//var normal []string
	var good []string
	//var bad []string

	for i := range mapids {
		switch {
		case i < 3: //assault
			good = []string{"DVA", "LUCIO"}
			break
		case i < 6: //control
			break
		case i < 9: //escort
			break
		case i < 12: //assaultEscort
			break
		}
		m := Map{Mapid: mapids[i]}
		c := Character{Charid: good[1]}
		hfm := HeroForMap{Map: mapids[i], Char: good[1]}
		//hfm := HeroForMap{Map: &m, Char: &c, Position: "offense", State: "good"}
		id, err := o.Insert(&hfm)
		if err == nil {
			fmt.Println("HeroForMap_ID:", id)
		} else {
			fmt.Println("ERROR : ", err)
		}

	}
}
