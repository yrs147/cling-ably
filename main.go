package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yrs147/cling-ably/ui"
)

func main() {

	// input := tea.NewProgram(input.InputModel())

	// // Run the spinner program
	// _, err := input.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Create a new spinner program
	// spinner := tea.NewProgram(ui.NewModel())

	// // Run the spinner program
	// _, err = spinner.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create a chat UI program (frontend.InitialModel() is used for chat UI)
	// Initialize a chat model with the desired username, room code, and language
	model := ui.InitialModel()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create a chat UI program with the initialized model
	p := tea.NewProgram(model)

	// Start the chat UI program
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

	// Once the chat UI program finishes, exit
	fmt.Println("Chat UI program finished")
}
