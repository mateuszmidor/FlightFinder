package application

import (
	"errors"
	"fmt"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure"
)

var NotFoundError = errors.New("not found error")

type Airports struct {
	flightsData infrastructure.FlightsData
}

func NewAirports(repo infrastructure.FlightsDataRepo) *Airports {
	flightsData := repo.Load()
	return &Airports{flightsData: flightsData}
}

func (a *Airports) ByIATACode(code string) (airports.Airport, error) {
	id :=  a.flightsData.Airports.GetByCode(code)
	if id == airports.NullID {
		return airports.Airport{}, fmt.Errorf("airport IATA code %q not found: %w",code, NotFoundError)
	}
	return a.flightsData.Airports.Get(id), nil
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
