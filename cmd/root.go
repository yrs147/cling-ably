package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yrs147/cling-ably/internal/chat"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cling",
	Short: "A CLI chat app using Ably",
	Long: `Cling is a CLI chat application that allows user to communicate
	with others in real time through their CLI!!!`,
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
var roomLang string

func init() {
	
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&username, "username", "u", "defaultUsername", "Your username for the chat")
	rootCmd.Flags().StringVarP(&roomCode, "room", "r", "defaultRoomCode", "Chat room code")
    rootCmd.Flags().StringVarP(&roomLang, "language", "l", "en", "Chat room language code")
}

func InitializeAblyAndSubscribe() error {
    // Initialize the Ably client with the username as the client ID
    client, err := chat.InitializeClient(username)
    if err != nil {
        fmt.Printf("Error initializing Ably client: %v\n", err)
        return err
    }

    chat.SubscribeToChat(client, roomCode, username,roomLang)
    

    return nil
}

