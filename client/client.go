package client

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	utils "github.com/mainak90/kube-test-ext/utils"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig *string

func Client() *kubernetes.Clientset {
	if utils.FileExists("/etc/rancher/k3s/k3s.yaml") {
		kubeconfig = flag.String("kubeconfig", filepath.Join("/etc", "rancher", "k3s", "k3s.yaml"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("The kubeconfig file cannot be loaded %v \n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("There is an error creating the client instance with the kubeconfig file %v \n", err)
		os.Exit(1)
	}

	return clientset
}
