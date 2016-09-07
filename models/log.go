package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Log struct {
	Data struct {
		Username string `json:"username"`
		Level    int    `json:"level"`
		Games    struct {
			Quick struct {
				Wins   string `json:"wins"`
				Lost   int    `json:"lost"`
				Played string `json:"played"`
			} `json:"quick"`
			Competitive struct {
				Wins   string `json:"wins"`
				Lost   int    `json:"lost"`
				Played string `json:"played"`
			} `json:"competitive"`
		} `json:"games"`
		Playtime struct {
			Quick       string `json:"quick"`
			Competitive string `json:"competitive"`
		} `json:"playtime"`
		Avatar      string `json:"avatar"`
		Competitive struct {
			Rank    string `json:"rank"`
			RankImg string `json:"rank_img"`
		} `json:"competitive"`
		LevelFrame string `json:"levelFrame"`
		Star       string `json:"star"`
	} `json:"data"`
}

func (this *Log) GetLog(tagid string, region string) *Log {
	tagid = strings.Replace(tagid, "#", "-", 1)
	escaped := url.QueryEscape(tagid)

	resp, err := http.Get("https://api.lootbox.eu/pc/" + region + "/" + escaped + "/profile")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var log *Log

	json.NewDecoder(resp.Body).Decode(&log)

	fmt.Println(log)

	return log
}
