package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_text_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCity(contents)

	const resultSize = 20
	expectedUrls := []string{
		"http://album.zhenai.com/u/1974534161", "http://album.zhenai.com/u/1197317604", "http://album.zhenai.com/u/1693313153",
	}
	expectedUsers := []string{
		"User 心悦", "User 折翼", "User 醉思恋人",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+
			"items; but had %d", resultSize, len(result.Items))
	}
	for i, city := range expectedUsers {
		if result.Items[i].(string) != city {
			t.Errorf("expected user #%d: %s; but was %s", i, city, result.Items[i].(string))
		}
	}
}
