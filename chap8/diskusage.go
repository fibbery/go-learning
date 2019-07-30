package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")
var directory = flag.String("f", ".", "the directory that be analyze memory usage")
var semaphore = make(chan struct{}, 20)
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func walkDir(dir string, filesizes chan<- int64, group *sync.WaitGroup) {
	defer group.Done()
	if cancelled() {
		return
	}
	for _, children := range directoryEntries(dir) {
		if children.IsDir() {
			group.Add(1)
			absolutepath := filepath.Join(dir, children.Name())
			go walkDir(absolutepath, filesizes, group)
		} else {
			filesizes <- children.Size()
		}
	}

}

func directoryEntries(dir string) []os.FileInfo {
	defer func() {<-semaphore}()
	select {
	case semaphore <- struct{}{}:
	case <-done:
		return nil
	}
	entries, e := ioutil.ReadDir(dir)
	if e != nil {
		return nil
	}
	return entries
}

func main() {
	defer logtime()()
	flag.Parse()
	fileSizes := make(chan int64)
	var group sync.WaitGroup
	group.Add(1)
	go walkDir(*directory, fileSizes, &group)
	go func() {
		group.Wait()
		close(fileSizes)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// verbose
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case <- done:
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func logtime() func() {
	start := time.Now()
	fmt.Printf("method start\n")
	return func() {
		fmt.Printf("\nmethod cost time : %s", time.Since(start).String())
	}
}

func printDiskUsage(nfiles int64, nbytes int64) {
	fmt.Printf("\r%d files, %.4f GB", nfiles, float64(nbytes)/(1000*1024*1024))
}
