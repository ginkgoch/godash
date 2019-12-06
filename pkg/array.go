package pkg

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

func Reverse(items DashSlice) DashSlice {
	length := len(items)

	halfLen := length / 2

	for i := 0; i < halfLen; i++ {
		items[i], items[length-1-i] = items[length-1-i], items[i]
	}

	return items
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

func Initial(items DashSlice) DashSlice {
	length := len(items)
	result := DashSlice{}

	if length > 1 {
		result = append(result, items[0:length-1]...)
	}

	return result
}

func Intersection(items1 DashSlice, items2 DashSlice) DashSlice {
	result := DashSlice{}

	for _, item := range items1 {
		if _, ok := IndexOf(items2, item); ok {
			result = append(result, item)
		}
	}

	return result
}

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

func IntersectionWith(items1 DashSlice, items2 DashSlice, comparison Comparison) DashSlice {
	result := DashSlice{}

	for _, item := range items1 {
		if _, ok := FindIndexWith(items2, item, comparison); ok {
			result = append(result, item)
		}
	}

	return result
}

func Join(items DashSlice, separator string) string {
	var strs []string

	for _, item := range items {
		strs = append(strs, fmt.Sprintf("%v", item))
	}

	return strings.Join(strs, separator)
}

func Last(items DashSlice) interface{} {
	var result interface{}

	length := len(items)
	if length > 0 {
		result = items[length-1]
	}

	return result
}

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

func Pull(items *DashSlice, values ...interface{}) DashSlice {
	result := PullAllWith(items, values, func(i1 interface{}, i2 interface{}) bool {
		return i1 == i2
	})

	return result
}

func PullAll(items *DashSlice, values DashSlice) DashSlice {
	result := Pull(items, values...)
	return result
}

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

func Slice(items DashSlice, from int, to int) DashSlice {
	return items[from:to]
}

func Tail(items DashSlice) DashSlice {
	if len(items) > 0 {
		return items[1:]
	}

	return DashSlice{}
}

func Take(items DashSlice, n int) DashSlice {
	length := len(items)
	if n >= length {
		n = length
	}

	return items[0:n]
}

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

func TakeRight(items DashSlice, n int) DashSlice {
	length := len(items)
	if n >= length {
		n = length
	}

	from := length - n
	return items[from:]
}

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

//TODO: sortedIndex
//TODO: sortedIndexBy
//TODO: sortedIndexOf
//TODO: sortedLastIndex
//TODO: sortedLastIndexBy
//TODO: sortedLastIndexOf
//TODO: sortedUniq
//TODO: sortedUniqBy
