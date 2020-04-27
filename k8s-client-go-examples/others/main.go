// Note: the example only works with the code within the same release/branch.
package others

import (
	"context"
	"flag"
	"fmt"

	"k8s.io/apimachinery/pkg/fields"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	"os"
	"path/filepath"
	"time"
)

func main() {

	var (
		clientset *kubernetes.Clientset
	)

	kubeconfig := flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	if clientset, err = kubernetes.NewForConfig(config); err != nil {
		panic(err.Error())
	}

	useInformer(clientset)
	//useNewListWatchFromClient(clientset)

}

func useNewListWatchFromClient(clientset *kubernetes.Clientset) {
	var (
		pod v1.Pod
	)

	podList, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(podList.Items))

	for _, pod = range podList.Items {
		fmt.Printf("Pod %d\n", pod.Status.Phase)
	}

	watchlist := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", v1.NamespaceDefault,
		fields.Everything())
	_, controller := cache.NewInformer(
		watchlist,
		&v1.Pod{},
		time.Second*0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Printf("add: %s \n", obj)
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Printf("delete: %s \n", obj)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				fmt.Printf("old: %s, new: %s \n", oldObj, newObj)
			},
		},
	)
	stop := make(chan struct{})
	go controller.Run(stop)
	for {
		time.Sleep(time.Second)
	}
}

/*
	Preferred way to use Informer instead of watcher

*/
func useInformer(clientset *kubernetes.Clientset) {
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	podInformer := informerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Printf("add: %s \n", obj)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Printf("delete: %s \n", obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Printf("old: %s, new: %s \n", oldObj, newObj)
		},
	})
	informerFactory.Start(wait.NeverStop)
	informerFactory.WaitForCacheSync(wait.NeverStop)
	pod, _ := podInformer.Lister().Pods(v1.NamespaceDefault).Get("kubia-hbcbk")
	fmt.Println(pod.Name)
}
