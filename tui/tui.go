package tui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var inputView *gocui.View
var chatView *gocui.View

func CreateGui() (*gocui.Gui, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}

	g.SetManagerFunc(Layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return nil, err
	}
	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, sendMessage); err != nil {
		return nil, err
	}

	// go drawchat(channel, username, g)

	return g, nil
}

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// Chatroom section
	if v, err := g.SetView("chatroom", 0, 0, maxX-21, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Chat Room"
		v.FgColor = gocui.ColorCyan
		v.Autoscroll = true
		v.Wrap = true
		// Add any logic to display chat messages here
	}

	// Input field
	if v, err := g.SetView("input", 0, maxY-3, maxX-21, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Input"
		v.Editable = true
		// Add logic for handling user input here
	}

	// Members online section
	if v, err := g.SetView("members", maxX-20, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Members Online"
		v.FgColor = gocui.ColorGreen
		v.Autoscroll = true
		v.Wrap = true
		// Add logic to display online members here
		// For now, you can add some dummy data
		dummyMembers := []string{"User1", "User2", "User3", "User4"}
		for _, member := range dummyMembers {
			fmt.Fprintln(v, member)
		}
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func sendMessage(g *gocui.Gui, v *gocui.View) error {
	if inputView != nil {
		// message := inputView.Buffer()
		inputView.Clear()
		inputView.SetCursor(0, 0)

		// You can use a function here to send the message to the chat room
		// For now, we'll print the message to the chat area
		// addChatMessage(fmt.Sprintf("<%s>: %s", username, message))

		// Replace the following line with actual sending logic
		// sendToChatRoom(message)
	}

	return nil
}

func addChatMessage(message string) {
	_, SizeY := chatView.Size()
	if chatView != nil {
		// Append the message to the chat view
		fmt.Fprintln(chatView, message)

		// Scroll the chat view to the bottom
		_ = chatView.SetOrigin(0, SizeY-3)
	}
}
