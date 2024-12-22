package ecs

import (
	"sync"

	"github.com/RAshkettle/termhack/factories"
	"github.com/yohamta/donburi"
)

// We want World available as a Singleton so we can always get it
var (
	singleton donburi.World
	once      = sync.Once{}
)

func GetWorld() donburi.World {
	once.Do(func() {
		singleton = donburi.NewWorld()
	})

	return singleton
}

func SpawnNewWorld() {
	world := GetWorld()
	factories.SpawnNewPlayer(world)
	//
}
