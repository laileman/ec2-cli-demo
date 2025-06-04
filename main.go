package main

import (
	"aws-proxy/cmd"
)

func main() {
	// The main function is the entry point of the application.
	// It will execute the root command which includes all subcommands.
	cmd.Execute()
}
