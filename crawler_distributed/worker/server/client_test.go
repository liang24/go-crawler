package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/liang24/go-crawler/crawler_distributed/config"
	"github.com/liang24/go-crawler/crawler_distributed/rpcsupport"
	"github.com/liang24/go-crawler/crawler_distributed/worker"
)

func TestCrawlService(t *testing.T) {

	// Call rpc service
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "https://album.zhenai.com/u/1715998969",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "男",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
