package main

import (
	"fmt"
	"math"
)

func main() {
	var area float64
	fmt.Print("Area: ")
	fmt.Scanln(&area)

	if area <= 0 {
		fmt.Printf("Expected area more than 0. Got %f\n", area)
		return
	}

	diameter := 2 * math.Sqrt(area/math.Pi)
	length := diameter * math.Pi

	fmt.Println("Length: ", length)
	fmt.Println("Diameter: ", diameter)
}
