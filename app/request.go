package main

import "strings"

type Request struct {
	method  string
	path    string
	version string
	headers map[string]string
}

func ParseRequest(lines []string) *Request {
	startLine := strings.Split(lines[0], " ")

	headers := make(map[string]string)
	for _, line := range lines[1:] {
		if line != "" {
			split := strings.Split(line, ": ")
			headers[split[0]] = split[1]
		}
	}

	return &Request{
		method:  startLine[0],
		path:    startLine[1],
		version: startLine[2],
		headers: headers,
	}
}
