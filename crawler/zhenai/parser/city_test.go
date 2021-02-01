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

	result := ParseCity(contents, "http://www.zhenai.com/zhenghun/aba")

	const resultSize = 21
	expectedUrls := []string{
		"https://album.zhenai.com/u/1974534161", "https://album.zhenai.com/u/1197317604", "https://album.zhenai.com/u/1693313153",
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
}
