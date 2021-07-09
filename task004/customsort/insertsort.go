package customsort

// Сортировка вставками на месте
func SortInPlace(array []int) {
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
