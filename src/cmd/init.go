/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
  "log"

  "github.com/Reikimann/goNavigate/src/db"
  "github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
  Use:   "init",
  Short: "Initializes the directory database",
  Run: func(cmd *cobra.Command, args []string) {
    d, err := db.OpenDatabase()
    if err != nil {
      log.Fatal(err)
    }
    defer d.Database.Close()

    d.InitDB()
  },
}

func init() {
  rootCmd.AddCommand(initCmd)
}
