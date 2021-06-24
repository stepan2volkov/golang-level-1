package main

import (
	"fmt"
	"task002/subtasks"
)

func main() {
	_, err := fmt.Print("Choose number of subtask:\n1 Circle\n2 Rectangle\n3 Three digit number\n\nYou choice: ")
	if err != nil {
		panic(err)
	}

	var subtaskNumber int
	_, err = fmt.Scanln(&subtaskNumber)
	if err != nil {
		panic(err)
	}

	switch subtaskNumber {
	case 1:
		subtasks.Circle()
	case 2:
		subtasks.Rectangle()
	case 3:
		subtasks.ThreeDigitNumber()
	default:
		if _, err = fmt.Printf("Unknown subtask number: %d\n", subtaskNumber); err != nil {
			panic(err)
		}
	}
}
