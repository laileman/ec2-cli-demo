package main

import (
	"github.com/laileman/ec2-cli-demo/cmd"
)

func main() {
	// The main function is the entry point of the application.
	// It will execute the root command which includes all subcommands.
	cmd.Execute()
}
