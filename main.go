package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ably/ably-go/ably"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
	key := os.Getenv("ABLY_KEY")

	client, err := ably.NewRealtime(ably.WithKey(key))
	if err != nil {
		fmt.Printf("Error creating Ably client: %v\n", err)
		return
	}

	defer client.Close()

	// Prompt for a username
	fmt.Print("Enter your username: ")
	var username string
	fmt.Scanln(&username)

	checkSubscribeToEvent(client, username)
}

func checkSubscribeToEvent(client *ably.Realtime, username string) {
	// Connect to the Ably Channel with a specific name (e.g., "chat")
	channelName := "chat"
	channel := client.Channels.Get(channelName)

	unsubscribe := subscribeToEvent(channel, username)

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

	unsubscribe()
}

func subscribeToEvent(channel *ably.RealtimeChannel, username string) func() {
	// Subscribe to messages sent on the channel with a specific event name (e.g., "message")
	unsubscribe, err := channel.Subscribe(context.Background(), "message", func(msg *ably.Message) {
		fmt.Printf("Received message from %v: '%v'\n", msg.ClientID, msg.Data)
	})
	if err != nil {
		fmt.Printf("Error subscribing to channel: %v\n", err)
		return nil
	}

	// Announce presence with the username
	err = channel.Presence.Enter(context.Background(), username)
	if err != nil {
		fmt.Printf("Error announcing presence: %v\n", err)
		return nil
	}

	return unsubscribe
}

func publish(channel *ably.RealtimeChannel, username, message string) {
	// Publish the message to the Ably Channel with a specific event name (e.g., "message")
	err := channel.Publish(context.Background(), "message", fmt.Sprintf("[%s]: %s", username, message))
	if err != nil {
		fmt.Printf("Error publishing message: %v\n", err)
	}
}

