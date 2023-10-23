package main

import (
	"fmt"
	"strings"
)

type Response struct {
	code    int
	headers map[string]string
	body    string
}

func NewResponse(code int) *Response {
	return &Response{code: code, headers: make(map[string]string)}
}

// func (res *Response) addHeader(name string, value string) *Response {
// 	res.headers[name] = value
// 	return res
// }

func (res *Response) addTextBody(value string) *Response {
	res.body = value
	res.headers["Content-Type"] = "text/plain"
	res.headers["Content-Length"] = fmt.Sprint(len(value))
	return res
}

var codeToMsg = map[int]string{
	200: "OK",
	404: "Not Found",
}

func (res *Response) toString() string {
	lines := []string{
		fmt.Sprintf("HTTP/1.1 %d %s", res.code, codeToMsg[res.code]),
	}

	for key, value := range res.headers {
		lines = append(lines, fmt.Sprintf("%s: %s", key, value))
	}

	if res.body != "" {
		lines = append(lines, "", res.body)
	}

	return strings.Join(lines, "\r\n") + "\r\n\r\n"
}
