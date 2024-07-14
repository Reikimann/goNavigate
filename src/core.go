package gonavigate

import (
  "fmt"
  "os"

  "github.com/Reikimann/goNavigate/src/tui"

  tea "github.com/charmbracelet/bubbletea"
)

// TODO: Handle options: func Run(opts Options) (int, error) {
func Run() {
  // See NOTES.md
  tty, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
  if err != nil {
    panic(err)
  }
  defer tty.Close()

  // BUG: The TUI doesn't clear on ctrl+c. Can this be fixed with bubbletea?
  p := tea.NewProgram(tui.NewModel(), tea.WithOutput(tty))
  finalModel, err := p.Run()
  if err != nil {
    fmt.Printf("There has been an error: %v", err)
    os.Exit(1)
  }

  if m, ok := finalModel.(tui.Model); ok {
    if !m.DBContainsDirs() {
      os.Exit(1)
    }

    fmt.Println(m.SelectedDir().Path)
  } else {
    fmt.Println("Failed to assert model type")
    os.Exit(1)
  }
}
