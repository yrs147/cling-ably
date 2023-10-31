package chat

import (
	"context"
	"fmt"
	"os"

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
func SubscribeToChat(client *ably.Realtime, roomName, username string, lang string) ([]string, error) {
	channel := client.Channels.Get(roomName)
	var messages []string

	_, err := channel.SubscribeAll(context.Background(), func(msg *ably.Message) {
		// Check if the message is not from the current user
		if msg.ClientID != username {
			text, err := translate.Translate(msg.Data, lang)
			if err != nil {
				fmt.Print(err)
			}
			fmt.Printf("[%s]: %s\n", msg.ClientID, text)
			messages = append(messages, text)
		}
	})
	if err != nil {
		err := fmt.Errorf("subscribing to channel: %w", err)
		fmt.Println(err)
	}

	return messages, err
}

func Publishing(channel *ably.RealtimeChannel, msg string) {

	// reader := bufio.NewReader(os.Stdin)

	err := channel.Publish(context.Background(), "message", msg)
	// await confirmation that message was received by Ably
	if err != nil {
		err := fmt.Errorf("publishing to channel: %w", err)
		fmt.Println(err)
	}

	// for {
	// 	// text, _ := reader.ReadString('\n')
	// 	// text = strings.ReplaceAll(text, "\n", "")
	// 	// final := translate.TranslateMsg(text,"en")

	// 	// Publish the message typed in to the Ably Channel
	// 	err := channel.Publish(context.Background(), "message", msg)
	// 	// await confirmation that message was received by Ably
	// 	if err != nil {
	// 		err := fmt.Errorf("publishing to channel: %w", err)
	// 		fmt.Println(err)
	// 	}
	// }
}
