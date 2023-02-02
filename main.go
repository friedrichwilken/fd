package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

    "github.com/friedrichwilken/fd/pkg/models/hjkl"
)

func main() {
	p := tea.NewProgram(hjkl.New())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
