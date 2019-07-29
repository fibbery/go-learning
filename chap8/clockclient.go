package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, e := net.Dial("tcp", "localhost:9999")
	if e != nil {
		log.Fatal(e)
	}
	defer conn.Close()
	for {
		if _, e := io.Copy(os.Stdout, conn); e != nil {
			log.Fatal(e)
		}
	}
}
