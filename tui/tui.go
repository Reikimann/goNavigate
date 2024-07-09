package tui

import (
  "github.com/Reikimann/goNavigate/db"
  tea "github.com/charmbracelet/bubbletea"
)

type model struct {
  directories []db.Directory // The directories in the list
  cursor int // Cursor position
}

func initialModel() model {
  return model{
    directories: make([]db.Directory, 0),
  }
}

func (m model) Init() tea.Cmd {
  return nil
}

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//   switch msg := msg.(type) {
//   case tea.KeyMsg:
//     switch msg.String() {
//     case "ctrl+c", "q":
//       return m, tea.Quit
//     case "up", "k":
//       if m.cursor > 0 {
//         m.cursor--
//       }
//     case "down", "j":
//       if m.cursor < len(m.directories) - 1 {
//         m.cursor++
//       }
//     case "enter":
//       // TODO: Handle navigating to the directory
//     }
//   }
// }
