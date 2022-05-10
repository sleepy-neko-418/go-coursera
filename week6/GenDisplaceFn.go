package main

import (
	"fmt"
	"math"
)

func main() {

    var a, v0, s0, t float64
	
	fmt.Println("Please enter value for acceleration: ")
	fmt.Scanln(&a)
    fmt.Println("Please enter value for initial velocity:")
	fmt.Scanln(&v0)
    fmt.Println("Please enter value for initial displacement:")
	fmt.Scanln(&s0)
    fmt.Println("Compute the displacement at what time?:")
	fmt.Scanln(&t)

    fn := GenDisplaceFn(a, v0, s0)

    fmt.Println("Displacement: ",fn(t))
}


func GenDisplaceFn (a, v0, s0 float64) func (t float64) float64 {
    fn:=func (t float64) float64 {
        return (a * math.Pow(t, 2) / 2) + v0 * t + s0
    }
        return fn
}
