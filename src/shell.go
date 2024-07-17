package gonavigate

import (
  "fmt"
  "os"
  "strings"
  "embed"
  "text/template"
)

//go:embed templates/*.txt
var templateFS embed.FS

// ShellName represents valid shell names
type ShellName string

// Supported shells
const (
  Zsh  ShellName = "zsh"
  // Bash ShellName = "bash"
  // Fish ShellName = "fish"
)

func IsValidShell(value string) (ShellName, bool) {
  switch ShellName(strings.ToLower(value)) {
  // case Zsh, Bash, Fish:
  case Zsh:
    return ShellName(value), true
  default:
    return "", false
  }
}

// Struct of options passed by the cmd/init cmd to the RenderShellFuncs
type Opts struct {
  Cmd string
}

// Outputs shell commands to make quick navigation work
func RenderShellFuncs(shell ShellName, opts Opts) {
  tmplFile := fmt.Sprintf("templates/%s.txt", shell)

  tmpl := template.Must(template.ParseFS(templateFS, tmplFile))

  err := tmpl.Execute(os.Stdout, opts)
  if err != nil {
    fmt.Printf("Error executing template: %v\n", err)
    return
  }
}
