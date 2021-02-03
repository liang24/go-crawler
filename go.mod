module github.com/liang24/go-crawler

go 1.15

require (
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	golang.org/x/text v0.3.5
	gopkg.in/olivere/elastic.v5 v5.0.86
)

replace (
	github.com/liang24/go-crawler/crawler/engine => ./crawler/engine
	github.com/liang24/go-crawler/crawler/fetcher => ./crawler/fetcher
	github.com/liang24/go-crawler/crawler/model => ./crawler/model
	github.com/liang24/go-crawler/crawler/persist => ./crawler/persist
	github.com/liang24/go-crawler/crawler/zhenai/factory => ./crawler/zhenai/factory
	github.com/liang24/go-crawler/crawler/zhenai/parser => ./crawler/zhenai/parser
)
