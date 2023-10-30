package chat

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/ably/ably-go/ably"
	"github.com/joho/godotenv"
	"github.com/yrs147/cling-ably/internal/translate"
)

// InitializeClient initializes the Ably client.
func InitializeClient(username string) (*ably.Realtime, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading environment variables file")
		return nil, err
	}
	key := os.Getenv("ABLY_KEY")

	// Initialize the client with the username as the client ID
	client, err := ably.NewRealtime(ably.WithKey(key), ably.WithClientID(username))
	if err != nil {
		return nil, err
	}

	return client, nil
}

// SubscribeToChat subscribes to chat messages in the specified room.
func SubscribeToChat(client *ably.Realtime, roomName, username string, lang string) {
	channel := client.Channels.Get(roomName)

	_, err := channel.SubscribeAll(context.Background(), func(msg *ably.Message) {
		// Check if the message is not from the current user
		if msg.ClientID != username {
			text,err := translate.Translate(msg.Data,lang)
			if err!=nil {
				fmt.Print(err)
			}
			fmt.Printf("[%s]: %s\n", msg.ClientID, text)
		}
	})
	if err != nil {
		err := fmt.Errorf("subscribing to channel: %w", err)
		fmt.Println(err)
	}

	// Let the user send messages
	publishing(channel)
}


func subscribeToEvent(channel *ably.RealtimeChannel, username string) (func(), error) {
	unsubscribe, err := channel.Subscribe(context.Background(), "message", func(msg *ably.Message) {
		if msg.ClientID != username {
			// Display received messages from other users
			fmt.Printf("[%s]: %v\n", msg.ClientID, msg.Data)
		}
	})
	if err != nil {
		fmt.Printf("Error subscribing to channel: %v\n", err)
		return nil, err
	}

	// Subscribe to presence events (people entering and leaving) on the channel
	_, pErr := channel.Presence.SubscribeAll(context.Background(), func(msg *ably.PresenceMessage) {
		if msg.Action == ably.PresenceActionEnter {
			fmt.Printf("%v has entered the chat\n", msg.ClientID)
		} else if msg.Action == ably.PresenceActionLeave {
			fmt.Printf("%v has left the chat\n", msg.ClientID)
		}
	})
	if pErr != nil {
		err := fmt.Errorf("subscribing to presence in channel: %w", pErr)
		fmt.Println(err)
	}

	return unsubscribe, nil
}

func publishing(channel *ably.RealtimeChannel) {
	
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.ReplaceAll(text, "\n", "")
		// final := translate.TranslateMsg(text,"en")

		// Publish the message typed in to the Ably Channel
		err := channel.Publish(context.Background(), "message", text)
		// await confirmation that message was received by Ably
		if err != nil {
			err := fmt.Errorf("publishing to channel: %w", err)
			fmt.Println(err)
		}
	}
}
