package main

import (
	"fmt"
	"time"
)

// The order of the words are not deterministic
// Because sometimes the second goroutines run first and sometimes the first goroutines run first
func main() {
	go fmt.Println("Hello")
	go fmt.Println("World")

	time.Sleep(1 * time.Second)
}
