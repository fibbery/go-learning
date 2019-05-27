package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer,"url.path = %s", request.URL.Path)
}
