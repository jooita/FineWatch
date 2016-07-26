package models

import (
	_ "github.com/astaxie/beego/orm"
)

type CnC struct {
	Char1   *Character
	Char2   *Character
	State   string
	explain string
}

/*
func (this *MnC) GetMaps(charid string) orm.ParamsList {
	var maps orm.ParamsList

	o := orm.NewOrm()
	o.Using("default")

	_, err := o.QueryTable(this).Filter("Char", charid).ValuesFlat(&maps, Map.Mapid)
	if err == nil {
		return maps
	}
	return nil
}

func (this *MnC) GetChars(mapid string) orm.ParamsList {
	var chars orm.ParamsList

	o := orm.NewOrm()
	o.Using("default")

	_, err := o.QueryTable(this).Filter("Map", mapid).ValuesFlat(&chars, Char.Charid)
	if err == nil {
		return chars
	}

	return nil
}
*/
