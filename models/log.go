package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Log struct {
	Tagid   string
	Level   float64
	Avatar  string
	Qwins   string
	Qlost   float64
	Qplayed float64
}

func (this *Log) GetLog(tagid string, region string) *Log {
	tagid = strings.Replace(tagid, "#", "-", 1)
	escaped := url.QueryEscape(tagid)

	resp, err := http.Get("https://api.lootbox.eu/pc/" + region + "/" + escaped + "/profile")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	logData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(logData))

	var mapData interface{}

	if err := json.Unmarshal(logData, &mapData); err != nil {

		panic(err.Error())
	}

	data := mapData.(map[string]interface{})["data"].(map[string]interface{})

	games := data["games"].(map[string]interface{})
	g_quick := games["quick"].(map[string]interface{})
	g_competitive := games["competitive"].(map[string]interface{})

	playtime := data["playtime"].(map[string]interface{})

	competitive := data["competitive"].(map[string]interface{})

	var log *Log
	log = &Log{Tagid: data["username"].(string), Level: data["level"].(float64), Avatar: data["avatar"].(string), Qwins: g_quick["wins"].(string)}

	return log
}
