package components

import (
	"github.com/yohamta/donburi"
)

// These components contain no data but exist just to identify an entity
type PlayerMarker struct{}

var Player = donburi.NewComponentType[PlayerMarker]()

type MonsterMarker struct{}

var Monster = donburi.NewComponentType[MonsterMarker]()

type MapTileMarker struct{}

var MapTile = donburi.NewComponentType[MapTileMarker]()

type ContainerMarker struct{}

var Container = donburi.NewComponentType[ContainerMarker]()
