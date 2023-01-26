/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ritage",
	Short: "a CLI to calculate age",
	Long: `We calculate your age through your name.
We are Expanding so feel free to give your Feedback to ritik12343@gmail.com`,
	Run: func(cmd *cobra.Command, args []string) {
		c := color.New(color.FgHiYellow, color.Bold).Add(color.Underline)
		d := color.New(color.FgCyan, color.Bold)
		red := color.New(color.FgRed, color.Bold)
		fmt.Println()
		red.Println("Welcome to Ritage")
		fmt.Println()
		d.Println("* It is ", cmd.Root().Short)
		d.Print("* To use it run  -> ")
		c.Println("ritage get firstname")
		fmt.Println("")
		color.Blue(cmd.Root().Long)
		fmt.Println("")
		red.Println("Thank You")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ritage.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
