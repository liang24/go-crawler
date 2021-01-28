package engine

import (
	"io"
	"net/http"
)

type Request struct {
	Url                string
	ParserFunc         ParserFunc
	NewHttpRequestFunc func(string, string, io.Reader) (*http.Request, error)
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

type ParserFunc func(contents []byte, url string) ParseResult

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
