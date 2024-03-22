package main

import (
	"aws"
	"main/cmd"
)

func main() {
	aws.InitClient()
	cmd.Execute()
}
