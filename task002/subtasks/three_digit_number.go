package subtasks

import "fmt"

func ThreeDigitNumber() {
	var number int16

	fmt.Print("Enter three digit number: ")
	fmt.Scanln(&number)

	if number < 100 || number > 999 {
		fmt.Println("The number should contain three digit.")
		return
	}

	units := number % 10
	number = (number - units) / 10
	tens := number % 10
	hundreds := (number - tens) / 10

	fmt.Printf("Hundreds: %d, tens: %d, units: %d\n", hundreds, tens, units)
}
