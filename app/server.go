package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Panic("Failed to bind to port 4221")
	}

	conn, err := l.Accept()
	if err != nil {
		log.Panic("Error accepting connection: ", err.Error())
	}

	buf := make([]byte, 0)
	conn.Read(buf)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}
