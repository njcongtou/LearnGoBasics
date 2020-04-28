package mock

import "fmt"

type Retriever struct {
	Contents string
}

/**
  只要实现了Get方法就可以
*/
func (r Retriever) Get(url string) string {
	fmt.Println("input string: ", url)
	return r.Contents
}
