package tui

import (
	"context"
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/yrs147/cling-ably/internal/chat"
	"github.com/ably/ably-go/ably"
)

var inputView *gocui.View

func CreateGui(channel, username, language string) (*gocui.Gui, error) {
	// Initialize the Ably client with the username as the client ID
	client, err := chat.InitializeClient(username)
	if err != nil {
		return nil, err
	}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}
	

	maxX, maxY := g.Size()

	// Chatroom section
	if ov, err := g.SetView("chatroom", 0, 0, maxX-21, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return nil, err
		}
		ov.Title = "Chat Room"
		ov.FgColor = gocui.ColorCyan
		ov.Autoscroll = true
		ov.Wrap = true
	}

	// Input field
	if iv, err := g.SetView("input", 0, maxY-3, maxX-21, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return nil, err
		}
		iv.Title = "Input"
		iv.Editable = true
		inputView = iv
		if _, err := g.SetCurrentView("input"); err != nil {
			return nil, err
		}
	}

	// Members online section
	if v, err := g.SetView("members", maxX-20, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return nil, err
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

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return nil, err
	}

	err = g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, iv *gocui.View) error {
		iv.Rewind()
		if len(iv.Buffer()) >= 2 {
			message := iv.Buffer()
			ablych := client.Channels.Get(channel)
			err := ablych.Publish(context.Background(), "message", message)
			if err != nil {
				log.Println("Failed to publish message:", err)
			}
			iv.Clear()
			iv.SetCursor(0, 0)
		}
		return nil
	})

	if err != nil {
		log.Println("Cannot bind the enter key:", err)
	}

	// Subscribe to the Ably channel
	ablych := client.Channels.Get(channel)
	_, err = ablych.SubscribeAll(context.Background(), func(msg *ably.Message) {
		if msg.ClientID != username {
			text := msg.Data.(string)
			g.Update(func(g *gocui.Gui) error {
				ov, err := g.View("chatroom")
				if err != nil {
					log.Println("Cannot get chatroom view:", err)
					return nil
				}
				_, err = fmt.Fprintf(ov, "[%s]: %s\n", msg.ClientID, text)
				if err != nil {
					log.Println("Cannot print to chatroom view:", err)
				}
				return nil
			})
		}
	})

	if err != nil {
		log.Println("Failed to subscribe to channel:", err)
	}

	return g, nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
