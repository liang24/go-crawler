module github.com/liang24/go-crawler

go 1.15

require (
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	golang.org/x/text v0.3.5
	gopkg.in/olivere/elastic.v5 v5.0.86
)

replace (
	github.com/liang24/go-crawler/engine => ./engine
	github.com/liang24/go-crawler/fetcher => ./fetcher
	github.com/liang24/go-crawler/model => ./model
	github.com/liang24/go-crawler/persist => ./persist
	github.com/liang24/go-crawler/zhenai/factory => ./zhenai/factory
	github.com/liang24/go-crawler/zhenai/parser => ./zhenai/parser
)
