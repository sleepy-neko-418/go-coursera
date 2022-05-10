package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c *Cow) Eat() {
	fmt.Println("grass")
}

func (c *Cow) Move() {
	fmt.Println("walk")
}

func (c *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (b *Bird) Eat() {
	fmt.Println("worms")
}

func (b *Bird) Move() {
	fmt.Println("fly")
}

func (b *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	name string
}

func (s *Snake) Eat() {
	fmt.Println("mice")
}

func (s *Snake) Move() {
	fmt.Println("slither")
}

func (s *Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)

	for {
		fmt.Print("> ")

		tokens := getUserInput()
		if !isValidInput(tokens) {
			fmt.Println("Invalid request!")
			continue
		}

		switch tokens[0] {
		case "newanimal":
			err := createNewAnimal(animals, tokens[1], tokens[2])
			if err != nil {
				fmt.Println(err)
			}
		case "query":
			queryAnimal(animals, tokens[1], tokens[2])
		}
	}
}

func getUserInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	return strings.Fields(input)
}

func isValidInput(tokens []string) bool {
	validCommands := []string{"newanimal", "query"}
	validAnimalType := []string{"cow", "bird", "snake"}
	validRequests := []string{"eat", "move", "speak"}

	if len(tokens) != 3 {
		return false
	}

	first := strings.ToLower(tokens[0])
	last := strings.ToLower(tokens[2])

	if !contains(validCommands, first) {
		return false
	}

	if tokens[0] == "newanimal" && !contains(validAnimalType, last) {
		return false
	}

	if tokens[0] == "query" && !contains(validRequests, last) {
		return false
	}

	return true
}

func createNewAnimal(animals map[string]Animal, name string, animalType string) error {
	var newAnimal Animal

	_, ok := animals[name]
	if ok {
		return errors.New("An animal with that name already existed. Abort!")
	}

	switch animalType {
	case "cow":
		newAnimal = &Cow{name}
		animals[name] = newAnimal
	case "bird":
		newAnimal = &Bird{name}
		animals[name] = newAnimal
	case "snake":
		newAnimal = &Snake{name}
		animals[name] = newAnimal
	}

	fmt.Println("Created it!")
	return nil
}

func queryAnimal(animals map[string]Animal, name string, request string) {
	animal, ok := animals[name]
	if !ok {
		fmt.Println("No animal with that name exists. Abort!")
		return
	}

	switch request {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func contains(elements []string, element string) bool {
	for _, e := range elements {
		if e == element {
			return true
		}
	}

	return false
}
