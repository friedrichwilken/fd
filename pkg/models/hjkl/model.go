package hjkl

import (
    	"fmt"
        "log"
        "os"

	tea "github.com/charmbracelet/bubbletea"

    
    res "github.com/friedrichwilken/fd/pkg/resources"
)

type selection struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
}

func New() selection{
	return selection{
		// Our to-do list is a grocery list
		choices: dirContent(), 
	}
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
		    //todo move one dir down
        }
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m selection) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		//checked := " " // not selected
		//if _, ok := m.selected[i]; ok {
		//	checked = "x" // selected!
		//}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func dirContent() []string {    
    entries, err := os.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }
 
    names := []string{}
    for _, e := range entries {
        names = append(names, entryToString(e))
    }
    return names
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