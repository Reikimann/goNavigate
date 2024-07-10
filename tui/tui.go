package tui

import (
  "fmt"

  "github.com/Reikimann/goNavigate/db"
  tea "github.com/charmbracelet/bubbletea"
)

// Program model/state
type model struct {
  directories []db.Directory // The directories in the list
  cursor int // Cursor position
  err error
  selected db.Directory
}

type (
  dirMsg []db.Directory
  errMsg struct{ err error }
)

// Implements the error interface for errMsg
func (e errMsg) Error() string{ return e.err.Error() } // TODO: Why is this smart?

// Creates an empty model
func NewModel() tea.Model {
  return model{}
}

// TODO: Try to pass the InitialModel to the init function. Change to small initialModel()
func getDirectories() tea.Msg {
  d, err := db.OpenDatabase()
  if err != nil {
    return errMsg{err}
  }
  defer d.Database.Close()

  directories, err := d.ListDirectories()
  if err != nil {
    return errMsg{err}
  }

  return dirMsg(directories)
}

func (m model) Init() tea.Cmd {
  return getDirectories
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case dirMsg:
    m.directories = msg
    return m, nil
  case errMsg:
    m.err = msg
    return m, tea.Quit
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q", "esc":
      return m, tea.Quit
    case "up", "k", "ctrl+k":
      if m.cursor > 0 {
        m.cursor--
      }
    case "down", "j", "ctrl+j":
      if m.cursor < len(m.directories) - 1 {
        m.cursor++
      }
    case "enter":
      // TODO: Handle navigating to the directory
      m.selected = m.directories[m.cursor]
      return m, tea.Quit
    }
  }

  return m, nil
}

func (m model) View() string {
  if m.err != nil {
    return fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err)
  }

  if (m.selected != db.Directory{}) {
    return fmt.Sprintln(m.selected.Path)
  }

  s := "Which directory would you like to navigate to?\n\n"

  for i, choice := range m.directories {
    cursor := " "

    // Is the cursor pointing at this choice?
    if m.cursor == i {
      cursor = ">"
    }

    // Render the row
    s += fmt.Sprintf("%s %s\n", cursor, choice.Path)
  }

  s += "\nPress q to quit.\n"
  return s
}
