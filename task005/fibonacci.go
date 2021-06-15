package main

import (
	"fmt"
	"os"
)

var cacheCalls uint = 0
var calcCalls uint = 0

// Удобно для сбора статистики создать алиас и методы для него
type FibonacciCache map[uint]uint

func (cache FibonacciCache) put(n, number uint) {
	cache[n] = number
}

func (cache FibonacciCache) get(n uint) (uint, bool) {
	value, ok := cache[n]
	if ok {
		// Фиксируем только те вызовы, которые были полезны
		cacheCalls++
	}
	return value, ok
}

var fibonacciNumbers = FibonacciCache{}

func fibonacciRecursive(n uint) (number uint) {
	number, ok := fibonacciNumbers.get(n)
	if ok {
		return
	}
	calcCalls++
	defer func() {
		fibonacciNumbers.put(n, number)
	}()
	if n < 2 {
		number = n
		return
	}
	number = fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
	return
}

func main() {
	var fibonacciNum uint

	fmt.Print("Enter Fibonacci Num: ")
	if _, err := fmt.Scanln(&fibonacciNum); err != nil {
		fmt.Println("You should use integer number")
		os.Exit(1)
	}

	fmt.Println("Result:", fibonacciRecursive(fibonacciNum))

	// Печать статистики
	fmt.Printf("Cache: %d, Calculation: %d\n", cacheCalls, calcCalls)
}
