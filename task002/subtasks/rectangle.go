package subtasks

import "fmt"

func Rectangle() {
	var a, b float32

	fmt.Print("a: ")
	fmt.Scanln(&a)
	fmt.Print("b: ")
	fmt.Scanln(&b)

	area := a * b
	fmt.Println(area)
}
