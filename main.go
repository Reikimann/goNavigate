/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
  "log"
  "fmt"

  "github.com/Reikimann/goNavigate/cmd"
  "github.com/Reikimann/goNavigate/db"
  // tea "github.com/charmbracelet/bubbletea"
)

func main() {
  if err := db.OpenDatabase(); err != nil {
    log.Fatalf("Failed to open database: %v", err)
  }

  cmd.Execute()
  directories, err := db.ListDirectories()
  if err != nil {
    log.Fatalf("Failed to list directories: %v", err)
  }

  for _, d := range directories {
    fmt.Printf("\nPath: %s\nRecurse: %v\nLastNavigation: %v\nTimesNavigated: %v\n\n", d.Path,
                                                                                      d.Recurse,
                                                                                      d.LastNavigation,
                                                                                      d.TimesNavigated)
  }
  // os.Chdir("/home/reikimann/dox")
}
