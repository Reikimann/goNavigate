/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Reikimann/goNavigate/db"

	"github.com/spf13/cobra"
)

// Null value makes this default to false
var recurse bool

var addCmd = &cobra.Command{
  Use:   "add [dirPath(s)]",
  Short: "Adds a directory to the list of directories",
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    var paths []string

    // Gets the absolute file path from input
    for _, p := range args {
      p, err := filepath.Abs(p)
      if err != nil {
        log.Fatal(err)
      }

      // Checks if the file exists
      fileInfo, err := os.Stat(p)
      if err != nil {
        fmt.Printf("Path %s doesn't exist\n", p)
        continue
      }

      if !fileInfo.IsDir() {
        fmt.Printf("%s is not a directory\n", p)
        continue
      }

      paths = append(paths, p)
    }

    // If the filepath exists, then open and commit them to the database
    d, err := db.OpenDatabase()
    if err != nil {
      log.Fatalf("Failed to open database: %v", err)
    }
    defer d.Database.Close()

    // Commit filepaths to database
    err = d.AddDirectories(paths, recurse)
    if err != nil {
      log.Fatalf("Failed to add directories: %v", err)
    }
  },
}

func init() {
  rootCmd.AddCommand(addCmd)

  addCmd.Flags().BoolVarP(&recurse, "recurse", "r", false, "Toggles if the added directory is recursable")
}
