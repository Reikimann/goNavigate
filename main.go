/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
  "log"

  "github.com/reikimann/goNavigate/cmd"
  "github.com/reikimann/goNavigate/db"
)

func main() {
  if err := db.OpenDatabase(); err != nil {
    log.Fatalf("Failed to open database: %v", err)
  }
  cmd.Execute()
}
