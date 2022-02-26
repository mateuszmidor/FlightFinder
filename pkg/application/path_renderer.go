package application

import (
	"github.com/mateuszmidor/FlightFinder/pkg/domain/pathfinding"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure"
)

type PathRenderer interface {
	Render(paths []pathfinding.Path, flightsData *infrastructure.FlightsData)
}
