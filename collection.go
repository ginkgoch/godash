package godash

import (
	"math/rand"
	"sort"
)

// Creates an object composed of keys generated from the results of running each element of collection thru iteratee.
// The corresponding value of each key is the number of times the key was returned by iteratee.
// The iteratee is invoked with one argument: (value).
func CountBy(items DashSlice, iteratee Iteratee) map[interface{}]int {
	counts := map[interface{}]int{}
	for _, item := range items {
		itemTmp := iteratee(item)
		count, found := counts[itemTmp]
		if !found {
			count = 1
		} else {
			count++
		}

		counts[itemTmp] = count
	}

	return counts
}

// Iterates over elements of collection and invokes iteratee for each element.
// The iteratee is invoked with three arguments: (value, index|key, collection).
// Iteratee functions may exit iteration early by explicitly returning false.
func ForEach(items DashSlice, action Action) DashSlice {
	for key, item := range items {
		action(item, key)
	}

	return items
}

// Iterates over elements of collection and invokes iteratee for each element.
// The iteratee is invoked with three arguments: (value, index|key, collection).
// Iteratee functions may exit iteration early by explicitly returning false.
func Each(items DashSlice, action Action) DashSlice {
	return ForEach(items, action)
}

// This method is like ForEach except that it iterates over elements of collection from right to left.
func ForEachRight(items DashSlice, action Action) DashSlice {
	length := len(items)

	for i := length - 1; i >= 0; i-- {
		action(items[i], i)
	}

	return items
}

// This method is like ForEach except that it iterates over elements of collection from right to left.
func EachRight(items DashSlice, action Action) DashSlice {
	return ForEachRight(items, action)
}

// Checks if predicate returns truthy for all elements of collection. Iteration is stopped once predicate returns falsy.
// The predicate is invoked with one argument: (value).
func Every(items DashSlice, predicate Predicate) bool {
	result := true

	for _, item := range items {
		if ok := predicate(item); !ok {
			result = false
			break
		}
	}

	return result
}

// Iterates over elements of collection, returning an array of all elements predicate returns truthy for.
// The predicate is invoked with one argument: (value).
func Filter(items DashSlice, predicate Predicate) DashSlice {
	result := DashSlice{}

	for _, item := range items {
		if ok := predicate(item); ok {
			result = append(result, item)
		}
	}

	return result
}

// Iterates over elements of collection from start index, returning the first element predicate returns truthy for.
// The predicate is invoked with one argument: (value).
func FindFrom(items DashSlice, predicate Predicate, start int) (interface{}, bool) {
	for _, item := range items[start:] {
		if found := predicate(item); found {
			return item, true
		}
	}

	return nil, false
}

// Iterates over elements of collection, returning the first element predicate returns truthy for.
// The predicate is invoked with one argument: (value).
func Find(items DashSlice, predicate Predicate) (interface{}, bool) {
	return FindFrom(items, predicate, 0)
}

// This method is like FindFrom except that it iterates over elements of collection from right to left.
func FindLastFrom(items DashSlice, predicate Predicate, start int) (interface{}, bool) {
	length := len(items)
	if start >= length {
		start = length - 1
	}

	for i := start; i >= 0; i-- {
		item := items[i]
		if found := predicate(item); found {
			return item, true
		}
	}

	return nil, false
}

// This method is like Find except that it iterates over elements of collection from right to left.
func FindLast(items DashSlice, predicate Predicate) (interface{}, bool) {
	return FindLastFrom(items, predicate, len(items)-1)
}

// Creates a flattened array of values by running each element in collection thru iteratee and
// flattening the mapped results. The iteratee is invoked with one argument: (value).
func FlatMap(items DashSlice, iteratee Iteratee) DashSlice {
	mappedItems := Map(items, iteratee)
	mappedItems = Flatten(mappedItems)
	return mappedItems
}

// This method is like FlatMap except that it recursively flattens the mapped results.
func FlatMapDeep(items DashSlice, iteratee Iteratee) DashSlice {
	mappedItems := Map(items, iteratee)
	mappedItems = FlattenDeep(mappedItems)
	return mappedItems
}

// This method is like FlatMap except that it recursively flattens the mapped results up to depth times.
func FlatMapDepth(items DashSlice, iteratee Iteratee, depth int) DashSlice {
	mappedItems := Map(items, iteratee)
	mappedItems = FlattenDepth(mappedItems, depth)
	return mappedItems
}

// Creates an array of values by running each element in collection thru iteratee.
// The iteratee is invoked with one argument: (value).
func Map(items DashSlice, iteratee Iteratee) DashSlice {
	mappedItems := DashSlice{}
	for _, item := range items {
		mappedItem := iteratee(item)
		mappedItems = append(mappedItems, mappedItem)
	}

	return mappedItems
}

// Creates an object composed of keys generated from the results of running each element of collection thru iteratee.
// The order of grouped values is determined by the order they occur in collection.
// The corresponding value of each key is an array of elements responsible for generating the key.
// The iteratee is invoked with one argument: (value).
func GroupBy(items DashSlice, iteratee Iteratee) map[interface{}]DashSlice {
	result := map[interface{}]DashSlice{}

	for _, item := range items {
		key := iteratee(item)

		el, found := result[key]
		if found {
			result[key] = append(el, item)
		} else {
			result[key] = append(DashSlice{}, item)
		}
	}

	return result
}

