/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/laileman/ec2-cli-demo/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop a ec2 instance",
	Long:  `stop a ec2 instance by providing the instance ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		instanceID, _ := cmd.Flags().GetString("id")
		pterm.Info.Printf("stop instance  %v... \n", instanceID)
		utils.StopInstance(instanceID)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
	var instanceID string
	stopCmd.Flags().StringVarP(&instanceID, "id", "i", "", "Instance ID")
	if err := stopCmd.MarkFlagRequired("id"); err != nil {
		pterm.Error.Println("Instance ID is required")
		return
	}
}
