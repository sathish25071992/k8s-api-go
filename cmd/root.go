package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "k8s",
	Short: "CLI to access K8s instance",
	Long: `A API Based CLI to access K8S instance`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
