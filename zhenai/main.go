package main

import (
	"github.com/liang24/go-crawler/engine"
	"github.com/liang24/go-crawler/persist"
	"github.com/liang24/go-crawler/scheduler"
	"github.com/liang24/go-crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.ConcurrentScheduler{},
		WorkerCount: 20,
		ItemSaver:   itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
