/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
  "github.com/reikimann/goNavigate/db"
  "github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
  Use:   "init",
  Short: "Initializes the directory database",
  Run: func(cmd *cobra.Command, args []string) {
    db.InitDB()
  },
}

func init() {
  rootCmd.AddCommand(initCmd)
}
