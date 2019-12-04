package pkg

import (
	"reflect"
)

// Creates an array of elements split into groups the length of size. If array can't be split evenly,
// the final chunk will be the remaining elements.
func Chunk(items DashSlice, size int) []DashSlice {
	dashSlices := []DashSlice{}

	for _, item := range items {
		sliceLength := len(dashSlices)
		if sliceLength == 0 || len(dashSlices[sliceLength-1]) == size {
			dashSlices = append(dashSlices, DashSlice{})
			sliceLength++
		}

		dashSlices[sliceLength-1] = append(dashSlices[sliceLength-1], item)
	}

	return dashSlices
}

// Creates an array with all falsy values removed. The values false, 0, "", nil are falsy.
func Compact(items DashSlice) DashSlice {
	dashSlice := DashSlice{}

	for _, item := range items {
		if item != nil && item != false && item != 0 && item != "" {
			dashSlice = append(dashSlice, item)
		}
	}

	return dashSlice
}

// Creates a new array concatenating array with any additional arrays and/or values.
func Concat(items DashSlice, newItems ...interface{}) DashSlice {
	result := append(DashSlice{}, items...)

	for _, newItem := range newItems {
		if reflect.TypeOf(newItem).Kind() == reflect.Slice {
			values := reflect.ValueOf(newItem)
			valuesLength := values.Len()
			for i := 0; i < valuesLength; i++ {
				result = append(result, values.Index(i).Interface())
			}
		} else {
			result = append(result, newItem)
		}
	}

	return result
}

// Creates an array of array values not included in the other given arrays using SameValueZero for equality comparisons.
// The order and references of result values are determined by the first array.
func Difference(items DashSlice, itemsToCompare ...interface{}) DashSlice {
	var result = DashSlice{}

	for _, item := range items {
		if _, ok := FindIndex(itemsToCompare, item); !ok {
			result = append(result, item)
		}
	}

	return result
}

// This method is like _.difference except that it accepts iteratee which is invoked for each element of array
// and values to generate the criterion by which they're compared.
// The order and references of result values are determined by the first array.
// The iteratee is invoked with one argument: (value).
func DifferenceBy(items DashSlice, itemsToCompare DashSlice, iteratee func(interface{}) interface{}) DashSlice {
	itemsNew := items.Map(iteratee)
	itemsToCompareNew := itemsToCompare.Map(iteratee)

	var result = DashSlice{}

	for i, item := range itemsNew {
		if _, ok := FindIndex(itemsToCompareNew, item); !ok {
			result = append(result, items[i])
		}
	}

	return result
}

func DifferenceWith(items DashSlice, itemsToCompare DashSlice,
	comparison Comparison) DashSlice {
	var result = DashSlice{}

	for i, item := range items {
		if _, ok := FindIndexWith(itemsToCompare, item, comparison); !ok {
			result = append(result, items[i])
		}
	}

	return result
}

// This method is like _.find except that it returns the index of the first element predicate returns truthy
// for instead of the element itself.
func FindIndex(items DashSlice, element interface{}) (int, bool) {
	var index = -1
	var ok bool

	for i, el := range items {
		if el == element {
			index = i
			ok = true
			break
		}
	}

	return index, ok
}

// Same to FindIndex. The difference is that, this method provides a comparison function to compare programmatically.
func FindIndexWith(items DashSlice, element interface{}, comparison Comparison) (int, bool) {
	var index = -1
	var ok bool

	for i, el := range items {
		if comparison(element, el) {
			index = i
			ok = true
			break
		}
	}

	return index, ok
}
