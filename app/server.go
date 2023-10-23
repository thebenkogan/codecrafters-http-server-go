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
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic("Error accepting connection: ", err.Error())
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	if _, err := conn.Read(buf); err != nil {
		log.Panic("Failed to read", err)
	}

	if _, err := conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n")); err != nil {
		log.Panic("Failed to write", err)
	}
}
