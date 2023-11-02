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

	fmt.Printf("🚀 Welcome to %s!\n", roomName)
	fmt.Printf("👉 Enter your message below: (type 'exit' to leave)\n")

	memberColors := make(map[string]int) 

	_, err := channel.SubscribeAll(context.Background(), func(msg *ably.Message) {
		if msg.ClientID != username {
			text, err := translate.Translate(msg.Data, lang)
			if err != nil {
				fmt.Println(err)
				return
			}
			
			// Determine and Assign color for the member
			color, exists := memberColors[msg.ClientID]
			if !exists {
				
				color = len(memberColors) + 1
				memberColors[msg.ClientID] = color
			}
			
		
			colorCode := 31 + color 
			
	
			fmt.Printf("\033[2K\r\033[%dm👉 %s: %s\n\033[0m👉 You: ", colorCode, msg.ClientID, text)
		}
	})

	if err != nil {
		fmt.Printf("Error subscribing to channel: %v\n", err)
	}

	publishing(channel)
}

func publishing(channel *ably.RealtimeChannel) {
	fmt.Print("👉 You: ")

	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.ReplaceAll(text, "\n", "")

		if text == "exit" {
			break
		}

		err := channel.Publish(context.Background(), "message", text)
		if err != nil {
			fmt.Printf("Error publishing message: %v\n", err)
		}
		
	}
	
}
