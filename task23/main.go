package main

import (
	"errors"
	"fmt"
)

/*
	Удалить i-ый элемент из слайса.
*/

// Удалить элемент без сохранения порядка
func delSliceElement[T any](slice []T, pos int, safeOrder bool) ([]T, error) {
	if pos < 0 || pos >= len(slice) {
		return nil, errors.New("index out of range")
	}
	result := make([]T, len(slice)-1)
	if safeOrder {
		copy(result, append(slice[:pos], slice[pos+1:]...))
	} else {
		slice[pos] = slice[len(slice)-1]
	}

	copy(result, slice)
	return result, nil
}

func main() {
	slice := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice = append(slice, 1)

	fmt.Println(slice)

	slice = []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result, _ := delSliceElement(slice, 5, false)
	fmt.Println(result)

	result, _ = delSliceElement(slice, 5, true)
	fmt.Println(result)
}
