package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"

	"congwang.com/learnGo/k8s-client-go-examples/common"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientset *kubernetes.Clientset
		err       error
	)

	// init k8s client
	if clientset, err = common.InitClient(); err != nil {
		fmt.Println(err)
		return
	}

	// watch for pod changes.
	/**
	Observation:
	when scaling in/out (delete or add a pod), AddFunc or DeleteFunc is only called once associated with 3 calls to UpdateFunc.
	*/
	watchlist := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", v1.NamespaceDefault, fields.Everything())
	_, controller := cache.NewInformer(
		watchlist,
		&v1.Pod{},
		time.Second*0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Printf("add: %s \n", obj.(*v1.Pod).Name)
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Printf("deleted: %s \n", obj.(*v1.Pod).Name)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				fmt.Printf("old: %s, new: %s \n", oldObj.(*v1.Pod).Name, newObj.(*v1.Pod).Name)
			},
		},
	)
	stop := make(chan struct{})
	go controller.Run(stop)
	for {
		time.Sleep(time.Second)
	}

}
