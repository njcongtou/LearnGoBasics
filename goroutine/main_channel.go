package main

import (
	"fmt"
	"time"
)

var ch = make(chan string, 10) // create a channel with buffer 10

func main() {
	for i := 0; i < 4; i++ {
		go download2("a.com/" + string(i+'0'))
	}

	for i := 0; i < 4; i++ {
		msg := <-ch // wait for msg from this channel
		fmt.Println("finish", msg)
	}
	fmt.Println("All Done!")
}

func download2(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url
}
