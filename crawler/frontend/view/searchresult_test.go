package view

import (
	"os"
	"testing"

	"github.com/liang24/go-crawler/crawler/engine"
	"github.com/liang24/go-crawler/crawler/frontend/model"
	common "github.com/liang24/go-crawler/crawler/model"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template_test.html")

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "https://album.zhenai.com/u/1715998969",
		Type: "zhenai",
		Id:   "1715998969",
		Payload: common.Profile{
			Name:      "梦的解析",
			Gender:    "男",
			Age:       23,
			Height:    162,
			Weight:    70,
			Income:    "5001-8000元",
			Marriage:  "未婚",
			Education: "大学本科",
			Xinzuo:    "魔羯座",
			Car:       "未买车",
			House:     "打算婚后购房",
			Hokou:     "重庆",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
