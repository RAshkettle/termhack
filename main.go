package main

import (
	"fmt"
	"os"

	"github.com/RAshkettle/termhack/ecs"
	"github.com/RAshkettle/termhack/systems"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	ecs.SpawnNewWorld()
	p := tea.NewProgram(systems.NewRenderModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error Initializing Game: %v", err)
		os.Exit(1)

	}
}
