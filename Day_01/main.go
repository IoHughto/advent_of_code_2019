package main

import (
	"fmt"
	"github.com/IoHughto/advent_of_code_2019/Day_01/fuelModule"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("data/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := parseData(bytes)
	if err != nil {
		log.Fatal(err)
	}

	simpleFuelSum, err := getSumOfData(data, getSimpleTotal)
	if err != nil {
		log.Fatal(err)
	}

	completeFuelSum, err := getSumOfData(data, getCompleteTotal)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Simple fuel required: %d\n", simpleFuelSum)
	fmt.Printf("Complete fuel required: %d\n", completeFuelSum)
}

func parseData(bytes []byte) ([]int64, error) {
	var data []int64

	inputData := strings.TrimSuffix(string(bytes), "\n")
	returnStrings := strings.Split(inputData, "\n")
	for _, returnString := range returnStrings {
		number, err := strconv.ParseInt(returnString, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", errInvalidSyntax, err)
		}
		data = append(data, number)
	}
	return data, nil
}

func getSumOfData(data []int64, fuelFunc func(module fuelModule.FuelModule) (int64, error)) (int64, error) {
	var fuelSum int64
	fuelSum = 0
	for _, datum := range data {
		module, err := fuelModule.New(datum)
		if err != nil {
			return 0, err
		}
		fuel, err := fuelFunc(module)
		if err != nil {
			return 0, err
		}
		fuelSum += fuel
	}

	return fuelSum, nil
}

// Sentinel error
var errInvalidSyntax = fmt.Errorf("int64 parse error")

func getSimpleTotal(module fuelModule.FuelModule) (int64, error) {
	return module.GetFuelNeeds()
}

func getCompleteTotal(module fuelModule.FuelModule) (int64, error) {
	return module.GetCompleteFuelNeeds()
}
