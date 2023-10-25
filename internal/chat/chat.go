package chat

import (
	"context"
	"fmt"
	"os"

	"github.com/ably/ably-go/ably"
	"github.com/joho/godotenv"
)

// InitializeClient initializes the Ably client.
func InitializeClient(ablyKey, username string) (*ably.Realtime, error) {
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
func SubscribeToChat(client *ably.Realtime, roomName, username string) (func(), error) {
	channel := client.Channels.Get(roomName)

	unsubscribe, err := subscribeToEvent(channel, username)
	if err != nil {
		return nil, err
	}
	// Let the user send messages
	for {
		fmt.Print("Type a message (or 'exit' to leave): ")
		var message string
		fmt.Scanln(&message)

		if message == "exit" {
			break
		}

		publish(channel, username, message)
	}

	return unsubscribe, nil
}

// PublishMessage publishes a message to the chat room.
// func PublishMessage(client *ably.Realtime, roomName, username, message string) error {
// 	channel := client.Channels.Get(roomName)

// 	err := publish(channel, username, message)
// 	return err
// }

func subscribeToEvent(channel *ably.RealtimeChannel, username string) (func(), error) {
	unsubscribe, err := channel.Subscribe(context.Background(), "message", func(msg *ably.Message) {
		fmt.Printf("Received message from %v: '%v'\n", msg.ClientID, msg.Data)
	})
	if err != nil {
		fmt.Printf("Error subscribing to channel: %v\n", err)
		return nil, err
	}

	err = channel.Presence.Enter(context.Background(), username)
	if err != nil {
		fmt.Printf("Error announcing presence: %v\n", err)
		return nil, err
	}

	return unsubscribe, nil
}

func publish(channel *ably.RealtimeChannel, username, message string) error {
	err := channel.Publish(context.Background(), "message", fmt.Sprintf("[%s]: %s", username, message))
	if err != nil {
		fmt.Printf("Error publishing message: %v\n", err)
	}

	return err
}
