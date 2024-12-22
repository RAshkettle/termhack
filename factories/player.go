package factories

import (
	"github.com/RAshkettle/termhack/components"
	"github.com/yohamta/donburi"
)

func SpawnNewPlayer(world donburi.World) {
	// Create the entity
	entity := world.Create(components.Player, components.Position)

	// Add to the world to get an ID
	entry := world.Entry(entity)

	// Set default values
	components.Position.SetValue(entry, components.PositionData{X: 10, Y: 10})

	// TODO: Change this position value to something real once we have map generation working
}
