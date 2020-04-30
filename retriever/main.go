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

接口变量里面有什么
1. 接口变量自带指针
2. 接口变量同样采用值传递，几乎不需要使用接口的指针
3. 指针接受者实现智能以指针方式使用，值接受者都可

*/
func main() {
	var (
		r Retriever
	)

	/**
	调用接口的不同的实现
	*/
	r = mock.Retriever{"fake content."}
	fmt.Println(download(r))
	r = real2.Retriever{}
	fmt.Println(download(r))

	// type assertion 判断
	if r3, ok := r.(real2.Retriever); ok {
		fmt.Println(r3.UserAgent)
	} else {
		fmt.Println("r2 is NOT a real2 retriever")
	}

	// switch 方式
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("agent:", v.UserAgent)
	}

}
