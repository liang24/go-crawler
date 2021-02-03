package controller

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/liang24/go-crawler/crawler/engine"
	"github.com/liang24/go-crawler/crawler/frontend/model"
	"github.com/liang24/go-crawler/crawler/frontend/view"
	"gopkg.in/olivere/elastic.v5"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		// DEBUG时使用
		// elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		// elastic.SetTraceLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

// localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult

	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	fmt.Printf("q=%s\n", rewriteQueryString(q))

	result.Query = q
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))

	//支持分页
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Za-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
