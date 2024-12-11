package tui

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/KapilSareen/ShellFox/pkg/fetch"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	isLoading    bool = false

	spinners = []spinner.Spinner{
		spinner.Line,
		spinner.Dot,
		spinner.MiniDot,
		spinner.Jump,
		spinner.Pulse,
		spinner.Points,
		spinner.Globe,
		spinner.Moon,
		spinner.Monkey,
	}

	LoadingMsgs = []string{
		"Hold tight, magic is happening..",
		"Loading, please wait..",
		"Good things take time..",
		"Almost there..",
		"Preparing your content..",
		"Hang on, we're getting your data..",
		"Just a moment..",
		"Your request is being processed..",
		"Stay tuned, loading in progress..",
	}

	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
)

type model struct {
	viewport    viewport.Model
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
	fullscreen  bool
	index       int
	spinner     spinner.Model
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
	vp.SetContent(lipgloss.NewStyle().
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("231")).
		Render(`
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

	sp := spinner.New()
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("31"))

	return model{
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
		fullscreen:  false,
		spinner:     sp,
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

			isLoading = true
			m.index = rand.Intn(len(LoadingMsgs))
			m.resetSpinner()
			return m, tea.Batch(m.spinner.Tick, fetchData(v))

		case "ctrl+l":
			m.index = rand.Intn(len(LoadingMsgs))
			isLoading = !isLoading
			m.resetSpinner()
			return m, m.spinner.Tick

		case "ctrl+f":
			m.fullscreen = !m.fullscreen
			if m.fullscreen {
				return m, tea.EnterAltScreen
			} else {
				return m, tea.ExitAltScreen
			}

		default:
			var cmd tea.Cmd
			m.textarea, cmd = m.textarea.Update(msg)
			return m, cmd
		}

	case fetchResponseMsg:
		isLoading = false
		m.messages = []string{string(msg)}
		m.viewport.SetContent(strings.Join(m.messages, "\n"))
		m.textarea.Reset()
		m.viewport.Height = 14
		m.viewport.GotoBottom()
		return m, nil

	case cursor.BlinkMsg:
		var cmd tea.Cmd
		m.textarea, cmd = m.textarea.Update(msg)
		return m, cmd

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	default:
		return m, nil
	}
}

type fetchResponseMsg string

func fetchData(url string) tea.Cmd {
	return func() tea.Msg {
		response := fetch.Fetch(url)
		return fetchResponseMsg(response)
	}
}

func (m model) View() string {
	if isLoading {
		return fmt.Sprintf(
			"%s\n\n %s %s",
			m.textarea.View(),
			m.spinner.View(),
			LoadingMsgs[m.index],
		) + "\n\n"
	}
	return fmt.Sprintf(
		"%s\n\n%s",
		m.textarea.View(),
		m.viewport.View(),
	) + "\n\n"
}

func (m *model) resetSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinners[m.index]
}
