package engine

import (
	"log"

	"github.com/liang24/go-crawler/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}
	counter := 0
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url, r.NewHttpRequestFunc)
		if err != nil {
			log.Printf("Fetcher: error "+"fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item #%d: %v\n", counter, item)
			counter++
		}
	}
}
