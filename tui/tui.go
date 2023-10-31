package tui

import (
	"fmt"
	"log"

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

	// Create a chat area view
	if v, err := g.SetView("chat", 0, 0, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Messages  -  <Chat Room>"
		v.FgColor = gocui.ColorRed
		v.Autoscroll = true
		v.Wrap = true
		chatView = v
	}

	// Create an input field view at the bottom
	if v, err := g.SetView("input", 0, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " New Message  -  <" + "User" + "> "
		v.FgColor = gocui.ColorWhite
		v.Editable = true
		err = v.SetCursor(0, 0)
		if err != nil {
			log.Println("Failed to set cursor:", err)
		}
		inputView = v
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
	_,SizeY := chatView.Size()
	if chatView != nil {
		// Append the message to the chat view
		fmt.Fprintln(chatView, message)

		// Scroll the chat view to the bottom
		_ = chatView.SetOrigin(0, SizeY-3)
	}
}
