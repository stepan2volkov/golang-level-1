package main

import (
	"fmt"
	"math/big"
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

	if _, err := fmt.Print("Enter Fibonacci Num: "); err != nil {
		panic(err)
	}
	if _, err := fmt.Scanln(&fibonacciNum); err != nil {
		panic("You should use integer number")
	}

	fmt.Println("Result:", f.Calculate(fibonacciNum))
}
