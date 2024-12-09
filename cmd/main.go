package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/KapilSareen/ShellFox/pkg/tui"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
	}

	
	fmt.Print("Exiting ShellFox\n")
}
