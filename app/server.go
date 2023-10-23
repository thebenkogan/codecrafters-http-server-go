package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Panic("Failed to bind to port 4221")
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic("Error accepting connection: ", err.Error())
		}
		go handleRequest(conn)
	}
}

var codeToMsg = map[int]string{
	200: "OK",
	404: "Not Found",
}

func buildResponse(code int) string {
	fmt.Println(code, codeToMsg[code])
	return fmt.Sprintf("HTTP/1.1 %d %s\r\n\r\n", code, codeToMsg[code])
}

type Request struct {
	method  string
	path    string
	version string
	headers map[string]string
}

func strToRequest(input string) *Request {
	lines := strings.Split(input, "\r\n")
	startLine := strings.Split(lines[0], " ")

	headers := make(map[string]string)
	// fmt.Println(lines, len(lines), input, startLine)
	// for _, line := range lines[1:] {
	// 	if line != "" {
	// 		split := strings.Split(line, ": ")
	// 		headers[split[0]] = split[1]
	// 	}
	// }

	return &Request{
		method:  startLine[0],
		path:    startLine[1],
		version: startLine[2],
		headers: headers,
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	buf, err := io.ReadAll(conn)
	if err != nil {
		log.Panic("Failed to read", err)
	}
	fmt.Println("burh???")
	req := strToRequest(string(buf))
	fmt.Println(req)
	fmt.Println(req.path == "/")

	var response string
	if req.path == "/" {
		response = buildResponse(200)
	} else {
		response = buildResponse(404)
	}
	fmt.Println("here", response)

	if _, err := conn.Write([]byte(response)); err != nil {
		log.Panic("Failed to write", err)
	}
}
