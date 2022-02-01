/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

// defineCmd represents the define command
var defineCmd = &cobra.Command{
	Use:   "define",
	Short: "Not yet implemented",
	Long:  `Not yet implemented`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("define not yet implemented")
	},
}

// todo: implement
func init() {
	registryCmd.AddCommand(defineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// defineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// defineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
