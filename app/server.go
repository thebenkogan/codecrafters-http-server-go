package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port int = 4221

type State struct {
	fileDir string
}

func main() {
	state := State{}

	args := os.Args
	if len(args) > 1 && args[1] == "--directory" {
		state.fileDir = args[2]
		fmt.Println("Using directory", state.fileDir)
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
		go handleRequest(&state, conn)
	}
}

var router = map[string]map[string](func(state *State, req *Request) *Response){
	"": {
		"GET": Home,
	},
	"echo": {
		"GET": Echo,
	},
	"user-agent": {
		"GET": UserAgent,
	},
	"files": {
		"GET":  GetFile,
		"POST": PostFile,
	},
}

func handleRequest(state *State, conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connection received")

	reader := bufio.NewReader(conn)
	req := ReadRequest(reader)

	fmt.Println("Request:", req)

	handler, ok := router[req.path[0]][req.method]
	var response *Response
	if ok {
		response = handler(state, req)
	} else {
		response = NewResponse(404)
	}

	resStr := response.toString()
	fmt.Println("Responding:", resStr)
	conn.Write([]byte(resStr))
}
