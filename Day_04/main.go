package main

import (
	"fmt"
	"github.com/IoHughto/advent_of_code_2019/Day_04/IntChecker"
)

func main() {
	start := 353096
	end := 843212
	counter := 0
	fullCounter := 0
	for i := start; i <= end; i++ {
		intChecker := IntChecker.New(i)
		if intChecker.Check() {
			counter++
		}
		if intChecker.FullCheck() {
			fullCounter++
		}
	}
	fmt.Println(counter)
	fmt.Println(fullCounter)
	//IntChecker.New(777777).FullCheck()
}
