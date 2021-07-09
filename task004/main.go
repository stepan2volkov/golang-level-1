package main

import (
	"fmt"
	"task004/customsort"
)

func main() {
	intNumbers := []int{10, 8, 6, 4, 2}
	customsort.SortInPlace(intNumbers)
	_, err := fmt.Println(intNumbers)
	if err != nil {
		panic(err)
	}
}
