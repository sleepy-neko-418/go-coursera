package week2

import (
	"fmt"
	"strconv"
)

func Trunc() {
	var input string

	fmt.Printf("Input a real number: ")
	fmt.Scan(&input)
	number, err := strconv.ParseFloat(input, 64)

	if err == nil {
		output := int(number)
		fmt.Printf("After truncate: %d\n", output)
	} else {
		fmt.Println("Not a valid input!")
	}
}
