/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nanobox-io/golang-scribble"
	"github.com/spf13/cobra"
)

type Hero struct{ Name string }
type Item struct{ 
	Name string
	Material map[string]int
}

var (
	db     *scribble.Driver
	dbFile string
	heros  []Hero
	items  []Item
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "acr",
	Short: "AC Rebellion Planer",
	Long:  `Ein kleiner cli-Planer für AC Rebellion.`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initDB()
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
	dbFile = os.ExpandEnv("$PWD/.acr")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&dbFile, "database", dbFile, "database file")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initDB() {
	var err error

	// a new scribble driver, providing the directory where it will be writing to,
	// and a qualified logger if desired
	db, err = scribble.New(dbFile, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	// Read heros from the database
	records, _ := db.ReadAll("heros")
	for _, f := range records {
		hero := Hero{}
		if err := json.Unmarshal([]byte(f), &hero); err != nil {
			fmt.Println("Error", err)
		}
		heros = append(heros, hero)
	}

	// Read items from the database
	records, _ = db.ReadAll("items")
	for _, f := range records {
		item := Item{}
		if err := json.Unmarshal([]byte(f), &item); err != nil {
			fmt.Println("Error", err)
		}
		items = append(items, item)
	}
}
