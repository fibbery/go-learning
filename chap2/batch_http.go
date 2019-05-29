package main

import (
	"fmt"
	"net/http"
)

func main() {
	ch := make(chan int)
	for i := 1; i < 10; i++ {
		go httpget("https://www.zhenai.com", ch)
	}
	for i := 1; i < 10; i++ {
		fmt.Println(<-ch)
	}

}

func httpget(url string, ch chan<- int) {
	req, _ := http.NewRequest("get", url, nil)
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	client := &http.Client{}
	res , err:= client.Do(req)
	if err != nil {
		fmt.Printf("get url error, %v\n", err)
		ch <- 1
		return
	}
	ch <- res.StatusCode
}
