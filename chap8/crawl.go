package main

import (
	"fmt"
	"os"
)

func crawl(url string) []string {
	fmt.Printf("正在爬%s的数据\n", url)
	return nil
}

func main() {
	worklist := make(chan []string)
	unseelinks := make(chan string)

	go func() {
		worklist <- os.Args[1:]
	}()

	for i := 0; i < 20; i++{
		go func() {
			for link := range unseelinks {
				foundurls := crawl(link)
				worklist <- foundurls
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _,url := range list {
			if !seen[url] {
				seen[url] = true
				unseelinks <- url
			}
		}
	}
}
