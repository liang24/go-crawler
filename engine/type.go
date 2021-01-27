package engine

import (
	"io"
	"net/http"
)

type Request struct {
	Url                string
	ParserFunc         func([]byte) ParseResult
	NewHttpRequestFunc func(string, string, io.Reader) (*http.Request, error)
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
