package main

import (
	"fmt"
	"log"
	"net/http"
)

type dolloars float32

func (d dolloars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dolloars

func (d database) list(response http.ResponseWriter, request *http.Request) {
	for item, price := range d {
		fmt.Fprintf(response, "%s : %s\n", item, price)
	}
}

func (d database) price(response http.ResponseWriter, request *http.Request) {
	item := request.URL.Query().Get("item")
	price, ok := d[item]
	if !ok {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(response, "no such item %q\n", item)
		return
	}
	fmt.Fprintf(response, "price : %s \n ", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5,}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("127.0.0.1:9080", nil))
}
