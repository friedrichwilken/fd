package hjkl

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	res "github.com/friedrichwilken/fd/pkg/resources"
)

var dbugmsg = ""
var allowDbug = true

type selection struct {
    current string
	choices []os.DirEntry // items on the to-do list
	cursor  int           // which to-do list item our cursor is pointing at
}

func New() selection {
    c, err := os.Getwd()
    if err != nil{
        panic(err)
    }

    s := selection{
        current:  c,
    }
    
    s.updateChoices()
    
    return s
}

func (m selection)updateChoices() {
    chs, err := os.ReadDir(m.current)
    if err != nil {
        panic(err)
    }

    m.choices = chs
}

func (m selection) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m selection) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
		    
        }
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m selection) View() string {
	s := ""

	if allowDbug {
		s = fmt.Sprintf("%s******************************************************************\n", s)
		s = fmt.Sprintf("%sDEBUG: %s\n", s, dbugmsg)
		s = fmt.Sprintf("%s******************************************************************\n", s)
	}

	//Show current dir
	s = fmt.Sprintf("%scurrent dir: %s\n", s, m.current)

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
