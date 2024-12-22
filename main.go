package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	SpawnNewWorld()

	p := tea.NewProgram(NewRenderModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error Initializing Game: %v", err)
		os.Exit(1)
	}
}
