package main

//This will be for all rendering to screen.
//This is done via bubbletea
//

import (
	tea "github.com/charmbracelet/bubbletea"
)

// This struct will contain all the renderable data that we need to show on the screen.
type RenderModel struct {
	Maptiles []string
	Items    []string
	Monsters []string
	Player   string
}

// This function does not get called but it necessary for the interface
// Instead, we used the normal new pattern to create the model.
func (r RenderModel) Init() tea.Cmd {
	return nil
}

// Listen for Key Events
// This triggers the Player's Turn
// Immediately after the player's turn, it should trigger the monster turn
func (r RenderModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

// All actual Drawing logic in our game will go here.
// This will redraw our screen on each rendering
func (r RenderModel) View() string {
	return ""
}

// Simple New Pattern to create our model.
func NewRenderModel() RenderModel {
	return RenderModel{
		Maptiles: []string{},
		Items:    []string{},
		Monsters: []string{},
		Player:   "",
	}
}
