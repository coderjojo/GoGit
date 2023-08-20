/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// # TODO: Add help
// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hashObject",
	Short: "Converts an file to a git object",
	Long: `Converrts and existing file to a git object

  ex: gogit `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hashObject called")
	},
}

func init() {
	rootCmd.AddCommand(hashObjectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashObjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashObjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
