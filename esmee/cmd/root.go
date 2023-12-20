/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"esmee/esmee/esmee"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "esmee",
	Short: "EsmeeDB",
	Long: `
Welcome to EsmeeDB CLI

This is a short documentation of the commands available. Refer to the online docs [TODO] for more details	

$ create [tree/branch]
The create commmand is used to create a new tree or branch

$ create tree
We can use this command to create a tree - $ create tree Food

$ climb [name of tree]
This is used to enter/climb a tree, e.g $ climb Vegetarian.
Once we climb up a tree, the prompt is updated with the name of the tree. And only then we can 
create a branch. e.g $ Food#

$ create branch:
To create a branch you must be in a tree. In order to create a branch "Vegetarian" on the tree
Inside the tree, we do # create branch Vegetarian
	
>Food#create branch Vegetarians
`,

	Run: func(cmd *cobra.Command, args []string) {
		esmee.StartEsmee()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
