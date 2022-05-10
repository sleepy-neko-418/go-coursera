package week5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SortIntegers() {
	integers := getUserInput()
	BubbleSort(integers)

	fmt.Println(integers)
}

func getUserInput() []int {
	integers := make([]int, 0, 10)

	fmt.Println("Input at most 10 integers.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	tokens := strings.Fields(input)

	for _, token := range tokens {
		number, err := strconv.Atoi(token)
		if err != nil {
			fmt.Printf("%s is not a valid integer! Skip.\n", token)
		} else {
			integers = append(integers, number)
		}
		if len(integers) >= 10 {
			break
		}
	}

	return integers
}

func BubbleSort(integers []int) {
	for i := len(integers) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if integers[j] > integers[j+1] {
				Swap(integers, j)
			}
		}
	}
}

func Swap(integers []int, index int) {
	integers[index], integers[index+1] = integers[index+1], integers[index]
}
