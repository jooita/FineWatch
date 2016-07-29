package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type CharsForMap struct {
	Id       int
	Map      *Map       `orm:"rel(fk)"`
	Char     *Character `orm:"rel(fk)"`
	Position string
	Explain  string `orm:"null"`
	Count    int    `orm:"default(0)"`
}

func (this *CharsForMap) GetCharsForMap(mapid string) []*CharsForMap {
	var cfminfo []*CharsForMap
	o := orm.NewOrm()
	o.Using("default")
	o.QueryTable(this).Filter("Map__Mapid", mapid).RelatedSel().All(&cfminfo)

	return cfminfo
}

func (this *CharsForMap) GetOf(mapid string) []*CharsForMap {
	var cfminfo []*CharsForMap
	o := orm.NewOrm()
	o.Using("default")
	o.QueryTable(this).Filter("Map__Mapid", mapid).Filter("Position", "offense").RelatedSel().All(&cfminfo)

	return cfminfo
}
func (this *CharsForMap) GetDf(mapid string) []*CharsForMap {
	var cfminfo []*CharsForMap
	o := orm.NewOrm()
	o.Using("default")
	o.QueryTable(this).Filter("Map__Mapid", mapid).Filter("Position", "defense").RelatedSel().All(&cfminfo)

	return cfminfo
}

func (this *CharsForMap) InsertCharsForMap() {

	fmt.Println("Insert CharsForMap!")

	o := orm.NewOrm()
	o.Using("default")

	cfm := make(chan *CharsForMap, 100)

	go func(cfm chan *CharsForMap) {
		for {
			id, err := o.Insert(<-cfm)
			if err == nil {
				fmt.Println("CharsForMap Insert :", id)
			}
		}
	}(cfm)

	var pos string

	for i := range mapids { //order : offense:good,bad,normal / defense:good,bad,normal

		switch {
		case i < 3: //Assault

			///////////////COMMON
			// OFFENSE
			pos = "offense"
			for _, ch := range []string{"pharah", "genji", "reaper"} {
				cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
			}

			// DEFENSE
			pos = "defense"
			for _, ch := range []string{"bastion", "torbjorn", "reinhardt", "mei"} {
				cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
			}

			switch mapids[i] {
			case "anubis":
				pos = "offense"
				for _, ch := range []string{"zenyatta", "lucio", "dva", "winston"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				pos = "defense"
				for _, ch := range []string{"zarya"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: "symmetra"}, Position: pos, Explain: "A 거점 한정"}
				break
			case "hanamura":
				pos = "offense"
				for _, ch := range []string{"tracer", "dva"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: "symmetra"}, Position: pos, Explain: "B 거점 한정"}
				pos = "defense"
				for _, ch := range []string{"roadhog", "hanzo"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				break
			case "volskaya":
				pos = "offense"
				for _, ch := range []string{"tracer", "reinhardt", "zarya", "lucio", "winston"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}

				pos = "defense"
				for _, ch := range []string{"hanzo", "soldier_76", "widowmaker", "zarya"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
			}
			break
		case i < 6: //Control

			///////////////COMMON
			// OFFENSE
			pos = "offense"
			for _, ch := range []string{"winston", "dva", "reaper", "lucio"} {
				cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
			}

			switch mapids[i] {
			case "ilios":
				for _, ch := range []string{"maccree", "pharah", "soldier_76", "roadhog", "genji", "hanzo"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				break
			case "nepal":
				for _, ch := range []string{"reinhardt", "roadhog", "junkrat", "mei"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				break
			case "lijiang":
				for _, ch := range []string{"reinhardt", "pharah", "soldier_76", "junkrat", "mei"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
			}
			break
		case i < 9: //Escort

			///////////////COMMON
			// OFFENSE
			pos = "offense"
			for _, ch := range []string{"dva", "genji", "winston", "reaper"} {
				cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
			}

			switch mapids[i] {
			case "dorado":
				pos = "offense"
				for _, ch := range []string{"lucio", "maccree"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				pos = "defense"
				for _, ch := range []string{"reinhardt", "roadhog", "junkrat", "pharah"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: "bastion"}, Position: pos, Explain: "경유지, 목적지 한정"}

				break
			case "route66":
				pos = "offense"
				for _, ch := range []string{"soldier_76", "reinhardt", "maccree", "pharah", "mercy"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				pos = "defense"
				for _, ch := range []string{"zarya", "maccree", "soldier_76", "symmetra", "mei", "junkrat", "mercy"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: "bastion"}, Position: pos, Explain: "목적지 한정"}
				cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: "torbjon"}, Position: pos, Explain: "목적지 한정"}
				break
			case "watchpoint_gibraltar":
				pos = "offense"
				for _, ch := range []string{"reinhardt", "soldier_76", "lucio"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				pos = "defense"
				for _, ch := range []string{"reinhardt", "roadhog", "zarya", "bastion", "widowmaker", "mercy"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
			}
			break
		case i < 12: //AssaultEscort

			///////////////COMMON
			// OFFENSE
			pos = "offense"
			for _, ch := range []string{"dva", "winston", "pharah"} {
				cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
			}
			// DEFENSE
			pos = "defense"
			for _, ch := range []string{"bastion", "torbjorn"} {
				cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
			}

			switch mapids[i] {
			case "hollywood":
				pos = "offense"
				for _, ch := range []string{"junkrat", "maccree", "reaper", "roadhog"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				pos = "defense"
				for _, ch := range []string{"maccree", "symmetra", "reaper", "pharah"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				break
			case "king's_row":
				pos = "offense"
				for _, ch := range []string{"junkrat", "genji", "tracer"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				pos = "defense"
				for _, ch := range []string{"widowmaker", "hanzo"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				break
			case "numbani":
				pos = "offense"
				for _, ch := range []string{"reaper", "soldier_76", "genji", "maccree", "tracer", "reinhardt"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				pos = "defense"
				for _, ch := range []string{"reinhardt", "widowmaker", "hanzo", "maccree", "junkrat", "mei"} {
					cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: ch}, Position: pos, Explain: ""}
				}
				cfm <- &CharsForMap{Map: &Map{Mapid: mapids[i]}, Char: &Character{Charid: "symmetra"}, Position: pos, Explain: "거점 한정"}
			}

		} //switch

	} //for

}
