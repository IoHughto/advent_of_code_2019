package fuelModule

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	fuelModuleTests := []fuelModuleTest{
		{1, nil, 1},
		{0, errInvalidMass, 0},
		{3000000000, nil, 3000000000},
	}
	for _, test := range fuelModuleTests {
		testNew(t, test)
	}
}

func TestFuelModule_GetFuelNeeds(t *testing.T) {
	fuelModuleTests := []fuelModuleTest{
		{1, nil, 0},
		{5, nil, 0},
		{9, nil, 1},
		{3000000006, nil, 1000000000},
	}
	for _, test := range fuelModuleTests {
		testGetFuelNeeds(t, test)
	}
}

func TestFuelModule_GetCompleteFuelNeeds(t *testing.T) {
	fuelModuleTests := []fuelModuleTest{
		{1, nil, 0},
		{5, nil, 0},
		{9, nil, 1},
		{1000, nil, 483},
	}
	for _, test := range fuelModuleTests {
		testGetCompleteFuelNeeds(t, test)
	}
}

func TestNoMass(t *testing.T) {
	newModule := FuelModule{}
	_, err := newModule.GetFuelNeeds()
	if !errors.Is(err, errNoMass) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", errNoMass, err)
	}
}

type fuelModuleTest struct {
	value    int64
	errType  error
	expected int64
}

func testNew(t *testing.T, test fuelModuleTest) {
	newModule, err := New(test.value)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
		return
	}
	if newModule.mass != test.expected {
		t.Errorf("Unexpected mass: expected (%d) got (%d)", test.expected, newModule.mass)
	}
}

func testGetFuelNeeds(t *testing.T, test fuelModuleTest) {
	newModule, err := New(test.value)
	if err != nil {
		t.Errorf("Unexpected error in New: %s", err)
		return
	}
	fuel, err := newModule.GetFuelNeeds()
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if fuel != test.expected {
		t.Errorf("Unexpected mass: expected (%d) got (%d)", test.expected, fuel)
	}
}

func testGetCompleteFuelNeeds(t *testing.T, test fuelModuleTest) {
	newModule, err := New(test.value)
	if err != nil {
		t.Errorf("Unexpected error in New: %s", err)
		return
	}
	fuel, err := newModule.GetCompleteFuelNeeds()
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if fuel != test.expected {
		t.Errorf("Unexpected mass: expected (%d) got (%d)", test.expected, fuel)
	}
}
