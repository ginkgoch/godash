package godash

import (
	"fmt"
	"reflect"
	"strings"
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

// Creates a new array concatenating array with any additional DashSlices.
func ConcatSlices(slices ...DashSlice) DashSlice {
	result := DashSlice{}
	for _, slice := range slices {
		result = Concat(result, slice...)
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

// Creates a slice of array excluding elements dropped from the beginning.
// Elements are dropped until predicate returns falsy.
// The predicate is invoked with two arguments: (value, index).
func DropWhile(items DashSlice, predicate Predicate) DashSlice {
	result := DashSlice{}

	started := false
	for _, item := range items {
		if !started && !predicate(item) {
			started = true
		}

		if started {
			result = append(result, item)
		}
	}

	return result
}

// Fills elements of array with value from start up to, but not including end.
func FillInRange(items DashSlice, value interface{}, start int, end int) {
	if start < 0 {
		start = 0
	}

	length := len(items)
	if end >= length {
		end = length
	}

	for i := start; i < end; i++ {
		items[i] = value
	}
}

// Fills elements of array with value.
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

// This method is like Find except that it returns the index of the first element predicate
// returns truthy for instead of the element itself.
func FindIndex(items DashSlice, predicate Predicate) (int, bool) {
	index := -1
	ok := false

	for i, el := range items {
		if predicate(el) {
			index = i
			ok = true
			break
		}
	}

	return index, ok
}

// This method is like IndexOf except that it iterates over elements of array from right to left.
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

// This method is like FindIndex except that it iterates over elements of collection from right to left.
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

// This method is like Find except that it returns the index of the first element
// predicate returns truthy for instead of the element itself.
func FindLastIndex(items DashSlice, predicate Predicate) (int, bool) {
	index := -1
	ok := false

	for i, el := range items {
		if predicate(el) {
			index = i
			ok = true
			break
		}
	}

	return index, ok
}

// Reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
func Reverse(items DashSlice) DashSlice {
	length := len(items)

	halfLen := length / 2

	for i := 0; i < halfLen; i++ {
		items[i], items[length-1-i] = items[length-1-i], items[i]
	}

	return items
}

// Gets the first element of slice.
func Head(items DashSlice) interface{} {
	if len(items) == 0 {
		return nil
	} else {
		return items[0]
	}
}

// Gets the first element of array.
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

// Flattens array a single level deep.
func Flatten(items DashSlice) DashSlice {
	result := DashSlice{}

	for _, item := range items {
		els := flattenRecursive(item, 0, 1)
		result = append(result, els...)
	}

	return result
}

// Recursively flattens array.
func FlattenDeep(items DashSlice) DashSlice {
	result := DashSlice{}

	for _, item := range items {
		els := flattenRecursive(item, 0, -1)
		result = append(result, els...)
	}

	return result
}

// Recursively flatten array up to depth times.
func FlattenDepth(items DashSlice, depth int) DashSlice {
	result := DashSlice{}

	for _, item := range items {
		els := flattenRecursive(item, 0, depth)
		result = append(result, els...)
	}

	return result
}

// This method returns an object composed from key-value pairs.
func FromPairs(pairs []DashSlice) map[interface{}]interface{} {
	result := make(map[interface{}]interface{})
	for _, pair := range pairs {
		if len(pair) == 0 {
			continue
		} else if len(pair) == 1 {
			result[pair[0]] = nil
		} else {
			result[pair[0]] = pair[1]
		}
	}

	return result
}

// Gets all but the last element of array.
func Initial(items DashSlice) DashSlice {
	length := len(items)
	result := DashSlice{}

	if length > 1 {
		result = append(result, items[0:length-1]...)
	}

	return result
}

// Creates an array of unique values that are included in all given arrays using SameValueZero for equality comparisons.
// The order and references of result values are determined by the first array.
func Intersection(items1 DashSlice, items2 DashSlice) DashSlice {
	result := DashSlice{}

	for _, item := range items1 {
		if _, ok := IndexOf(items2, item); ok {
			result = append(result, item)
		}
	}

	return result
}

// This method is like Intersection except that it accepts iteratee which is invoked for each element of each arrays
// to generate the criterion by which they're compared. The order and references of result values are determined by the
// first array. The iteratee is invoked with one argument: (value).
func IntersectionBy(items1 DashSlice, items2 DashSlice, iteratee Iteratee) DashSlice {
	result := DashSlice{}
	newItems1 := items1.Map(iteratee)
	newItems2 := items2.Map(iteratee)

	for i, item := range newItems1 {
		if _, ok := IndexOf(newItems2, item); ok {
			result = append(result, items1[i])
		}
	}

	return result
}

// This method is like _.intersection except that it accepts comparator which is invoked to compare elements of arrays.
// The order and references of result values are determined by the first array.
// The comparator is invoked with two arguments: (arrVal, othVal).
func IntersectionWith(items1 DashSlice, items2 DashSlice, comparison Comparison) DashSlice {
	result := DashSlice{}

	for _, item := range items1 {
		if _, ok := FindIndexWith(items2, item, comparison); ok {
			result = append(result, item)
		}
	}

	return result
}

// Converts all elements in array into a string separated by separator.
func Join(items DashSlice, separator string) string {
	var stringItems []string

	for _, item := range items {
		stringItems = append(stringItems, fmt.Sprintf("%v", item))
	}

	return strings.Join(stringItems, separator)
}

// Gets the last element of array.
func Last(items DashSlice) interface{} {
	var result interface{}

	length := len(items)
	if length > 0 {
		result = items[length-1]
	}

	return result
}

// Gets the element at index n of array. If n is negative, the nth element from the end is returned.
func Nth(items DashSlice, n int) interface{} {
	length := len(items)
	if n >= length || n < -length {
		return nil
	} else if n >= 0 {
		return items[n]
	} else {
		return items[length+n]
	}
}

// Removes all given values from array using SameValueZero for equality comparisons.
func Pull(items *DashSlice, values ...interface{}) DashSlice {
	result := PullAllWith(items, values, func(i1 interface{}, i2 interface{}) bool {
		return i1 == i2
	})

	return result
}

// This method is like Pull except that it accepts an array of values to remove.
func PullAll(items *DashSlice, values DashSlice) DashSlice {
	result := Pull(items, values...)
	return result
}

// This method is like PullAll except that it accepts comparator which is invoked to compare elements of array to values.
// The comparator is invoked with two arguments: (arrVal, othVal).
func PullAllWith(items *DashSlice, values DashSlice, comparison Comparison) DashSlice {
	result := DashSlice{}

	for _, item := range *items {
		if _, ok := FindIndexWith(values, item, comparison); !ok {
			result = append(result, item)
		}
	}

	*items = result
	return result
}

//Removes elements from array corresponding to indexes and returns an array of removed elements.
func PullAt(items *DashSlice, indices ...int) DashSlice {
	pulled := DashSlice{}
	result := DashSlice{}
	indicesSlice := NewDashSliceFromIntArray(indices...)

	for i, item := range *items {
		if _, ok := IndexOf(indicesSlice, i); ok {
			pulled = append(pulled, item)
		} else {
			result = append(result, item)
		}
	}

	*items = result
	return pulled
}

// Removes all elements from array that predicate returns truthy for and returns an array of the removed elements.
// The predicate is invoked with two arguments: (value, index).
func Remove(items *DashSlice, predicate Predicate) DashSlice {
	removed := DashSlice{}
	newItems := DashSlice{}

	for _, item := range *items {
		if predicate(item) {
			removed = append(removed, item)
		} else {
			newItems = append(newItems, item)
		}
	}

	*items = newItems
	return removed
}

// Creates a slice of array from start up to, but not including, end.
func Slice(items DashSlice, start int, end int) DashSlice {
	return items[start:end]
}

// Gets all but the first element of array.
func Tail(items DashSlice) DashSlice {
	if len(items) > 0 {
		return items[1:]
	}

	return DashSlice{}
}

// Creates a slice of array with n elements taken from the beginning.
func Take(items DashSlice, n int) DashSlice {
	length := len(items)
	if n >= length {
		n = length
	}

	return items[0:n]
}

// Creates a slice of array with elements taken from the beginning. Elements are taken until predicate returns falsy.
// The predicate is invoked with one argument: (value).
func TakeWhile(items DashSlice, predicate Predicate) DashSlice {
	var to int

	for i, item := range items {
		if !predicate(item) {
			break
		} else {
			to = i + 1
		}
	}

	return Take(items, to)
}

// Creates a slice of array with n elements taken from the end.
func TakeRight(items DashSlice, n int) DashSlice {
	length := len(items)
	if n >= length {
		n = length
	}

	from := length - n
	return items[from:]
}

// Creates a slice of array with elements taken from the end. Elements are taken until predicate returns falsy.
// The predicate is invoked with one argument: (value).
func TakeRightWhile(items DashSlice, predicate Predicate) DashSlice {
	from := len(items)

	for i := from - 1; i >= 0; i-- {
		if !predicate(items[i]) {
			from = i + 1
			break
		}
	}

	return items[from:]
}

// Creates an array of unique values, in order, from all given arrays using SameValueZero for equality comparisons.
func Union(slices ...DashSlice) DashSlice {
	result := ConcatSlices(slices...)
	result = Uniq(result)
	return result
}

// This method is like Uniq except that it accepts iteratee which is invoked for each element in array
// to generate the criterion by which uniqueness is computed. The order of result values is determined
// by the order they occur in the array. The iteratee is invoked with one argument: (value).
func UnionBy(iteratee Iteratee, slices ...DashSlice) DashSlice {
	result := ConcatSlices(slices...)
	result = UniqBy(result, iteratee)
	return result
}

// This method is like Uniq except that it accepts comparator which is invoked to compare elements of array.
// The order of result values is determined by the order they occur in the array.
// The comparator is invoked with two arguments: (arrVal, othVal).
func UnionWith(comparison Comparison, slices ...DashSlice) DashSlice {
	result := ConcatSlices(slices...)
	result = UniqWith(result, comparison)
	return result
}

// Creates a duplicate-free version of an array, using SameValueZero for equality comparisons,
// in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array.
func Uniq(items DashSlice) DashSlice {
	uniqMarks := make(map[interface{}]bool)
	result := DashSlice{}

	for _, item := range items {
		if !uniqMarks[item] {
			uniqMarks[item] = true
			result = append(result, item)
		}
	}

	return result
}

// This method is like Union except that it accepts iteratee which is invoked for each element of each arrays
// to generate the criterion by which uniqueness is computed. Result values are chosen from the first array
// in which the value occurs. The iteratee is invoked with one argument: (value).
func UniqBy(items DashSlice, iteratee Iteratee) DashSlice {
	uniqMarks := make(map[interface{}]interface{})
	result := DashSlice{}

	for _, item := range items {
		newItem := iteratee(item)
		if _, found := uniqMarks[newItem]; !found {
			uniqMarks[newItem] = item
			result = append(result, item)
		}
	}

	return result
}

// This method is like Uniq except that it accepts comparator which is invoked to compare elements of array.
// The order of result values is determined by the order they occur in the array.
// The comparator is invoked with two arguments: (arrVal, othVal).
func UniqWith(items DashSlice, comparison Comparison) DashSlice {
	result := DashSlice{}

	for _, item := range items {
		if _, found := FindIndexWith(result, item, comparison); !found {
			result = append(result, item)
		}
	}

	return result
}

// Creates an array excluding all given values using SameValueZero for equality comparisons.
func Without(items DashSlice, values ...interface{}) DashSlice {
	newItems := make(DashSlice, len(items))
	copy(newItems, items)
	Pull(&newItems, values...)

	return newItems
}

func xor2(i1 DashSlice, i2 DashSlice) DashSlice {
	ni1 := Uniq(i1)
	ni2 := Uniq(i2)
	result := DashSlice{}

	for _, item := range ni1 {
		if _, found := IndexOf(ni2, item); found {
			ni2 = Without(ni2, item)
		} else {
			result = append(result, item)
		}
	}

	return append(result, ni2...)
}

// Creates an array of unique values that is the symmetric difference of the given arrays.
// The order of result values is determined by the order they occur in the arrays.
func Xor(items ...DashSlice) DashSlice {
	length := len(items)
	if length == 0 {
		return DashSlice{}
	} else if length == 1 {
		result := append(DashSlice{}, items[0]...)
		return result
	} else {
		result := append(DashSlice{}, items[0]...)
		for i := 1; i < length; i++ {
			result = xor2(result, items[i])
		}

		return result
	}
}

// Creates an array of grouped elements, the first of which contains the first elements of the given arrays,
// the second of which contains the second elements of the given arrays, and so on.
func Zip(slices ...DashSlice) []DashSlice {
	maxLength := 0
	for _, slice := range slices {
		if maxLength < len(slice) {
			maxLength = len(slice)
		}
	}

	result := make([]DashSlice, 0)
	for i := 0; i < maxLength; i++ {
		item := DashSlice{}
		for _, slice := range slices {
			if i < len(slice) {
				item = append(item, slice[i])
			} else {
				item = append(item, nil)
			}
		}

		result = append(result, item)
	}

	return result
}

// This method is like Zip except that it accepts iteratee to specify how grouped values should be combined.
// The iteratee is invoked with the elements of each group: (...group).
func ZipWith(iteratee func([]interface{}) interface{}, slices ...DashSlice) DashSlice {
	zipped := Zip(slices...)

	result := DashSlice{}
	for _, item := range zipped {
		itemResult := iteratee(item)
		result = append(result, itemResult)
	}

	return result
}

//TODO: sortedIndex
//TODO: sortedIndexBy
//TODO: sortedIndexOf
//TODO: sortedLastIndex
//TODO: sortedLastIndexBy
//TODO: sortedLastIndexOf
//TODO: sortedUniq
//TODO: sortedUniqBy
//TODO: xorBy
//TODO: xorWith
