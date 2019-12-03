package fuelModule

import (
	"fmt"
	"math"
)

type FuelModule struct {
	mass int64
	fuel int64
}

var errNoMass = fmt.Errorf("mass was not initialized")
var errInvalidMass = fmt.Errorf("mass must be greater than 0")

func (fm *FuelModule) GetFuelNeeds() (int64, error) {
	if fm.mass == 0 {
		return 0, fmt.Errorf("%w", errNoMass)
	}

	if fm.fuel == 0 {
		requiredFuel := int64(math.Floor(float64(fm.mass)/3)) - 2

		if requiredFuel < 0 {
			requiredFuel = 0
		}

		fm.fuel = requiredFuel
	}

	return fm.fuel, nil
}

func New(mass int64) (FuelModule, error) {
	if mass <= 0 {
		return FuelModule{}, fmt.Errorf("%w: %d", errInvalidMass, mass)
	}
	return FuelModule{
		mass: mass,
	}, nil
}
