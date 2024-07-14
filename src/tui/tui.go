package tui

import (
	// "fmt"
	"strings"

	"github.com/Reikimann/goNavigate/src/db"
	"github.com/Reikimann/goNavigate/src/utils"
	tea "github.com/charmbracelet/bubbletea"
)

// Program Model/state
type Model struct {
  directories []db.Directory // The directories in the list
  cursor int // Cursor position
  selected db.Directory // The selected directory

  err error
  dbEmpty bool // If the DB empty
  quitting bool // If the program is to quit
}

type (
  dirMsg []db.Directory // This msg holds the directories from the DB
  dirEmptyMsg struct{} // This msg conveys that the DB is empty
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

  if len(directories) == 0 {
    return dirEmptyMsg{}
  }
  return dirMsg(directories)
}

func (m Model) Init() tea.Cmd {
  return getDirectories
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case errMsg:
    m.err = msg
    return m, tea.Quit
  case dirEmptyMsg:
    m.dbEmpty = true
    return m, tea.Quit
  case dirMsg:
    m.directories = msg
    return m, nil
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "esc":
      m.quitting = true
      return m, tea.Quit
    case "up", "k", "ctrl+k":
      if m.cursor > 0 {
        m.cursor--
      } else {
        m.cursor = len(m.directories) - 1
      }
    case "down", "j", "ctrl+j":
      if m.cursor < len(m.directories) - 1 {
        m.cursor++
      } else {
        m.cursor = 0
      }
    case "enter":
      m.selected = m.directories[m.cursor]
      return m, tea.Quit
    }
  }

  return m, nil
}

func (m Model) View() string {
  var b strings.Builder

  if m.err != nil {
    b.WriteString("\nWe had some trouble: " + m.err.Error() + "\n\n")
    return b.String()
  }

  // TODO: Figure out how to show this in a screen and make the user press q to quit
  if m.dbEmpty {
    b.WriteString("You haven't added any directories yet. Please check the help command.")
    return b.String()
  }

  // Return not printing TUI if a directory has been selected or the program is to shut down
  if m.selected != (db.Directory{}) || m.quitting {
    return b.String()
  }

  // TODO: Add a border around the path selection
  b.WriteString("Which directory would you like to navigate to?\n\n")

  for i, choice := range m.directories {
    cursor := " "

    // TODO: Make the line highlighted (lib lipgloss)
    if m.cursor == i {
      cursor = ">"
    }

    // Strips the /home/user/ from a given path.
    path, err := utils.DirPathStripHome(choice.Path)
    if err != nil {
      b.WriteString(" " + cursor + " " + choice.Path + "\n")
    } else {
      b.WriteString(" " + cursor + " " + path + "\n")
    }
  }

  return b.String()
}
