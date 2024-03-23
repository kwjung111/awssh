package cmd

import (
	"aws"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get your EC2 infos with your profile",
	Long:  `Get your EC2 infos : name, private ip - with your profile`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			aws.GetEc2List()
			return
		}

		if len(args) == 1 {
			aws.GetEc2List()
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
