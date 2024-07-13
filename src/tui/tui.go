package tui

import (
  "fmt"

  "github.com/Reikimann/goNavigate/src/db"
  tea "github.com/charmbracelet/bubbletea"
)

// Program Model/state
type Model struct {
  directories []db.Directory // The directories in the list
  cursor int // Cursor position
  selected db.Directory // The selected directory

  err error
  dbEmpty bool
}

type (
  dirMsg []db.Directory
  errMsg struct{ err error }
)

// Implements the error interface for errMsg
func (e errMsg) Error() string{ return e.err.Error() } // TODO: Why is this smart?

// Creates an empty model
func NewModel() tea.Model {
  return Model{
    selected: db.Directory{},
  }
}

func (m Model) SelectedDir() db.Directory {
  return m.selected
}

func (m Model) DBContainsDirs() bool {
  return len(m.directories) != 0
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

func (m Model) Init() tea.Cmd {
  return getDirectories
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case dirMsg:
    m.directories = msg

    // TODO: Somehow make this into a msg isEmptyMsg
    if len(m.directories) == 0 {
      m.dbEmpty = true
      return m, tea.Quit
    }
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

func (m Model) View() string {
  if m.err != nil {
    return fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err)
  }

  // TODO: Figure out how to show this in a screen and make the user press q to quit
  if m.dbEmpty {
    return fmt.Sprintln("You haven't added any directories yet. Please check the help command.")
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
