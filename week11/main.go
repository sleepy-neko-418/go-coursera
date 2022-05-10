package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ints := getUserInput()
	part_len := len(ints) / 4

	part_1 := make([]int, part_len)
	copy(part_1, ints[:part_len])

	part_2 := make([]int, part_len)
	copy(part_2, ints[part_len:part_len*2])

	part_3 := make([]int, part_len)
	copy(part_3, ints[part_len*2:part_len*3])

	part_4 := make([]int, len(ints)-part_len*3)
	copy(part_4, ints[part_len*3:])

	needSort := make(chan []int, 10)
	sorted := make(chan []int, 10)

	needSort <- part_1
	needSort <- part_2
	needSort <- part_3
	needSort <- part_4

	go sortArray(needSort, sorted)
	go sortArray(needSort, sorted)
	go sortArray(needSort, sorted)
	go sortArray(needSort, sorted)

	part_1 = <-sorted
	part_2 = <-sorted
	part_3 = <-sorted
	part_4 = <-sorted

	fmt.Println("Main thread recieved from output queue: ", part_1, part_2, part_3, part_4)
	fmt.Println("Start merging...")

	ints = append(part_1, part_2...)
	ints = append(ints, part_3...)
	ints = append(ints, part_4...)
	sort.Ints(ints)

	fmt.Printf("Merged. Result: %v\n", ints)
}

func getUserInput() []int {
	result := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a series of integers: ")
	scanner.Scan()
	tokens := strings.Fields(scanner.Text())
	for _, t := range tokens {
		number, err := strconv.Atoi(t)
		if err != nil {
			fmt.Printf("%s is not a valid number, skip!\n", t)
		} else {
			result = append(result, number)
		}
	}

	return result
}

func sortArray(in chan []int, out chan []int) {
	arr := <-in
	fmt.Printf("Recieved from input queue: %v\n", arr)
	fmt.Println("Sorting...")
	sort.Ints(arr)
	fmt.Printf("Sorted. Sending to output queue: %v\n\n", arr)
	out <- arr
}
