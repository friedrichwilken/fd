package hjkl

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/friedrichwilken/fd/pkg/path"
	res "github.com/friedrichwilken/fd/pkg/resources"
)

var dbugmsg = ""
var allowDbug = true

type model struct {
	current path.Path
	choices []os.DirEntry // items on the to-do list
	cursor  int           // which to-do list item our cursor is pointing at
}

func InitialModel() model {
	cur, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	chs, err := os.ReadDir(cur)
	if err != nil {
		panic(err)
	}

	s := model{
		current: path.New(cur, "/"),
		choices: chs,
	}

	return s
}

func (m model) enterDir() (tea.Model, tea.Cmd) {
	dir := m.choices[m.cursor]
	if !dir.IsDir() {
		return m, nil
	}

	m.current.GoToSub(dir.Name())

	chs, err := os.ReadDir(m.current.String())
	if err != nil {
		panic(err)
	}

	m.choices = chs
	m.cursor = 0

	return m, nil
}

func (m model) parentDir() (tea.Model, tea.Cmd) {
	err := m.current.GoToParent()
	if err != nil {
		// todo display "cant move to paren, there is none"
	}

	chs, err := os.ReadDir(m.current.String())
	if err != nil {
		panic(err)
	}

	m.choices = chs
	m.cursor = 0

	return m, nil
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "h", "left":
			return m.parentDir()

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ", "right", "l":
			return m.enterDir()
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	s := ""

	//Show current dir
	s = fmt.Sprintf("%scurrent dir: %s\n\n", s, m.current.String())

	// Iterate over our choices
	for i, choice := range m.choices {
		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf(" %s %s\n", cursor, entryToString(choice))
	}

	// The footer
	s += "\nPress q to quit.\n"

	if allowDbug {
		s = fmt.Sprintf("%s******************************************************************\n", s)
		s = fmt.Sprintf("%sDEBUG: %s\n", s, dbugmsg)
		s = fmt.Sprintf("%s******************************************************************\n", s)
	}

	// Send the UI for rendering
	return s
}

func entryToString(e os.DirEntry) string {
	// The retruned string will be either "<filesymbol> filename" or "<dirsymbol> dirname"
	// Define if we use a file or dir symbol
	s := ""
	if e.IsDir() {
		s = res.DirSymbol
	} else {
		s = res.FileSymbol
	}

	return fmt.Sprintf("%s %s", s, e.Name())
}
