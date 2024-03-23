/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"cfg"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Profile, keyFile, port uses default value if not defined
// name, tag ,host allows all value if not defined
type Item struct {
	Profile string
	Name    string
	Tag     string
	KeyFile string
	Port    int
}

// itemCmd represents the item command
var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "manage your item",
	Long:  `item includes profile,name,tag,host,keyFile,port`,
	Run: func(cmd *cobra.Command, args []string) {

		add, _ := cmd.Flags().GetBool("add")
		if add {
			if len(args) < 1 {
				fmt.Println("Provide a name for the item")
				return
			} else if len(args) > 1 {
				fmt.Println("item should added one by one")
				return
			}
			addItem(args[0])
			return
		}

		list, _ := cmd.Flags().GetBool("list")
		if list {
			getItemList()
		}

		for _, arg := range args {
			getItem(arg)
		}

	},
}

func init() {
	rootCmd.AddCommand(itemCmd)

	itemCmd.Flags().BoolP("add", "a", false, "Add a new item")
	itemCmd.Flags().BoolP("list", "l", false, "get list of all items")

}

func getItemList() {
	path := "."
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Failed to Read Directory")
	}

	for _, file := range files {
		if IsValidItemConf(file.Name()) {
			fmt.Println(file.Name())
		}
	}
}

func IsValidItemConf(name string) bool {

	vip := cfg.LoadItem(name)

	if err := vip.ReadInConfig(); err != nil {
		return false
	}

	if vip.IsSet("Profile") && vip.IsSet("Name") && vip.IsSet("Tag") && vip.IsSet("KeyFile") && vip.IsSet("port") {
		return true
	} else {
		return false
	}
}

func getItem(name string) Item {

	var item Item

	vip := cfg.LoadItem(name)

	if err := vip.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file , %f", err)
	}

	if err := vip.Unmarshal(&item); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	fmt.Printf("===Item : %v\n", name)
	fmt.Printf("Profile : %s\n", item.Profile)
	fmt.Printf("Name    : %s\n", item.Name)
	fmt.Printf("Tag     : %s\n", item.Tag)
	fmt.Printf("KeyFile : %s\n", item.KeyFile)
	fmt.Printf("Port    : %d\n", item.Port)

	return item
}

func addItem(name string) {
	reader := bufio.NewReader(os.Stdin)
	item := Item{}
	port := 0

	fmt.Print("Enter Profile: ")
	item.Profile, _ = reader.ReadString('\n')
	item.Profile = strings.TrimSpace(item.Profile)

	fmt.Print("Enter Name: ")
	item.Name, _ = reader.ReadString('\n')
	item.Name = strings.TrimSpace(item.Name)

	fmt.Print("Enter Tag: ")
	item.Tag, _ = reader.ReadString('\n')
	item.Tag = strings.TrimSpace(item.Tag)

	fmt.Print("Enter KeyFile: ")
	item.KeyFile, _ = reader.ReadString('\n')
	item.KeyFile = strings.TrimSpace(item.KeyFile)

	fmt.Print("Enter Port: ")
	portStr, _ := reader.ReadString('\n')
	portStr = strings.TrimSpace(portStr)

	if portStr != "" {
		portCnv, err := strconv.Atoi(portStr)
		if err != nil {
			fmt.Println("Error converting port to integer:", err)
			return
		}
		if portCnv < 0 || portCnv > 65535 {
			fmt.Println("port should between 0 and 65535")
		}
		port = portCnv
	}

	item.Port = port

	vip := viper.New()

	vip.Set("Profile", item.Profile)
	vip.Set("Name", item.Name)
	vip.Set("Tag", item.Tag)
	vip.Set("KeyFile", item.KeyFile)
	vip.Set("Port", item.Port)

	cfg.SaveItem(vip, name)
}
