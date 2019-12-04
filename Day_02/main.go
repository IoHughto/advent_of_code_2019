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
	newProgram := program.IntCode{}
	newProgram.Load(data)
	newProgram, err = newProgram.Replace([]program.ReplaceData{{1, 12}, {2, 2}})
	if err != nil {
		log.Fatal(err)
	}
	result, err := newProgram.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %d", result)
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

// sentinel errors
var errIntConversion = fmt.Errorf("cannot convert to string")
