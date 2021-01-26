package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getNodeList() ([]string, error) {
	var home string
	var kubeconfig string
	var err error
	var nodes []string

	home, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	kubeconfig = filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nodes, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nodes, err
	}
	ctx, _ := context.WithCancel(context.Background())
	nodelist, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nodes, err
	}
	for _, n := range nodelist.Items {
		nodes = append(nodes, n.GetName())
	}
	return nodes, nil
}

func main() {
	nodes, err := getNodeList()
	if err != nil {
		panic(err)
	}
	fmt.Println(nodes)
}
