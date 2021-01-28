package main

import (
	"github.com/liang24/go-crawler/engine"
	"github.com/liang24/go-crawler/scheduler"
	"github.com/liang24/go-crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.ConcurrentScheduler{},
		WorkerCount: 20,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
