package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type ServerAddress struct {
	ip   string
	port int
}

func (s *ServerAddress) String() string {
	return s.ip + ":" + strconv.Itoa(s.port)
}

func (s *ServerAddress) Set(str string) error {
	data := strings.Split(str, ":")
	s.ip = data[0]
	s.port, _ = strconv.Atoi(data[1])
	return nil
}

func ServerAddressFlag(name string, ip string, port int, usage string) *ServerAddress {
	tmp := ServerAddress{ip, port,}
	flag.Var(&tmp, name, usage)
	return &tmp
}

var server = ServerAddressFlag("server","127.0.0.1", 9090,"enter server address")

func main() {
	flag.Parse()
	fmt.Print(server)
}
