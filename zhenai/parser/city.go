package parser

import (
	"regexp"
	"strings"

	"github.com/liang24/go-crawler/engine"
	"github.com/liang24/go-crawler/zhenai/factory"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	genderRe  = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
)

func ParseCity(contents []byte) engine.ParseResult {
	userMatches := profileRe.FindAllSubmatch(contents, -1)
	// genderMatches := genderRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for i, _ := range userMatches {
		user := userMatches[i]
		// gender := genderMatches[i]

		result.Items = append(result.Items, "User "+string(user[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:                strings.Replace(string(user[1]), "http:", "https:", 1),
			ParserFunc:         ParseProfile,
			NewHttpRequestFunc: factory.NewRequest,
		})
	}

	return result
}
