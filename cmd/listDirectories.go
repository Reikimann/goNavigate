/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
  "fmt"
  "log"
  "github.com/Reikimann/goNavigate/db"

  "github.com/spf13/cobra"
)

var listDirectoriesCmd = &cobra.Command{
  Use:   "listDirs",
  Short: "Prints the DB directories table to stdout",
  Run: func(cmd *cobra.Command, args []string) {
    d, err := db.OpenDatabase()
    if err != nil {
      log.Fatalf("Failed to open database: %v", err)
    }
    defer d.Database.Close()

    directories, err := d.ListDirectories()
    if err != nil {
      log.Fatalf("Failed to list directories: %v", err)
    }

    for _, d := range directories {
      fmt.Printf("\nID: %v\nPath: %s\nRecurse: %v\nLastNavigation: %v\nTimesNavigated: %v\n", d.ID,
                                                                                              d.Path,
                                                                                              d.Recurse,
                                                                                              d.LastNavigation,
                                                                                              d.TimesNavigated)
    }
  },
}

func init() {
  rootCmd.AddCommand(listDirectoriesCmd)
}
