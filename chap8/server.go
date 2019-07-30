package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	server, e := net.Listen("tcp", "localhost:9090")
	if e != nil {
		log.Fatal(e.Error())
	}
	// handle channel
	go broadcast()
	log.Printf("server[%s] start !!!\n", server.Addr().String())
	for {
		conn, e := server.Accept()
		if e != nil {
			log.Print(e)
			continue
		}
		go handleConnect(conn)
	}

}

func handleConnect(conn net.Conn) {
	ch := make(chan string)
	// deal message to this client
	go clientWriter(conn, ch)

	// send login message
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + " : " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
