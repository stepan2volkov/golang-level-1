package main

import (
	"fmt"
	"math"
	"os"
)

func exitWithError(errorText string) {
	fmt.Printf("ERROR: %s\n", errorText)
	os.Exit(0)
}

func formatNumber(number float64) string {
	format := "%.3f"
	if _, frac := math.Modf(number); frac == 0 {
		format = "%.0f"
	}
	return fmt.Sprintf(format, number)
}

func printResult(a float64, b float64, op string, result float64) {
	fmt.Printf("%s %s %s = %s\n", formatNumber(a), op, formatNumber(b), formatNumber(result))
}

func main() {
	var a, b, result float64
	var op string

	fmt.Print("A: ")
	if _, err := fmt.Scanln(&a); err != nil {
		exitWithError("float number expected for A")
	}

	fmt.Print("Operation: ")
	if _, err := fmt.Scanln(&op); err != nil {
		exitWithError("something got wrong")
	}

	fmt.Print("B: ")
	if _, err := fmt.Scanln(&b); err != nil {
		exitWithError("float number expected for B")
	}

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "%":
		result = math.Mod(a, b)
	case "**":
		fallthrough
	case "^":
		result = math.Pow(a, b)
	case "/":
		if b > 0 {
			result = a / b
		} else {
			exitWithError("division by zero is prohibited")
		}
	default:
		errorText := fmt.Sprintf("operation '%s' is not valid", op)
		exitWithError(errorText)
	}
	printResult(a, b, op, result)
}
