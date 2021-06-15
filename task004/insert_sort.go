package main

import (
	"fmt"
)

// Сортировка вставками на месте
func sortInPlace(array []int) {
	for current_index, current_value := range array {
		prev_index := current_index - 1

		for prev_index >= 0 && current_value < array[prev_index] {
			// Передвигаем все элементы слева, которые больше текущего, на одну позицию вправо
			array[prev_index+1] = array[prev_index]
			prev_index--
		}
		array[prev_index+1] = current_value
	}
}

// Сортировка вставками, которая не изменяет первоначальный массив
func sort(array []int) []int {
	newArray := make([]int, len(array))
	copy(newArray, array)
	sortInPlace(newArray)
	return newArray
}

func main() {
	intNumbers := []int{10, 8, 6, 4, 2}

	fmt.Println("INIT:\t\t\t", intNumbers)
	fmt.Println("SORTED:\t\t\t", sort(intNumbers))
	fmt.Println("INIT:\t\t\t", intNumbers)

	sortInPlace(intNumbers)
	fmt.Println("SORTED IN PLACE:\t", intNumbers)
}
