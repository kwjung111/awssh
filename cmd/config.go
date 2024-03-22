package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "set default awssh configuration",
	Long:  `set config : `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")

		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter your default profile : ")
		profile, _ := reader.ReadString('\n')
		profile = strings.TrimSpace(profile)

		viper.Set("profile", profile)

		err := viper.WriteConfigAs("config.yaml")
		if err != nil {
			fmt.Println("Error saving config:", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
