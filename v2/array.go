package godash

import (
	"reflect"
)

// Creates an array of elements split into groups the length of size. If array can't be split evenly,
// the final chunk will be the remaining elements.
func Chunk[E any](items []E, size int) [][]E {
	dashSlices := [][]E{}

	for _, item := range items {
		sliceLength := len(dashSlices)
		if sliceLength == 0 || len(dashSlices[sliceLength-1]) == size {
			dashSlices = append(dashSlices, []E{})
			sliceLength++
		}

		dashSlices[sliceLength-1] = append(dashSlices[sliceLength-1], item)
	}

	return dashSlices
}

// Creates an array with all falsy values removed. The values false, 0, "", nil are falsy.
func Compact[E any](items []E) []E {
	dashSlice := []E{}

	for _, item := range items {
		if !reflect.DeepEqual(item, nil) &&
			!reflect.DeepEqual(item, false) &&
			!reflect.DeepEqual(item, 0) &&
			!reflect.DeepEqual(item, "") {
			dashSlice = append(dashSlice, item)
		}
	}

	return dashSlice
}

// Creates a new array concatenating array with any additional arrays and/or values.
func Concat[E any](items []E, newItems []E) []E {
	result := append([]E{}, items...)

	for _, newItem := range newItems {
		result = append(result, newItem)
	}

	return result
}
