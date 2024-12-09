package tui

import (
    "fmt"
    "strings"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
    "github.com/KapilSareen/ShellFox/pkg/fetch"
)

type model struct {
	viewport    viewport.Model
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
}

func InitialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Enter a url to visit"
	ta.Focus()

	ta.Prompt = ">> "
    
    ta.SetWidth(15)
    ta.SetHeight(1)
	ta.FocusedStyle.Base = ta.FocusedStyle.Base.Foreground(lipgloss.Color("31")).Border(lipgloss.DoubleBorder(), true)
    
	ta.CharLimit = 250

	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

    vp := viewport.New(50, 14)
    vp.SetContent(lipgloss.NewStyle().Align(lipgloss.Center).Render(`
      /$$$$$$  /$$                 /$$ /$$         /$$$$$$$$                 
     /$$__  $$| $$                | $$| $$        | $$_____/                 
    | $$  \__/| $$$$$$$   /$$$$$$ | $$| $$        | $$     /$$$$$$  /$$   /$$
    |  $$$$$$ | $$__  $$ /$$__  $$| $$| $$ /$$$$$$| $$$$$ /$$__  $$|  $$ /$$/
     \____  $$| $$  \ $$| $$$$$$$$| $$| $$|______/| $$__/| $$  \ $$ \  $$$$/ 
     /$$  \ $$| $$  | $$| $$_____/| $$| $$        | $$   | $$  | $$  >$$  $$ 
    |  $$$$$$/| $$  | $$|  $$$$$$$| $$| $$        | $$   |  $$$$$$/ /$$/\  $$
     \______/ |__/  |__/ \_______/|__/|__/        |__/    \______/ |__/  \__/

        Welcome to Shell-Fox!
        A CLI-based browser for maniacs

`))


	ta.KeyMap.InsertNewline.SetEnabled(false)
    
	return model{
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.textarea.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case "enter":
			v := m.textarea.Value()

			if v == "" {
				return m, nil
			}

			resp, err := fetch.Fetch(v)
            if err != nil {
                m.messages = []string{err.Error()}
                return m, tea.Quit
            }
			m.messages = []string{resp}
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Reset()
            m.viewport.Height = 14
			m.viewport.GotoBottom()

			return m, nil
            

		default:
			var cmd tea.Cmd
			m.textarea, cmd = m.textarea.Update(msg)
			return m, cmd
		}

	case cursor.BlinkMsg:
		var cmd tea.Cmd
		m.textarea, cmd = m.textarea.Update(msg)
		return m, cmd

	default:
		return m, nil
	}
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.textarea.View(),
		m.viewport.View(),
	) + "\n\n"
}