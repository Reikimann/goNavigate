/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
  "fmt"
  "log"

  "github.com/reikimann/goNavigate/db"
  "github.com/spf13/cobra"
)

// Null value makes this default to false
var recurse bool

var addCmd = &cobra.Command{
  Use:   "add [dirPath(s)]",
  Short: "Adds a directory to the list of directories",
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    err := db.AddDirectories(args, recurse)
    if err != nil {
      log.Fatalf("Failed to add directories: %v", err)
    }
    fmt.Println("Directories added successfully.")
  },
}

func init() {
  rootCmd.AddCommand(addCmd)

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // addCmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
  addCmd.Flags().BoolVarP(&recurse, "recurse", "r", false, "Toggles if the added directory is recursable")
}
