package frontend

// A simple program demonstrating the text area component from the Bubbles
// component library.

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
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
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(2)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 5)
	vp.SetContent(`Welcome to the chat room!
Type a message and press Enter to send.`)

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
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case tea.KeyEnter:
			m.messages = append(m.messages, m.senderStyle.Render("You: ")+m.textarea.Value())
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Reset()
			m.viewport.GotoBottom()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m model) View() string {
    // Create a larger border style for your chat UI
    borderStyle := lipgloss.NewStyle().
        BorderStyle(lipgloss.DoubleBorder()).
        BorderForeground(lipgloss.Color("63")).
        MarginTop(1).
        Padding(1, 2).
        Width(34)

    // Define the chatroom heading
    heading := " Chatroom "

    // Create a style for the heading
    headingStyle := lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("250")).
        Background(lipgloss.Color("63"))

    // Combine the heading and chat UI
    borderedChatUI := borderStyle.Render(fmt.Sprintf(
        "%s\n%s\n\n%s",
        headingStyle.Render(heading),
        m.viewport.View(),
        m.textarea.View(),
    ) + "\n\n")

    return borderedChatUI
}



// func (m model) View() string {
//     // Create a larger border style for your chat UI
//     borderStyle := lipgloss.NewStyle().
//         BorderStyle(lipgloss.DoubleBorder()).
//         BorderForeground(lipgloss.Color("63")).
//         MarginTop(1).
//         Padding(1, 2).
//         Width(34)

//     // Define the chatroom heading
//     heading := " Chatroom "

//     // Calculate the heading position to center it
//     headingWidth := len(heading)
//     leftPadding := (34 - headingWidth - 4) / 2 // Subtract border padding
//     rightPadding := leftPadding
//     if 34%2 != 0 {
//         rightPadding++
//     }

//     // Create a style for the heading
//     headingStyle := lipgloss.NewStyle().
//         Bold(true).
//         Foreground(lipgloss.Color("250")).
//         Background(lipgloss.Color("63")).
//         Padding(0, leftPadding, 0, rightPadding)

//     // Calculate the chat UI content height and adjust it
//     chatUIHeight := 16 // Adjust the height as needed
//     textAreaHeight := chatUIHeight - 4 // Adjust for the input area
//     textAreaTopPadding := chatUIHeight - textAreaHeight

//     // Create a style for the chat UI content
//     chatUIStyle := lipgloss.NewStyle().
//         Width(34). // Adjust the width as needed
//         Height(chatUIHeight - 4) // Adjust for the border and heading

//     // Combine the heading and chat UI
//     borderedChatUI := borderStyle.Render(fmt.Sprintf(
//         "%s\n%s\n%s",
//         headingStyle.Render(heading),
//         chatUIStyle.Render(m.viewport.View()),
//         strings.Repeat("\n", textAreaTopPadding)+m.textarea.View(),
//     ) + "\n\n")

//     return borderedChatUI
// }
