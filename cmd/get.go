package cmd

import (
	"aws"
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get your EC2 infos with your profile",
	Long:  `Get your EC2 infos : name, private ip - with your profile`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
		aws.GetEc2List()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
