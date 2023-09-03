/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	objectType string
	writeFlag  bool
)

// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hash-object [-w] [-t TYPE] FILE",
	Short: "Compute object ID optionally create a blob from a file",
	Long: `-t Specify the type TYPE "blob, commit, tag, tree" 
         -w write Actually write object into the database
         path Read object from <file>
  `,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		fmt.Printf("hashObject called on path %s", path)
	},
}

func init() {
	rootCmd.AddCommand(hashObjectCmd)

	hashObjectCmd.Flags().StringVarP(&objectType, "type", "t", "blob", `Specify the type
    (blob, commit, tag, tree)`)

	hashObjectCmd.Flags().BoolVarP(&writeFlag, "write", "w", false, `Actually write the 
    object into the database`)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashObjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashObjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
