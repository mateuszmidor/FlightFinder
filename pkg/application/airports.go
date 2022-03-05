package application

import (
	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure"
)

type Airports struct {
	flightsData infrastructure.FlightsData
}

func NewAirports(repo infrastructure.FlightsDataRepo) *Airports {
	flightsData := repo.Load()
	return &Airports{flightsData: flightsData}
}

func (a *Airports) AllAirports() airports.Airports {
	return a.flightsData.Airports
}

func (a *Airports) AirportsByCountry(twoLettersCode string) airports.Airports {
	result := make(airports.Airports, 0)
	for _, airport := range a.flightsData.Airports {
		if airport.Name() == twoLettersCode {
			result = append(result, airport)
		}
	}
	return result
}
