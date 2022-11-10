package godash

import (
	"fmt"
	"reflect"
	"strings"
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

	result = append(result, newItems...)

	return result
}

// Creates a new array concatenating array with any additional DashSlices.
func ConcatSlices[E any](slices ...[]E) []E {
	result := []E{}

	for _, slice := range slices {
		result = Concat(result, slice)
	}

	return result
}

// Creates an array of array values not included in the other given arrays using SameValueZero for equality comparisons.
// The order and references of result values are determined by the first array.
func Difference[E any](items []E, itemsToCompare []E) []E {
	var result = []E{}

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
func DifferenceBy[E any](items []E, itemsToCompare []E, iteratee Iteratee[E, E]) []E {
	itemsNew := Map(items, iteratee)
	itemsToCompareNew := Map(itemsToCompare, iteratee)

	var result = []E{}

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
func DifferenceWith[E any](items []E, itemsToCompare []E,
	comparison Comparison[E]) []E {
	var result = []E{}

	for i, item := range items {
		if _, ok := FindIndexWith(itemsToCompare, item, comparison); !ok {
			result = append(result, items[i])
		}
	}

	return result
}

// Creates a slice of array with n elements dropped from the beginning.
func Drop[E any](items []E, count int) []E {
	result := []E{}

	for i := count; i < len(items); i++ {
		result = append(result, items[i])
	}

	return result
}

// Creates a slice of array with n elements dropped from the end.
func DropRight[E any](items []E, count int) []E {
	result := []E{}
	length := len(items) - count

	for i := 0; i < length; i++ {
		result = append(result, items[i])
	}

	return result
}

// Creates a slice of array excluding elements dropped from the beginning.
// Elements are dropped until predicate returns falsy.
// The predicate is invoked with two arguments: (value, index).
func DropWhile[E any](items []E, predicate Predicate[E]) []E {
	result := []E{}

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

// This method is like _.find except that it returns the index of the first element predicate returns truthy
// for instead of the element itself.
func IndexOf[E any](items []E, element E) (int, bool) {
	var index = -1
	var ok bool

	for i, el := range items {
		if reflect.DeepEqual(el, element) {
			index = i
			ok = true
			break
		}
	}

	return index, ok
}

// Fills elements of array with value from start up to, but not including end.
func FillInRange[E any](items []E, value E, start int, end int) {
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
func Fill[E any](items []E, fillValue E) {
	FillInRange(items, fillValue, 0, len(items))
}

// Same to IndexOf. The difference is that, this method provides a comparison function to compare programmatically.
func FindIndexWith[E any](items []E, element E, comparison Comparison[E]) (index int, ok bool) {
	index = -1
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
func FindIndex[E any](items []E, predicate Predicate[E]) (int, bool) {
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
func LastIndexOf[E any](items []E, element E) (int, bool) {
	var index = -1
	var ok bool
	length := len(items)

	reversed := make([]E, length)
	copy(reversed, items)
	Reverse(reversed)

	for i, el := range reversed {
		if reflect.DeepEqual(el, element) {
			index = length - i - 1
			ok = true
			break
		}
	}

	return index, ok
}

// This method is like FindIndex except that it iterates over elements of collection from right to left.
func FindLastIndexWith[E any](items []E, element E, comparison Comparison[E]) (int, bool) {
	var index = -1
	var ok bool
	length := len(items)

	reversed := make([]E, length)
	copy(reversed, items)
	Reverse(reversed)

	for i, el := range reversed {
		if comparison(element, el) {
			index = length - 1 - i
			ok = true
			break
		}
	}

	return index, ok
}

// This method is like Find except that it returns the index of the first element
// predicate returns truthy for instead of the element itself.
func FindLastIndex[E any](items []E, predicate Predicate[E]) (int, bool) {
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
func Reverse[E any](items []E) []E {
	length := len(items)

	halfLen := length / 2

	for i := 0; i < halfLen; i++ {
		items[i], items[length-1-i] = items[length-1-i], items[i]
	}

	return items
}

// Gets the first element of slice.
func Head[E any](items []E) *E {
	if len(items) == 0 {
		return nil
	} else {
		return &items[0]
	}
}

// Gets the first element of array.
func First[E any](items []E) *E {
	return Head(items)
}

// This method returns an object composed from key-value pairs.
func FromPairsAny(pairs [][]any) map[any]any {
	result := make(map[any]any)
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

// This method returns an object composed from key-value pairs.
func FromPairs[K DashComparable, V any](pairs []KeyValuePair[K, V]) map[K]V {
	result := make(map[K]V)
	for _, pair := range pairs {
		result[pair.key] = pair.value
	}

	return result
}

func Map[E, V any](slice []E, iteratee func(E) V) []V {
	result := []V{}
	for _, item := range slice {
		result = append(result, iteratee(item))
	}

	return result
}

func Filter[E any](slice []E, predicate Predicate[E]) []E {
	result := []E{}
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// Gets all but the last element of array.
func Initial[E any](slice []E) []E {
	length := len(slice)
	result := []E{}

	if length > 1 {
		result = append(result, slice[0:length-1]...)
	}

	return result
}

// Creates an array of unique values that are included in all given arrays using SameValueZero for equality comparisons.
// The order and references of result values are determined by the first array.
func Intersection[E any](items1 []E, items2 []E) (intersectedItems []E) {
	intersectedItems = []E{}

	for _, item := range items1 {
		if _, ok := IndexOf(items2, item); ok {
			intersectedItems = append(intersectedItems, item)
		}
	}

	return
}

// This method is like Intersection except that it accepts iteratee which is invoked for each element of each arrays
// to generate the criterion by which they're compared. The order and references of result values are determined by the
// first array. The iteratee is invoked with one argument: (value).
func IntersectionBy[E, T any](items1 []E, items2 []E, iteratee Iteratee[E, T]) (intersectedItems []E) {
	intersectedItems = []E{}
	newItems1 := Map(items1, iteratee)
	newItems2 := Map(items2, iteratee)

	for i, item := range newItems1 {
		if _, ok := IndexOf(newItems2, item); ok {
			intersectedItems = append(intersectedItems, items1[i])
		}
	}

	return
}

// This method is like _.intersection except that it accepts comparator which is invoked to compare elements of arrays.
// The order and references of result values are determined by the first array.
// The comparator is invoked with two arguments: (arrVal, othVal).
func IntersectionWith[E any](items1 []E, items2 []E, comparison Comparison[E]) (intersectedItems []E) {
	intersectedItems = []E{}

	for _, item := range items1 {
		if _, ok := FindIndexWith(items2, item, comparison); ok {
			intersectedItems = append(intersectedItems, item)
		}
	}

	return
}

// Converts all elements in array into a string separated by separator.
func Join[E any](items []E, separator string) string {
	var stringItems []string

	for _, item := range items {
		stringItems = append(stringItems, fmt.Sprintf("%v", item))
	}

	return strings.Join(stringItems, separator)
}

// Gets the last element of array.
func Last[E any](items []E) (exists bool, lastItem E) {
	length := len(items)
	if length > 0 {
		exists = true
		lastItem = items[length-1]
	}

	return
}

// Gets the element at index n of array. If n is negative, the nth element from the end is returned.
func Nth[E any](items []E, n int) (exists bool, item E) {
	length := len(items)

	if n >= length || n < -length {
		return
	} else if n >= 0 {
		exists = true
		item = items[n]
		return
	} else {
		exists = true
		item = items[n+length]
		return
	}
}

// Removes all given values from array using SameValueZero for equality comparisons.
func Pull[E DashComparable](items *[]E, values ...E) []E {
	comparison := func(i1 E, i2 E) bool {
		return i1 == i2
	}

	result := PullAllWith(items, values, comparison)
	return result
}

// This method is like Pull except that it accepts an array of values to remove.
func PullAll[E DashComparable](items *[]E, values []E) []E {
	return Pull(items, values...)
}

// This method is like PullAll except that it accepts comparator which is invoked to compare elements of array to values.
// The comparator is invoked with two arguments: (arrVal, othVal).
func PullAllWith[E any](items *[]E, values []E, comparison Comparison[E]) []E {
	result := []E{}

	for _, item := range *items {
		if _, ok := FindIndexWith(values, item, comparison); !ok {
			result = append(result, item)
		}
	}

	*items = result
	return result
}

//Removes elements from array corresponding to indexes and returns an array of removed elements.
func PullAt[E any](items *[]E, indices ...int) (pulled []E) {
	result := []E{}

	for i, item := range *items {
		if _, ok := IndexOf(indices, i); ok {
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
func Remove[E any](items *[]E, predicate Predicate[E]) (removed []E) {
	newItems := []E{}

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
func Slice[E any](items []E, start int, end int) []E {
	if start < 0 {
		start = 0
	}

	length := len(items)
	if end > length {
		end = length
	}

	return items[start:end]
}

// Gets all but the first element of array.
func Tail[E any](items []E) (result []E) {
	if len(items) > 0 {
		result = items[1:]
	} else {
		result = []E{}
	}

	return
}

// Creates a slice of array with n elements taken from the beginning.
func Take[E any](items []E, n int) (results []E) {
	length := len(items)
	if n > length {
		n = length
	}

	results = items[0:n]
	return
}

// Creates a slice of array with elements taken from the beginning. Elements are taken until predicate returns falsy.
// The predicate is invoked with one argument: (value).
func TakeWhile[E any](items []E, predicate Predicate[E]) []E {
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
func TakeRight[E any](items []E, n int) []E {
	length := len(items)
	if n >= length {
		n = length
	}

	from := length - n
	return items[from:]
}

// Creates a slice of array with elements taken from the end. Elements are taken until predicate returns falsy.
// The predicate is invoked with one argument: (value).
func TakeRightWhile[E any](items []E, predicate Predicate[E]) []E {
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
func Union[E DashComparable](slices ...[]E) []E {
	result := ConcatSlices(slices...)
	result = Uniq(result)
	return result
}

// This method is like Uniq except that it accepts iteratee which is invoked for each element in array
// to generate the criterion by which uniqueness is computed. The order of result values is determined
// by the order they occur in the array. The iteratee is invoked with one argument: (value).
func UnionBy[I any, O DashComparable](iteratee Iteratee[I, O], slices ...[]I) []I {
	result := ConcatSlices(slices...)
	result = UniqBy(result, iteratee)
	return result
}

// This method is like Uniq except that it accepts comparator which is invoked to compare elements of array.
// The order of result values is determined by the order they occur in the array.
// The comparator is invoked with two arguments: (arrVal, othVal).
func UnionWith[E any](comparison Comparison[E], slices ...[]E) []E {
	result := ConcatSlices(slices...)
	result = UniqWith(result, comparison)
	return result
}

// Creates a duplicate-free version of an array, using SameValueZero for equality comparisons,
// in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array.
func Uniq[E DashComparable](items []E) []E {
	uniqMarks := make(map[E]bool)
	result := []E{}

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
func UniqBy[I any, O DashComparable](items []I, iteratee Iteratee[I, O]) []I {
	uniqMarks := make(map[O]I)
	result := []I{}

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
func UniqWith[E any](items []E, comparison Comparison[E]) []E {
	result := []E{}

	for _, item := range items {
		if _, found := FindIndexWith(result, item, comparison); !found {
			result = append(result, item)
		}
	}

	return result
}

func xor2[E DashComparable](i1 []E, i2 []E) []E {
	ni1 := Uniq(i1)
	ni2 := Uniq(i2)
	result := []E{}

	for _, item := range ni1 {
		if _, found := IndexOf(ni2, item); found {
			ni2 = Without(ni2, item)
		} else {
			result = append(result, item)
		}
	}

	return append(result, ni2...)
}

// Creates an array excluding all given values using SameValueZero for equality comparisons.
func Without[E DashComparable](items []E, values ...E) []E {
	newItems := make([]E, len(items))
	copy(newItems, items)
	Pull(&newItems, values...)

	return newItems
}

// Creates an array of unique values that is the symmetric difference of the given arrays.
// The order of result values is determined by the order they occur in the arrays.
func Xor[E DashComparable](items ...[]E) []E {
	length := len(items)
	if length == 0 {
		return []E{}
	} else if length == 1 {
		result := append([]E{}, items[0]...)
		return result
	} else {
		result := append([]E{}, items[0]...)
		for i := 1; i < length; i++ {
			result = xor2(result, items[i])
		}

		return result
	}
}

// Creates an array of grouped elements, the first of which contains the first elements of the given arrays,
// the second of which contains the second elements of the given arrays, and so on.
func Zip[E any](slices ...[]E) [][]E {
	maxLength := 0
	for _, slice := range slices {
		if maxLength < len(slice) {
			maxLength = len(slice)
		}
	}

	result := make([][]E, 0)
	for i := 0; i < maxLength; i++ {
		item := []E{}
		for _, slice := range slices {
			if i < len(slice) {
				item = append(item, slice[i])
			} else {
				var v E
				item = append(item, v)
			}
		}

		result = append(result, item)
	}

	return result
}

// This method is like Zip except that it accepts iteratee to specify how grouped values should be combined.
// The iteratee is invoked with the elements of each group: (...group).
func ZipWith[E any, V any](iteratee Iteratee[[]E, V], slices ...[]E) []V {
	zipped := Zip(slices...)

	result := []V{}
	for _, item := range zipped {
		itemResult := iteratee(item)
		result = append(result, itemResult)
	}

	return result
}

// <-- later below ->
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
