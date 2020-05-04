package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 4; i++ {
		wg.Add(1) // add a counter
		go download("a.com/" + string(i+'0'))
	}
	wg.Wait() // wait for all goroutines to finish.
	fmt.Println("All Done!")
}

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	wg.Done()
}
