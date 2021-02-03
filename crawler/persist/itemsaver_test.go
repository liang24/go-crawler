package persist

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/liang24/go-crawler/crawler/engine"
	"github.com/liang24/go-crawler/crawler/model"
	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "https://album.zhenai.com/u/1715998969",
		Type: "zhenai",
		Id:   "1715998969",
		Payload: model.Profile{
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

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = Save(client, "dating_profile", expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected: %v", actual, expected)
	}
}
