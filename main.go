/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
  "fmt"
  "os"

  "github.com/Reikimann/goNavigate/src/cmd"
  goNavigate "github.com/Reikimann/goNavigate/src"
)

func exit(code int, err error) {
  if code == goNavigate.ExitError && err != nil {
    fmt.Fprintln(os.Stderr, err.Error())
  }
  os.Exit(code)
}

func main() {
  cmd.Execute()
}
