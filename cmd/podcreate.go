package cmd

import (
	"context"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	CoreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(podcreateCmd)
}

var podcreateCmd = &cobra.Command {
	Use: "podcreate",
	Short: "Create a pod in the K8S",
	Long: "Create the pod with containers in the current K8S",
	RunE: func(cmd *cobra.Command, args []string) error {
		var home string
		var kubeconfig string
		var err error

		pod := CoreV1.Pod {
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-pod",
			},
			Spec: CoreV1.PodSpec {
				Containers: []CoreV1.Container {
					{
						Name: "test-container",
						Image: "hello-world:latest",
					},
				},
			},
		}

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
		_, err = clientset.CoreV1().Pods("default").Create(ctx, &pod, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		return nil
	},
}

