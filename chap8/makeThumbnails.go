package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func ImageFile(filename string) (string, error) {
	fmt.Printf("open file[%s]\n", filename)
	return "", nil
}

func makeThumbnails(filenames []string) {
	for _, filename := range filenames {
		if _, err := ImageFile(filename); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, filename := range filenames {
		go ImageFile(filename)
	}
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var group sync.WaitGroup
	for filename := range filenames {
		group.Add(1)
		go func() {
			defer group.Done()
			thumb, e := ImageFile(filename)
			if e != nil {
				log.Println(e)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}()

		go func() {
			group.Wait()
			close(sizes)
		}()

		var total int64
		for size := range sizes {
			total += size
		}
		return total
	}
}
