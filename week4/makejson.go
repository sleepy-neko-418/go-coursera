package week4

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func MakeJson() {
	scanner := bufio.NewScanner(os.Stdin)
	person := make(map[string]string)

	fmt.Print("Enter name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter address: ")
	scanner.Scan()
	address := scanner.Text()

	person["name"] = name
	person["address"] = address

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error: cannot encode data to json!")
	} else {
		fmt.Println(string(jsonData))
	}
}
