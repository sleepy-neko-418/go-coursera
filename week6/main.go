package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := getInputToFloat("acceleration")
	v := getInputToFloat("velocity")
	s := getInputToFloat("initial displacement")

	displaceFn := GenDisplaceFn(s, a, v)

	t := getInputToFloat("time")

	fmt.Println("Total displacement: ", displaceFn(t))
}

func getInputToFloat(name string) float64 {
	var input string

	fmt.Print("Enter ", name, ": ")
	fmt.Scan(&input)
	for {
		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Print("Not a valid number. Enter again: ")
			fmt.Scan(&input)
			continue
		}
		return number
	}
}

func GenDisplaceFn(s, a, v float64) func(float64) float64 {
	return func(t float64) float64 {
		return s + 0.5*a*t*t + v*t
	}
}
