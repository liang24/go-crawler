package parser

import (
	"io/ioutil"
	"testing"

	"github.com/liang24/go-crawler/engine"
	"github.com/liang24/go-crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_text_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "男", "https://album.zhenai.com/u/1715998969")

	expected := engine.Item{
		Url:  "https://album.zhenai.com/u/1715998969",
		Type: "zhenai",
		Id:   "1715998969",
		Payload: model.Profile{
			Name:      "梦的解析",
			Gender:    "男",
			Age:       23,
			Height:    162,
			Weight:    0,
			Income:    "5001-8000元",
			Marriage:  "未婚",
			Education: "大学本科",
			Xinzuo:    "魔羯座",
			Car:       "未买车",
			House:     "打算婚后购房",
			Hokou:     "重庆",
		},
	}

	actual := result.Items[0]

	if expected != actual {
		t.Errorf("expected: %v; but was %v", expected, actual)
	}
}
