package parser

import (
	"regexp"
	"strings"

	"github.com/liang24/go-crawler/crawler/engine"
	"github.com/liang24/go-crawler/crawler/zhenai/factory"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	genderRe  = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
	cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)">下一页</a>`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	profileMatches := profileRe.FindAllSubmatch(contents, -1)
	genderMatches := genderRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for i, _ := range profileMatches {
		profile := profileMatches[i]
		gender := string(genderMatches[i][1])

		url := strings.Replace(string(profile[1]), "http:", "https:", 1)

		result.Requests = append(result.Requests, engine.Request{
			Url:                url,
			Parser:             NewProfileParser(gender),
			NewHttpRequestFunc: factory.NewRequest,
		})
	}

	matches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}

type ProfileParser struct {
	gender string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.gender)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ParseProfile", p.gender
}

func NewProfileParser(gender string) *ProfileParser {
	return &ProfileParser{
		gender: gender,
	}
}
