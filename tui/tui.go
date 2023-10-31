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
		chatView = v // Store chatView for future use
	}

	// Input field
	if v, err := g.SetView("input", 0, maxY-3, maxX-21, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Input"
		v.Editable = true
		inputView = v // Store inputView for future use
		if _, err := g.SetCurrentView("input"); err != nil {
			return err
		}
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
		message := inputView.Buffer()
		inputView.Clear()
		inputView.SetCursor(0, 0)

		// Display the message in the chatroom as "You"
		addChatMessage("You: " + message)
	}

	return nil
}

func addChatMessage(message string) {
	if chatView != nil {
		// Append the message to the chat view
		fmt.Fprintln(chatView, message)
	}
}
