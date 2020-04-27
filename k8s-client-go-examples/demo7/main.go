package main

import (
	"context"
	"fmt"

	"congwang.com/learnGo/k8s-client-go-examples/common"
	core_v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	var (
		clientset *kubernetes.Clientset
		tailLines int64
		req       *rest.Request
		res       rest.Result
		logs      []byte
		err       error
	)

	// init k8s client
	if clientset, err = common.InitClient(); err != nil {
		goto FAIL
	}

	// create a POD log request
	req = clientset.CoreV1().Pods("default").GetLogs("kubia-hbcbk", &core_v1.PodLogOptions{Container: "kubia", TailLines: &tailLines})

	// req.Stream() is same as Do() call

	// send request
	if res = req.Do(context.TODO()); res.Error() != nil {
		err = res.Error()
		goto FAIL
	}

	// get results
	if logs, err = res.Raw(); err != nil {
		goto FAIL
	}

	fmt.Println("Container Output: ", string(logs))
	return

FAIL:
	fmt.Println(err)
	return
}
