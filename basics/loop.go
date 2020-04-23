package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
  只有for
  没有while
*/

func main() {
	var (
		file    *os.File
		err     error
		scanner *bufio.Scanner
	)

	if file, err = os.Open("abc.txt"); err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(file)
	/**
	变成了一个while循环
	*/
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	/**
	死循环
	go routine 内部都是死循环，互相通信，很有意思
	*/
	for {
		fmt.Println("dead loop.")
	}

}
