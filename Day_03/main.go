package main

import (
	"fmt"
	"github.com/IoHughto/advent_of_code_2019/Day_03/wire"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	rawData, err := ioutil.ReadFile("data/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	wireData, err := getWireData(string(rawData))
	if err != nil {
		log.Fatal(err)
	}

	wires, err := getWires(wireData)
	if err != nil {
		log.Fatal(err)
	}

	err = traverseWires(wires)
	if err != nil {
		log.Fatal(err)
	}

	intersections, err := wires[0].Intersect(wires[1])
	if err != nil {
		log.Fatal(err)
	}

	manhattanIntersection, err := getClosestIntersection(intersections, "manhattan")
	alongIntersection, err := getClosestIntersection(intersections, "along")
	fmt.Printf("%+v\n", manhattanIntersection)
	fmt.Printf("%+v\n", alongIntersection)
}

func getWireData(rawData string) ([]string, error) {
	wireData := strings.Split(rawData, "\n")
	if len(wireData) < 2 {
		return nil, fmt.Errorf("%w: %d", errOneWire, len(wireData))
	}

	return wireData, nil
}

func getWires(wireData []string) ([]wire.Wire, error) {
	var wires []wire.Wire
	for _, wireDatum := range wireData {
		newWire := wire.Wire{}
		err := newWire.Load(strings.Split(wireDatum, ","))
		if err != nil {
			return nil, err
		}
		wires = append(wires, newWire)
	}

	return wires, nil
}

func traverseWires(wires []wire.Wire) error {
	for _, individualWire := range wires {
		err := individualWire.Traverse()
		if err != nil {
			return err
		}
	}

	return nil
}

func getClosestIntersection(intersections []wire.Intersection, lengthType string) (wire.Intersection, error) {
	closest := wire.Intersection{}
	length := 0
	for _, intersection := range intersections {
		thisLength, err := intersection.GetLength(lengthType)
		if err != nil {
			return wire.Intersection{}, err
		}
		if length == 0 {
			closest = intersection
			length = thisLength
		}
		if thisLength < length {
			length = thisLength
			closest = intersection
		}
	}
	return closest, nil
}

var errOneWire = fmt.Errorf("not enough wires were provided")
