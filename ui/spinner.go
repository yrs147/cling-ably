package frontend

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
	"time"
)

type Model struct {
	spinner spinner.Model
	startTime time.Time
}

func NewModel() Model {
	s := spinner.NewModel()
	s.Spinner = spinner.Globe
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return Model{spinner: s, startTime: time.Now()}
}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)

	// Check if 5 seconds have passed
	if time.Since(m.startTime) >= 5 * time.Second {
		return m, tea.Quit
	}

	return m, cmd
}

func (m Model) View() string {
	str := "   Loading..."
	return str + m.spinner.View() + "\n"
}
