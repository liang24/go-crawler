package parser

import (
	"context"
	"regexp"
	"strings"

	"github.com/liang24/go-crawler/engine"
	"github.com/liang24/go-crawler/zhenai/factory"
	"gopkg.in/olivere/elastic.v5"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	genderRe  = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
	cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)">下一页</a>`)
)

var client *elastic.Client

func init() {
	var err error
	client, err = elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
}

func ParseCity(contents []byte, _ string) engine.ParseResult {
	profileMatches := profileRe.FindAllSubmatch(contents, -1)
	genderMatches := genderRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for i, _ := range profileMatches {
		profile := profileMatches[i]
		gender := string(genderMatches[i][1])

		url := strings.Replace(string(profile[1]), "http:", "https:", 1)
		id := extractString([]byte(url), idRe)

		_, err := client.Get().
			Index("dating_profile").
			Type("zhenai").
			Id(id).
			Do(context.Background())
		if err == nil { //表示存在
			continue
		}

		result.Requests = append(result.Requests, engine.Request{
			Url:                url,
			ParserFunc:         ProfileParser(gender),
			NewHttpRequestFunc: factory.NewRequest,
		})
	}

	matches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}

func ProfileParser(gender string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, string(gender[1]), url)
	}
}
