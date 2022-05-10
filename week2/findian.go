package week2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindIan() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a string: ")
	scanner.Scan()
	input := scanner.Text()

	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
