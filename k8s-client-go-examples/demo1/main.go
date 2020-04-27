package main

import (
	"context"
	"fmt"

	"congwang.com/learnGo/k8s-client-go-examples/common"
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientSet *kubernetes.Clientset
		podsList  *core_v1.PodList
		pod       v1.Pod
		err       error
	)

	// init k8s client
	if clientSet, err = common.InitClient(); err != nil {
		goto FAIL
	}

	if podsList, err = clientSet.CoreV1().Pods("default").List(context.TODO(), meta_v1.ListOptions{}); err != nil {
		goto FAIL
	}

	for _, pod = range podsList.Items {
		fmt.Printf("Pod %d\n", pod.Name)
	}

FAIL:
	fmt.Println(err)
	return
}
