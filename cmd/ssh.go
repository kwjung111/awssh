package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ssh called")
		connectViaSystemSSH(args)
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)

}

func connectViaSystemSSH(args []string) {
	key := "your key file "
	host := "host"
	port := "port"
	cmd := exec.Command("bash", "-c", "ssh", "-i", key, host, "-p", port)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("ssh failed to start : %v", err)
	}

}
