package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yrs147/cling-ably/internal/chat"
)

var rootCmd = &cobra.Command{
    Use:   "cling",
    Short: "A CLI chat app using Ably",
    Long:  "Cling is a CLI chat application that allows users to communicate in real-time through their CLI",
    // Run: func(cmd *cobra.Command, args []string) {
    //     // Your chat logic goes here, using the username, roomCode, and message flags.
    //     // For example, you can use these flags to create/join a chat room or send messages.
    // },
}


// rootCmd represents the base command when called without any subcommands
var cling = &cobra.Command{
	Use:   "connect",
	Short: "Connect command is used to create chatrooms",
	Long: `Connect command is used to create chatroom by passing username and roomname as flag`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		err := InitializeAblyAndSubscribe()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var username string
var roomCode string

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cling-ably.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cling.Flags().StringVarP(&username, "username", "u", "defaultUsername", "Your username for the chat")
	cling.Flags().StringVarP(&roomCode, "room", "r", "defaultRoomCode", "Chat room code")

	rootCmd.AddCommand(cling)
}

func InitializeAblyAndSubscribe() error {
    // Initialize the Ably client with the username as the client ID
    client, err := chat.InitializeClient(roomCode, username)
    if err != nil {
        fmt.Printf("Error initializing Ably client: %v\n", err)
        return err
    }

    unsubscribe, err := chat.SubscribeToChat(client, roomCode, username)
    if err != nil {
        fmt.Printf("Error subscribing to chat: %v\n", err)
        return err
    }
    defer unsubscribe()

    // Add any additional logic or commands you want to execute here

    return nil
}

