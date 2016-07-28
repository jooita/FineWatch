package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

var mapids = []string{"anubis", "hanamura", "volskaya", "ilios", "nepal", "lijiang", "dorado", "route66", "watchpoint_gibraltar", "hollywood", "king's_row", "numbani"}

var charids = []string{"genji", "maccree", "pharah", "reaper", "soldier_76", "tracer", "bastion", "hanzo", "junkrat", "mei", "torbjorn", "widowmaker", "dva", "reinhardt", "roadhog", "winston", "zarya", "lucio", "mercy", "symmetra", "zenyatta", "ana"}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(Map))
	orm.RegisterModelWithPrefix("tb_", new(Character))
	orm.RegisterModelWithPrefix("tb_", new(MapList))
	orm.RegisterModelWithPrefix("tb_", new(CharsForMap))
}

func ErrHandle(id int64, err error) {
	if err == nil {
		fmt.Println("Insert :", id)
	}
}
