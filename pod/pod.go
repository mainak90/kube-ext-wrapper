package pod

import (
	"fmt"

	"github.com/mainak90/kube-test-ext/client"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListPod() {
	clientset := client.Client()
	namespace := "kube-system"
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range pods.Items {
		fmt.Println("Podname is ", pod.GetName(), "\n")
		fmt.Println("Node name is ", pod.Spec.NodeName, "\n")
	}
}
