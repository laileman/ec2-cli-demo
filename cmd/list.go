/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ec2-cli-demo/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all ec2 instances",
	Long:  `list all aws ec2 instances`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list all instances info")
		pterm.Info.Println("list all instances info")
		utils.ListInstances()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
