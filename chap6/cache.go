package main

import (
	"fmt"
	"sync"
)

var (
	cache = struct {
		sync.Mutex
		mapping map[string]string
	}{
		mapping: make(map[string]string),
	}
)

func Lookeup(key string) string {
	cache.Lock()
	defer cache.Unlock()
	return cache.mapping[key]
}

func Add(key,value string) {
	cache.Lock()
	defer cache.Unlock()
	cache.mapping[key] = value
}

func main() {
	Add("jiangnenghua", "615542")
	fmt.Println(Lookeup("jiangnenghua"))
}
