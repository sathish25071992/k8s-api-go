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
	rootCmd.AddCommand(nslistCmd)
}

var nslistCmd = &cobra.Command {
	Use: "nslist",
	Short: "List the namespace in the K8S",
	Long: "List the running namespace in the current K8S",
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
		nslist, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
		if err != nil {
			return err
		}
		for _, n := range nslist.Items {
			fmt.Println(n.GetName())
		}
		return nil
	},
}

