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
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "create ssh alias for socket5 proxy",
	Long:  `create ssh alias for socket5 proxy`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		user, _ := cmd.Flags().GetString("user")
		pterm.Info.Println("create ssh proxy alias")
		//fmt.Printf("start instance  %v ... ", id)
		utils.Ssh(file, user)
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
	var file string
	var user string
	sshCmd.Flags().StringVarP(&file, "file", "f", "", "ssh config file")
	sshCmd.Flags().StringVarP(&user, "user", "u", "", "ssh user")
	if err := sshCmd.MarkFlagRequired("file"); err != nil {
		pterm.Error.Println("SSH config file is required")
		return
	}
	if err := sshCmd.MarkFlagRequired("user"); err != nil {
		pterm.Error.Println("SSH user is required")
		return
	}
}
