package main

import (
	"flag"
	"log"
	"net/rpc"
	"strings"

	"github.com/liang24/go-crawler/crawler/duplicate"
	"github.com/liang24/go-crawler/crawler/engine"
	"github.com/liang24/go-crawler/crawler/scheduler"
	"github.com/liang24/go-crawler/crawler/zhenai/parser"
	"github.com/liang24/go-crawler/crawler_distributed/config"
	itemsaver "github.com/liang24/go-crawler/crawler_distributed/persist/client"
	"github.com/liang24/go-crawler/crawler_distributed/rpcsupport"
	worker "github.com/liang24/go-crawler/crawler_distributed/worker/client"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")

	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()

	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	duplicator := duplicate.NewRedisDuplicator("127.0.0.1:6379", "")

	e := engine.ConcurrentEngine{
		Scheduler:         &scheduler.ConcurrentScheduler{},
		WorkerCount:       100,
		ItemChan:          itemChan,
		RequestProcessor:  processor,
		RequestDuplicator: duplicator,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}
