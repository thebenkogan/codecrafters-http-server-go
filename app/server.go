package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

const port int = 4221

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Panic("Failed to bind to port", port)
	}
	defer l.Close()

	fmt.Println("Listening on port", port)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic("Error accepting connection: ", err.Error())
		}
		go handleRequest(conn)
	}
}

type Request struct {
	method  string
	path    string
	version string
	headers map[string]string
}

func parseRequest(lines []string) *Request {
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

func handleRequest(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connection received")

	s := bufio.NewScanner(conn)
	s.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for s.Scan() {
		if text := s.Text(); text != "" {
			lines = append(lines, text)
		} else {
			break
		}
	}

	fmt.Println("Request:", strings.Join(lines, ", "))

	req := parseRequest(lines)
	pathParts := strings.Split(req.path, "/")[1:]

	var response *Response
	switch pathParts[0] {
	case "":
		response = NewResponse(200)
	case "echo":
		response = NewResponse(200).addTextBody(strings.Join(pathParts[1:], "/"))
	case "user-agent":
		response = NewResponse(200).addTextBody(req.headers["User-Agent"])
	default:
		response = NewResponse(404)
	}

	resStr := response.toString()
	fmt.Println("Responding:", resStr)
	conn.Write([]byte(resStr))
}
