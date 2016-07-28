package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Character struct {
	Charid  string `orm:"pk"`
	Class   string
	Explain string
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
	} else {
		fmt.Println("GetChar:", char.Charid, char.Class, char.Explain)
	}

	return &char
}

func (this *Character) InsertChar() {

	fmt.Println("Insert Character!")

	o := orm.NewOrm()
	o.Using("default")
	for i := range charids {
		chars := Character{Charid: charids[i]}
		switch {
		case i < 6:
			chars.Class = "offense"
			break
		case i < 12:
			chars.Class = "defense"
			break
		case i < 17:
			chars.Class = "tank"
			break
		case i < 22:
			chars.Class = "support"
		}
		id, err := o.Insert(&chars)
		if err == nil {
			fmt.Println("Char ID : ", id)
		}
	}
}
