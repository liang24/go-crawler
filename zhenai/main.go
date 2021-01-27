package main

import (
	"github.com/liang24/go-crawler/engine"
	"github.com/liang24/go-crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
