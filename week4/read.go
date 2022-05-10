package week4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	firstname string
	lastname  string
}

func Read() {
	people := []Person{}

	stdinScanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter file name: ")
	stdinScanner.Scan()
	fileName := stdinScanner.Text()

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		tokens := strings.Split(line, " ")
		p := Person{
			firstname: tokens[0],
			lastname:  tokens[1],
		}
		people = append(people, p)
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, p := range people {
		fmt.Println("First name: ", p.firstname, " Last name: ", p.lastname)
	}
}
