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
	Short: "make ssh connection with your defined items",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Println("wrong arguments")
		}
		fmt.Println("ssh called")
		connectViaSystemSSH(args)
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}

func IsThereMatchingItem(name string) {

}

func connectViaSystemSSH(args []string) {

	if len(args) < 1 {
		fmt.Println("item argument needed")
		return
	}

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
