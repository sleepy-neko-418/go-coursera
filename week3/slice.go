package week3

import (
	"fmt"
	"strconv"
)

func Slice() {
	mySlice := make([]int, 0, 3)
	input := ""

	for input != "X" {
		fmt.Print("Input an integer: ")
		fmt.Scan(&input)

		if input == "X" {
			continue
		}

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: not a valid integer!")
			continue
		}

		index := binarySearch(mySlice, number)

		if index == len(mySlice) {
			mySlice = append(mySlice, number)
		} else {
			mySlice = append(mySlice[:index+1], mySlice[index:]...)
			mySlice[index] = number
		}
		fmt.Println(mySlice)
	}
}

func binarySearch(mySlice []int, number int) int {
	low := 0
	high := len(mySlice) - 1

	for low <= high {
		mid := (low + high) / 2
		switch {
		case mySlice[mid] > number:
			high = mid - 1
		case mySlice[mid] < number:
			low = mid + 1
		case mySlice[mid] == number:
			if mid == high || mySlice[mid+1] > number {
				return mid
			}
			low = mid + 1
		}
	}
	return low
}
