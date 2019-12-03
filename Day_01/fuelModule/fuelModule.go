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
	return fm.getFuelNeeds(getFuelFromMass)
}

func (fm *FuelModule) GetCompleteFuelNeeds() (int64, error) {
	return fm.getFuelNeeds(getCompleteFuelFromMass)
}

func (fm *FuelModule) getFuelNeeds(fuelFunc func(int64) int64) (int64, error) {
	if fm.mass == 0 {
		return 0, fmt.Errorf("%w", errNoMass)
	}

	if fm.fuel == 0 {
		requiredFuel := fuelFunc(fm.mass)

		if requiredFuel < 0 {
			requiredFuel = 0
		}

		fm.fuel = requiredFuel
	}

	return fm.fuel, nil
}

func getCompleteFuelFromMass(value int64) int64 {
	requiredFuel := getFuelFromMass(value)

	if requiredFuel <= 0 {
		return 0
	}

	return requiredFuel + getCompleteFuelFromMass(requiredFuel)
}

func getFuelFromMass(value int64) int64 {
	return int64(math.Floor(float64(value)/3)) - 2
}

func New(mass int64) (FuelModule, error) {
	if mass <= 0 {
		return FuelModule{}, fmt.Errorf("%w: %d", errInvalidMass, mass)
	}
	return FuelModule{
		mass: mass,
	}, nil
}
