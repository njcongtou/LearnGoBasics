package main

import (
	"fmt"

	"congwang.com/learnGo/retriever/mock"
	"congwang.com/learnGo/retriever/real2"
)

type Retriever interface {
	// 不用加func
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.google.com")
}

/**
接口的用法
1. 定义接口interface 和 他的func
2. 定义结构体，要实现这个func
3. 创建这个结构体，并且调用这个func
*/
func main() {
	var r Retriever

	/**
	调用接口的不同的实现
	*/
	r = mock.Retriever{"fake content."}
	r = real2.Retriever{}
	fmt.Println(download(r))

}
