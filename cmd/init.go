/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/coderjojo/gogit/internal/repository"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new, empty repository",
	Long: `Initialize a new, emptry respository by default at current 
  working directory
  Specify the path to Initialize at different location 

  gogit init path`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path string

		if len(args) > 0 {
			path = args[0]
		} else {
			path, _ = os.Getwd()
		}

		_, err := repository.RepoCreate(path)
		if err != nil {
			fmt.Println("Error Creating Repository: ", err)
			return
		}

		fmt.Printf("Initializing Repository at %s\n", path)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
