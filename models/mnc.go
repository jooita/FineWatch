package models

import "github.com/astaxie/beego/orm"

type MnC struct {
	Map     *Map
	Char    *Character
	State   string
	explain string
}

/*
func (this *MnC) GetMaps(charid string) orm.ParamsList {
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

func (this *MnC) GetChars(mapid string) orm.ParamsList {
	var chars orm.ParamsList

	o := orm.NewOrm()
	o.Using("default")

	_, err := o.QueryTable(this).Filter("Map", mapid).ValuesFlat(&chars, "Char.Charid")
	if err == nil {
		return chars
	}

	return nil
}
