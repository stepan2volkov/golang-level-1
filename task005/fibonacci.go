package main

import (
	"fmt"
	"math/big"
	"os"
)

type Fibonacci struct {
	cache map[int64]*big.Int
	// useCache is introduced to make benchmark and see difference in performance
	useCache bool
}

func (f *Fibonacci) Calculate(n int64) (result *big.Int) {
	result, exists := f.cache[n]
	if exists && f.useCache {
		return
	}
	if f.useCache {
		defer func() {
			f.cache[n] = result
		}()
	}
	if n < 2 {
		result = big.NewInt(n)
	} else {
		result = new(big.Int)
		result.Add(f.Calculate(n-1), f.Calculate(n-2))
	}
	return
}

func New(useCache bool) *Fibonacci {
	cache := make(map[int64]*big.Int)
	return &Fibonacci{cache, useCache}
}

func main() {
	f := New(true)
	var fibonacciNum int64

	fmt.Print("Enter Fibonacci Num: ")
	if _, err := fmt.Scanln(&fibonacciNum); err != nil {
		fmt.Println("You should use integer number")
		os.Exit(1)
	}

	fmt.Println("Result:", f.Calculate(fibonacciNum))
}
