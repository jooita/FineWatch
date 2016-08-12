package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

var mapids = []string{"temple_of_anubis", "hanamura", "volskaya_industries", "ilios", "nepal", "lijiang_tower", "dorado", "route_66", "watchpoint_gibraltar", "hollywood", "king's_row", "numbani"}

var kmapids = []string{"아누비스 신전", "하나무라", "볼스카야 인더스트리", "일리오스", "네팔", "리장 타워", "도라도", "66번 국도", "감시기지: 지브롤터", "할리우드", "왕의 길", "눔바니"}

var charids = []string{"genji", "maccree", "pharah", "reaper", "soldier_76", "tracer", "bastion", "hanzo", "junkrat", "mei", "torbjorn", "widowmaker", "dva", "reinhardt", "roadhog", "winston", "zarya", "lucio", "mercy", "symmetra", "zenyatta", "ana"}

var kcharids = []string{"겐지", "맥크리", "파라", "리퍼", "솔저: 76", "트레이서", "바스티온", "한조", "정크랫", "메이", "토르비욘", "위도우메이커", "D.VA", "라인하르트", "로드호그", "윈스턴", "자리야", "루시우", "메르시", "시메트라", "젠야타", "아나"}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(Map))
	orm.RegisterModelWithPrefix("tb_", new(Character))
	orm.RegisterModelWithPrefix("tb_", new(CharsForMap))
}

func ErrHandle(id int64, err error) {
	if err == nil {
		fmt.Println("Insert :", id)
	}
}
