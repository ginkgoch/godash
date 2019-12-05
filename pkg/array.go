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
		if _, ok := IndexOf(itemsToCompare, item); !ok {
			result = append(result, item)
		}
	}

	return result
}

// This method is like _.difference except that it accepts iteratee which is invoked for each element of array
// and values to generate the criterion by which they're compared.
// The order and references of result values are determined by the first array.
// The iteratee is invoked with one argument: (value).
func DifferenceBy(items DashSlice, itemsToCompare DashSlice, iteratee Iteratee) DashSlice {
	itemsNew := items.Map(iteratee)
	itemsToCompareNew := itemsToCompare.Map(iteratee)

	var result = DashSlice{}

	for i, item := range itemsNew {
		if _, ok := IndexOf(itemsToCompareNew, item); !ok {
			result = append(result, items[i])
		}
	}

	return result
}

// This method is like _.difference except that it accepts comparator which is invoked to compare elements of array to values.
//The order and references of result values are determined by the first array.
//The comparator is invoked with two arguments: (arrVal, othVal).
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

// Creates a slice of array with n elements dropped from the beginning.
func Drop(items DashSlice, count int) DashSlice {
	result := DashSlice{}

	for i := count; i < len(items); i++ {
		result = append(result, items[i])
	}

	return result
}

// Creates a slice of array with n elements dropped from the end.
func DropRight(items DashSlice, count int) DashSlice {
	result := DashSlice{}
	length := len(items) - count

	for i := 0; i < length; i++ {
		result = append(result, items[i])
	}

	return result
}

func DropWhile(items DashSlice, predict Prediction) DashSlice {
	result := DashSlice{}

	started := false
	for _, item := range items {
		if !started && !predict(item) {
			started = true
		}

		if started {
			result = append(result, item)
		}
	}

	return result
}

func FillInRange(items DashSlice, fillValue interface{}, from int, to int) {
	if from < 0 {
		from = 0
	}

	length := len(items)
	if to >= length {
		to = length
	}

	for i := from; i < to; i++ {
		items[i] = fillValue
	}
}

func Fill(items DashSlice, fillValue interface{}) {
	FillInRange(items, fillValue, 0, len(items))
}

// This method is like _.find except that it returns the index of the first element predicate returns truthy
// for instead of the element itself.
func IndexOf(items DashSlice, element interface{}) (int, bool) {
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

// Same to IndexOf. The difference is that, this method provides a comparison function to compare programmatically.
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

func FindIndex(items DashSlice, predict Prediction) (int, bool) {
	index := -1
	ok := false

	for i, el := range items {
		if predict(el) {
			index = i
			ok = true
			break
		}
	}

	return index, ok
}

func LastIndexOf(items DashSlice, element interface{}) (int, bool) {
	var index = -1
	var ok bool
	length := len(items)

	reversed := make(DashSlice, length)
	copy(reversed, items)
	Reverse(reversed)

	for i, el := range reversed {
		if el == element {
			index = length - i - 1
			ok = true
			break
		}
	}

	return index, ok
}

func FindLastIndexWith(items DashSlice, element interface{}, comparison Comparison) (int, bool) {
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

func FindLastIndex(items DashSlice, predict Prediction) (int, bool) {
	index := -1
	ok := false

	for i, el := range items {
		if predict(el) {
			index = i
			ok = true
			break
		}
	}

	return index, ok
}

func Reverse(items DashSlice) {
	length := len(items)

	halfLen := length / 2

	for i := 0; i < halfLen; i++ {
		items[i], items[length-1-i] = items[length-1-i], items[i]
	}
}

func Head(items DashSlice) interface{} {
	if len(items) == 0 {
		return nil
	} else {
		return items[0]
	}
}

func First(items DashSlice) interface{} {
	return Head(items)
}

func flattenRecursive(item interface{}, currentDepth int, depth int) DashSlice {
	currentDepth++
	result := DashSlice{}
	if reflect.TypeOf(item).Kind() == reflect.Slice {
		sv := reflect.ValueOf(item)
		for i := 0; i < sv.Len(); i++ {
			el := sv.Index(i).Interface()

			if depth == -1 || currentDepth < depth {
				els := flattenRecursive(el, currentDepth, depth)
				result = append(result, els...)
			} else {
				result = append(result, el)
			}
		}
	} else {
		result = append(result, item)
	}

	return result
}

func Flatten(items DashSlice) DashSlice {
	result := DashSlice{}

	for _, item := range items {
		els := flattenRecursive(item, 0, 1)
		result = append(result, els...)
	}

	return result
}

func FlattenDeep(items DashSlice) DashSlice {
	result := DashSlice{}

	for _, item := range items {
		els := flattenRecursive(item, 0, -1)
		result = append(result, els...)
	}

	return result
}

func FlattenDepth(items DashSlice, depth int) DashSlice {
	result := DashSlice{}

	for _, item := range items {
		els := flattenRecursive(item, 0, depth)
		result = append(result, els...)
	}

	return result
}
