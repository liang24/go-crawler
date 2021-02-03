package main

import (
	"github.com/liang24/go-crawler/crawler/engine"
	"github.com/liang24/go-crawler/crawler/persist"
	"github.com/liang24/go-crawler/crawler/scheduler"
	"github.com/liang24/go-crawler/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.ConcurrentScheduler{},
		WorkerCount:      20,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}
