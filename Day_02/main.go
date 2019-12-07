package main

import (
	"fmt"
	"github.com/IoHughto/advent_of_code_2019/Day_02/program"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	rawData, err := ioutil.ReadFile("data/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := getProgramData(string(rawData))
	if err != nil {
		log.Fatal(err)
	}
	noun, verb, err := findTargetInput(data, 19690720)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Noun: %d\nVerb: %d\n", noun, verb)
}

func getProgramData(rawData string) ([]int, error) {
	var data []int

	stringData := strings.Split(rawData, ",")
	for _, stringDatum := range stringData {
		if stringDatum == "" {
			continue
		}
		instruction, err := strconv.Atoi(stringDatum)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", errIntConversion, err)
		}
		data = append(data, instruction)
	}

	return data, nil
}

func findTargetInput(data []int, target int) (int, int, error) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			newProgram := program.IntCode{}
			newProgram.Load(data)
			replaceData := []program.ReplaceData{{1, i}, {2, j}}
			newProgram, err := newProgram.Replace(replaceData)
			if err != nil {
				return 0, 0, fmt.Errorf("%w: %s", errReplaceError, err)
			}
			result, err := newProgram.Run()
			if err != nil {
				return 0, 0, err
			}
			if result == target {
				return i, j, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("%w: %d", errTargetNotFound, target)
}

// sentinel errors
var errIntConversion = fmt.Errorf("cannot convert to string")
var errReplaceError = fmt.Errorf("could not replace inputs")
var errTargetNotFound = fmt.Errorf("target output not found")
