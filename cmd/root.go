/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	
	"github.com/nanobox-io/golang-scribble"
	"github.com/spf13/cobra"
)

type Heros struct{ Name string }

var (
	db     *scribble.Driver
	dbFile string 
	heros []string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "acr",
	Short: "AC Rebellion Planer",
	Long:  `Ein kleiner cli-Planer für AC Rebellion.`,
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
	rootCmd.PersistentFlags().StringVar(&dbFile, "database", dbFile, "database file")
	
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	
	cobra.OnInitialize(initDB)
}

func initDB() {
	var err error
	dbFile = os.ExpandEnv("$USERPROFILE\\.acr")

	// a new scribble driver, providing the directory where it will be writing to,
	// and a qualified logger if desired
	db, err = scribble.New(dbFile, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	// Read heros from the database
	heros, _ = db.ReadAll("heros")
}
