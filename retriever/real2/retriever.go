package real2

import (
	"net/http"
	"net/http/httputil"
	"time"
)

// 结构体中定义不同的变量
type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

/**
  只要实现了Get方法就可以
*/
func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
