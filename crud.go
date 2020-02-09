/* package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
)

func main() {
	kubeconfig := flag.String("kubeconfig", filepath.Join("/etc", "rancher", "k3s", "k3s.yaml"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Printf("The kubeconfig file cannot be loaded %v \n", err)
		os.Exit(1)
	}
	clienset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Encountered error trying to create the kubernetes client session %v \n", err)
		os.Exit(1)
	}
}
