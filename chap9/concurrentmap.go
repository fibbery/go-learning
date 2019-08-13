package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	response, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

type Memo struct {
	f     Func
	cache map[string]result
}

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	e     error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result),}
}

func (m *Memo) get(url string) (interface{}, error) {
	res, ok := m.cache[url]
	if !ok {
		res.value, res.e = m.f(url)
		m.cache[url] = res
	}
	return res.value, res.e
}

func main() {
	urls := [...]string{"http://www.baidu.com", "http://www.zhihu.com", "http://www.zhihu.com"}
	memo := New(httpGetBody)
	for _, url := range urls {
		start := time.Now()
		value, err := memo.get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}

	ch := make(chan int)
	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		close(ch)
		<-ch
		group.Done()
		fmt.Printf("hello,go")
	}()
	group.Wait()
}
