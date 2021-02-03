package main

import (
	"testing"
	"time"

	"github.com/liang24/go-crawler/crawler/engine"
	"github.com/liang24/go-crawler/crawler/model"
	"github.com/liang24/go-crawler/crawler_distributed/rpcsupport"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
		Url:  "https://album.zhenai.com/u/1715998969",
		Type: "zhenai",
		Id:   "1715998969",
		Payload: model.Profile{
			Name:      "梦的解析",
			Gender:    "男",
			Age:       23,
			Height:    162,
			Weight:    0,
			Income:    "5001-8000元",
			Marriage:  "未婚",
			Education: "大学本科",
			Xinzuo:    "魔羯座",
			Car:       "未买车",
			House:     "打算婚后购房",
			Hokou:     "重庆",
		},
	}

	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
