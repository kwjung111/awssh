package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "awssh : a tool to connect EC2 with ssh",
	Long:  `awssh : a tool to connect EC2 with ssh`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executed Root command")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
