package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/KapilSareen/ShellFox/pkg/tui"
	// "github.com/KapilSareen/ShellFox/pkg/fetch"
	// "github.com/KapilSareen/ShellFox/pkg/parse"
)

func main() {

	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
	}
                              // Example usage of Parser
	// html := `<html> <head> <title>Example Domain</title></head> <body> <div> <h1>Example Domain</h1> <p>This domain is for use in illustrative examples in documents. You may use this domain in literature without prior coordination or asking for permission.</p> <p><a href="https://www.iana.org/domains/example">More information...</a></p> </div> </body> </html>`
	// // Parse the html
	// node := parse.Parse(html)
	// // Print the node tree
	// fmt.Println("Printing DOM tree...")
	// node.PrintDOM("")
}
