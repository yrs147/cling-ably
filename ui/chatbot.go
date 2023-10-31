package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yrs147/cling-ably/internal/chat" // Import your chat package
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
	return nil
}
// 	m.textarea, cmd = m.textarea.Update(msg)
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	client, err := chat.InitializeClient("Nina")
	channel := client.Channels.Get("sample")

	m.textarea, cmd = m.textarea.Update(msg)

	if msg, ok := msg.(tea.KeyMsg); ok {
		if msg.Type == tea.KeyEnter {
			text := m.textarea.View()

			// Handle sending the message directly in the UI
			chat.Publishing(channel, text)
			if err != nil {
				fmt.Println("Error sending message:", err)
			}

			m.textarea = textarea.New() // Clear the textarea
		}
	}

	// Handle receiving messages directly in the UI
	newMessages, err := chat.SubscribeToChat(client, "sample", "Nina", "en")
	if err != nil {
		fmt.Println("Error receiving messages:", err)
	}

	// Update the messages slice with new messages
	m.messages = append(m.messages, newMessages...)

	return m, cmd
}



func (m model) View() string {
	borderStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("63")).
		MarginTop(1).
		Padding(1, 2).
		Width(34)

	heading := " Chatroom "

	headingStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("250")).
		Background(lipgloss.Color("63"))

	borderedChatUI := borderStyle.Render(fmt.Sprintf(
		"%s\n%s\n\n%s",
		headingStyle.Render(heading),
		m.textarea.View(),
		strings.Join(m.messages, "\n"),
	) + "\n\n")

	return borderedChatUI
}
