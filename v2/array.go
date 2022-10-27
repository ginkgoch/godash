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

func Initial[E any](slice []E) []E {
	length := len(slice)
	result := []E{}

	if length > 1 {
		result = append(result, slice[0:length-1]...)
	}

	return result
}

func Intersection[E any](items1 []E, items2 []E) (intersectedItems []E) {
	intersectedItems = []E{}

	for _, item := range items1 {
		if _, ok := IndexOf(items2, item); ok {
			intersectedItems = append(intersectedItems, item)
		}
	}

	return
}

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

func IntersectionWith[E any](items1 []E, items2 []E, comparison Comparison[E]) (intersectedItems []E) {
	intersectedItems = []E{}

	for _, item := range items1 {
		if _, ok := FindIndexWith(items2, item, comparison); ok {
			intersectedItems = append(intersectedItems, item)
		}
	}

	return
}

func Join[E any](items []E, separator string) string {
	var stringItems []string

	for _, item := range items {
		stringItems = append(stringItems, fmt.Sprintf("%v", item))
	}

	return strings.Join(stringItems, separator)
}

func Last[E any](items []E) (exists bool, lastItem E) {
	length := len(items)
	if length > 0 {
		exists = true
		lastItem = items[length-1]
	}

	return
}

//TODO: Nth
//TODO: Pull
//TODO: PullAll
//TODO: PullAllWith
//TODO: PullAt
//TODO: Remove
//TODO: Slice
//TODO: Tail
//TODO: Take
//TODO: TakeWhile
//TODO: TakeRight
//TODO: TakeRightWhile
//TODO: Union
//TODO: UnionBy
//TODO: UnionWith
//TODO: Uniq
//TODO: UniqBy
//TODO: UniqWith
//TODO: Without
//TODO: Zip
//TODO: ZipWith
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
