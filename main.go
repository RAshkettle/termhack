package main

import (
	"fmt"
	"os"

	"github.com/RAshkettle/termhack/ecs"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	ecs.SpawnNewWorld()
	p := tea.NewProgram(NewRenderModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error Initializing Game: %v", err)
		os.Exit(1)
	}
}