// Checks if value is in collection. If collection is a string, it's checked for a substring of value,
// otherwise SameValueZero is used for equality comparisons. If fromIndex is negative, it's used as the offset from
// the end of collection.
func Includes(items DashSlice, value interface{}) bool {
	_, found := IndexOf(items, value)
	return found
}

// Reduces collection to a value which is the accumulated result of running each element in collection thru iteratee,
// where each successive invocation is supplied the return value of the previous. If accumulator is not given,
// the first element of collection is used as the initial value. The reducer is invoked with two arguments:
// (accumulator, value).
func ReduceWithInitial(items DashSlice, reducer Reducer, initial interface{}) interface{} {
	result := initial
	for _, item := range items {
		result = reducer(result, item)
	}

	return result
}

// Reduces collection to a value which is the accumulated result of running each element in collection thru iteratee,
// where each successive invocation is supplied the return value of the previous. If accumulator is not given,
// the first element of collection is used as the initial value. The reducer is invoked with two arguments:
// (accumulator, value).
func Reduce(items DashSlice, reducer Reducer) interface{} {
	length := len(items)
	if length == 0 {
		return nil
	} else if length == 1 {
		return items[0]
	} else {
		return ReduceWithInitial(items[1:], reducer, items[0])
	}
}

// This method is like ReduceWithInitial except that it iterates over elements of collection from right to left.
func ReduceRightWithInitial(items DashSlice, reducer Reducer, initial interface{}) interface{} {
	result := initial
	for i := len(items) - 1; i >= 0; i-- {
		item := items[i]
		result = reducer(result, item)
	}

	return result
}

// This method is like Reduce except that it iterates over elements of collection from right to left.
func ReduceRight(items DashSlice, reducer Reducer) interface{} {
	length := len(items)
	if length == 0 {
		return nil
	} else if length == 1 {
		return items[0]
	} else {
		return ReduceRightWithInitial(items[0:length-1], reducer, items[length-1])
	}
}

// The opposite of Filter; this method returns the elements of collection that predicate does not return truthy for.
func Reject(items DashSlice, predicate Predicate) DashSlice {
	result := DashSlice{}
	for _, item := range items {
		if ok := predicate(item); !ok {
			result = append(result, item)
		}
	}

	return result
}

// Gets a random element from collection.
func Sample(items DashSlice) interface{} {
	n := rand.Intn(len(items))
	return items[n]
}

// Gets n random elements at unique keys from collection up to the size of collection.
func SampleSize(items DashSlice, n int) DashSlice {
	result := DashSlice{}

	for i := 0; i < n; i++ {
		length := len(items)
		if length == 0 {
			break
		}
		r := rand.Intn(len(items))
		result = append(result, items[r])
		items = append(items[:r], items[r+1:]...)
	}

	return result
}

// Creates an array of shuffled values.
func Shuffle(items DashSlice) DashSlice {
	return SampleSize(items, len(items))
}

// Gets the size of collection.
func Size(items DashSlice) int {
	return len(items)
}

// Checks if predicate returns truthy for any element of collection. Iteration is stopped once predicate returns truthy.
// The predicate is invoked with one argument: (value).
func Some(items DashSlice, predicate Predicate) bool {
	found := false

	for _, item := range items {
		if ok := predicate(item); ok {
			found = true
			break
		}
	}

	return found
}

// Creates an array of elements, sorted in ascending order by the results of running each element in a collection
// thru each iteratee. This method performs a stable sort, that is, it preserves the original sort order of equal
// elements. The iteratee is invoked with one argument: (value).
func SortByInt(items DashSlice, iteratee IterateeToInt) DashSlice {
	result := append(DashSlice{}, items...)
	sort.SliceStable(result, func(i, j int) bool {
		el1, el2 := iteratee(result[i]), iteratee(result[j])
		return el1 < el2
	})

	return result
}

// Creates an array of elements, sorted in ascending order by the results of running each element in a collection
// thru each iteratee. This method performs a stable sort, that is, it preserves the original sort order of equal
// elements. The iteratee is invoked with one argument: (value).
func SortByFloat64(items DashSlice, iteratee IterateeToFloat) DashSlice {
	result := append(DashSlice{}, items...)
	sort.SliceStable(result, func(i, j int) bool {
		el1, el2 := iteratee(result[i]), iteratee(result[j])
		return el1 < el2
	})

	return result
}

// Creates an array of elements, sorted in ascending order by the results of running each element in a collection
// thru each iteratee. This method performs a stable sort, that is, it preserves the original sort order of equal
// elements. The iteratee is invoked with one argument: (value).
func SortByString(items DashSlice, iteratee IterateeToString) DashSlice {
	result := append(DashSlice{}, items...)
	sort.SliceStable(result, func(i, j int) bool {
		el1, el2 := iteratee(result[i]), iteratee(result[j])
		return el1 < el2
	})

	return result
}
