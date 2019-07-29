package main

import (
	"log"
	"net"
	"time"
)

func main() {
	server, e := net.Listen("tcp", "localhost:9999")
	if e != nil {
		log.Fatal(e)
	}
	log.Printf("this server listening at %s\n", server.Addr().String())
	for {
		conn, e := server.Accept()
		if e != nil {
			log.Fatal(e)
		}
		handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func() {
		log.Printf("client[%s] close connection", conn.RemoteAddr())
		conn.Close()
	}()
	for {
		bytes := []byte(time.Now().Format("15:04:05\n"))
		_, err := conn.Write(bytes)
		if err != nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}
