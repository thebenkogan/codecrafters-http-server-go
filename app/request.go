package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Request struct {
	method  string
	path    []string
	version string
	headers map[string]string
	body    []byte
}

func ReadRequest(reader *bufio.Reader) *Request {
	res := &Request{headers: make(map[string]string)}

	for {
		line, _ := reader.ReadString(byte('\n'))
		line = strings.TrimSuffix(line, "\r\n")
		if res.method == "" {
			// reading first line
			startLine := strings.Split(line, " ")
			res.method = startLine[0]
			res.path = strings.Split(startLine[1], "/")[1:]
			res.version = startLine[2]
		} else if line != "" {
			// reading header
			split := strings.Split(line, ": ")
			res.headers[split[0]] = split[1]
		} else {
			// reading blank line, might be a body after this
			if lengthStr, ok := res.headers["Content-Length"]; ok {
				length, _ := strconv.Atoi(lengthStr)
				body := make([]byte, length)
				io.ReadFull(reader, body)
				res.body = body
			}
			break
		}
	}

	return res
}
