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
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a ec2 instance",
	Long:  `start a ec2 instance by providing the instance ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		pterm.Info.Printf("start instance  %v... \n", id)
		//fmt.Printf("start instance  %v ... ", id)
		utils.StartInstance(id)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	var instanceID string
	startCmd.Flags().StringVarP(&instanceID, "id", "i", "", "Instance ID")
	if err := startCmd.MarkFlagRequired("id"); err != nil {
		pterm.Error.Println("Instance ID is required")
		return
	}
}
