/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/coderjojo/gogit/internal/repository"
	"github.com/spf13/cobra"
)

// catFileCmd represents the catFile command
var catFileCmd = &cobra.Command{
	Use:   "cat-file TYPE OBJECT",
	Short: "cat the object to the standard output",
	Long:  `Prints out an existing git object to the standard output `,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		objectType := args[0]
		object := args[1]
		repo, err := repository.RepoFind(".", true)

		if err != nil {
			fmt.Errorf("Error Occurred %s", err)
		}

		repository.CatFile(repo, object, objectType)

		fmt.Printf("objecttype : %s , object : %s", objectType, object)
	},
}

func init() {
	rootCmd.AddCommand(catFileCmd)
	catFileCmd.ValidArgs = []string{"blob", "commit", "tag", "tree"}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
