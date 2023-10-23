package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const port int = 4221

type Server struct {
	fileDir string
}

func main() {
	server := Server{}

	args := os.Args
	if len(args) > 1 && args[1] == "--directory" {
		server.fileDir = args[2]
		fmt.Println("Using directory", server.fileDir)
	}

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
		go server.handleRequest(conn)
	}
}

func (s *Server) handleRequest(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connection received")

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		if text := scanner.Text(); text != "" {
			lines = append(lines, text)
		} else {
			break
		}
	}

	fmt.Println("Request:", strings.Join(lines, ", "))

	req := ParseRequest(lines)
	pathParts := strings.Split(req.path, "/")[1:]

	var response *Response
	switch pathParts[0] {
	case "":
		response = NewResponse(200)
	case "echo":
		response = NewResponse(200).addTextBody(strings.Join(pathParts[1:], "/"))
	case "user-agent":
		response = NewResponse(200).addTextBody(req.headers["User-Agent"])
	case "files":
		fileName := pathParts[1]
		file, err := os.Open(s.fileDir + "/" + fileName)
		if os.IsNotExist(err) {
			response = NewResponse(404)
		} else {
			defer file.Close()
			buffer, _ := io.ReadAll(file)
			response = NewResponse(200).attachFile(buffer)
		}
	default:
		response = NewResponse(404)
	}

	resStr := response.toString()
	fmt.Println("Responding:", resStr)
	conn.Write([]byte(resStr))
}
