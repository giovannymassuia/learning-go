/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("category called with name: %s\n", name)

	},

	// PreRun and PostRun are called before and after the Run function
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("pre run: category called")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("post run: category called")
	},

	// RunE is called when an error occurs
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("run error: category called")
	},
}

func init() {
	rootCmd.AddCommand(categoryCmd)

	categoryCmd.PersistentFlags().StringP("name", "n", "", "Category name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
