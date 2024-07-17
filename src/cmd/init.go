/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
  "fmt"

  goNav "github.com/Reikimann/goNavigate/src"

  "github.com/spf13/cobra"
)

// The cmdString for quickly navigating using the shell scripts
var cmdString string

var initCmd = &cobra.Command{
  Use:   "init [shellName]",
  Short: "Generate shell configuration",
  // Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
  // ValidArgs: []string{"zsh", "bash"},
  Args: func(cmd *cobra.Command, args []string) error {
    if len(args) != 1 {
      return fmt.Errorf("requires exactly one argument")
    }
    if _, valid := goNav.IsValidShell(args[0]); !valid {
      return fmt.Errorf("invalid shell name: %s. Allowed values are: zsh.", args[0])
    }
    return nil
  },
  Run: func(cmd *cobra.Command, args []string) {
    shell, _ := goNav.IsValidShell(args[0])

    opts := goNav.Opts{
      Cmd: cmdString,
    }

    goNav.RenderShellFuncs(shell, opts)
  },
}

func init() {
  rootCmd.AddCommand(initCmd)

  initCmd.Flags().StringVar(&cmdString, "cmd", "g", "Changes the prefix of the `g` command [default: g]")
}
