package models

import (
	"fmt"
	"net/http"
)

type Log struct {
	Tagid  string `orm:"pk"`
	Region string
	Most   string
	KD     string
}

func (this *Log) Get(tagid string, region string) *Log {
	var log *Log

	req, err := http.NewRequest("GET", "https://api.lootbox.eu/pc/kr/%EB%B2%A0%EC%A7%80%EB%B0%80-3657/profile", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		println("ERROR:", err)
		// handle err
	}
	defer resp.Body.Close()

	fmt.Printf("%v\n", resp)

	return log
}
