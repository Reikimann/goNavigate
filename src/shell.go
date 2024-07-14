package gonavigate

import (
	"fmt"
	"strings"
)

// import (
//   "os"
//   "text/template"
// )

// ShellName represents valid shell names
type ShellName string

const (
  Zsh  ShellName = "zsh"
  Bash ShellName = "bash"
  Fish ShellName = "fish"
)

func IsValidShell(value string) bool {
  switch ShellName(strings.ToLower(value)) {
  case Zsh:
    return true
  default:
    return false
  }
}

// // The init function passes the shell
// // Make a switch statement
// https://github.com/ajeetdsouza/zoxide/bloc/main/src/cmd/init.rs

func Render(shell string, cmd string) {
  // switch source := shell {

  // }
  fmt.Printf("shell: %s, cmd: %s", shell, cmd)
}
