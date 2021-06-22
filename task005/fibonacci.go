package main

import (
	"fmt"
	"os"
)

type Fibonacci struct {
	cache map[uint]uint
}

func (f *Fibonacci) Calculate(n uint) (result uint) {
	result, exists := f.cache[n]
	if exists {
		return
	}
	defer func() {
		f.cache[n] = result
	}()
	if n < 2 {
		result = n
	} else {
		result = f.Calculate(n-1) + f.Calculate(n-2)
	}
	return
}

func New() *Fibonacci {
	cache := make(map[uint]uint)
	return &Fibonacci{cache}
}

func main() {
	f := New()
	var fibonacciNum uint

	fmt.Print("Enter Fibonacci Num: ")
	if _, err := fmt.Scanln(&fibonacciNum); err != nil {
		fmt.Println("You should use integer number")
		os.Exit(1)
	}

	fmt.Println("Result:", f.Calculate(fibonacciNum))
}
