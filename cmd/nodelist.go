package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nodelistCmd)
}

var nodelistCmd = &cobra.Command {
	Use: "nodelist",
	Short: "List the nodes in the K8S",
	Long: "List the running nodes in the current K8S",
	RunE: func(cmd *cobra.Command, args []string) error {
		var home string
		var kubeconfig string
		var err error

		home, err = os.UserHomeDir()
		if err != nil {
			return err
		}
		kubeconfig = filepath.Join(home, ".kube", "config")

		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
		}
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			return err
		}
		ctx, _ := context.WithCancel(context.Background())
		nodelist, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
		if err != nil {
			return err
		}
		for _, n := range nodelist.Items {
			fmt.Println(n.GetName())
		}
		return nil
	},
}

