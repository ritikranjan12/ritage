/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get your age",
	Long: `We predict your age by your name.
To use please pass your first name.
An example is shown here :
	ritage get [firstname]
Note : please remove the square bracket.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("Thanks for using!")
		fmt.Println()
		time.Sleep(2 * time.Second)
		c := color.New(color.FgCyan)
		c.Println("Please Wait we are Calculating...")
		fmt.Println()
		name := args[0]
		url := `https://api.agify.io/?name=` + name
		res, _ := http.Get(url)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		x := map[string]int64{}
		json.Unmarshal([]byte(string(body)), &x)
		age := x["age"]
		red := color.New(color.FgRed, color.Bold)
		whiteBackground := red.Add(color.BgWhite)
		whiteBackground.Print("Your Age is : ", age, "  ")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
