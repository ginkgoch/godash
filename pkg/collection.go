package pkg

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

func ForEach(items DashSlice, action Action) DashSlice {
	for key, item := range items {
		action(item, key)
	}

	return items
}

func Each(items DashSlice, action Action) DashSlice {
	return ForEach(items, action)
}

func ForEachRight(items DashSlice, action Action) DashSlice {
	length := len(items)

	for i := length - 1; i >= 0; i-- {
		action(items[i], i)
	}

	return items
}

func EachRight(items DashSlice, action Action) DashSlice {
	return ForEachRight(items, action)
}

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

func Filter(items DashSlice, predicate Predicate) DashSlice {
	result := DashSlice{}

	for _, item := range items {
		if ok := predicate(item); ok {
			result = append(result, item)
		}
	}

	return result
}

func FindFrom(items DashSlice, predicate Predicate, from int) (interface{}, bool) {
	for _, item := range items[from:] {
		if found := predicate(item); found {
			return item, true
		}
	}

	return nil, false
}

func Find(items DashSlice, predicate Predicate) (interface{}, bool) {
	return FindFrom(items, predicate, 0)
}

func FindLastFrom(items DashSlice, predicate Predicate, from int) (interface{}, bool) {
	length := len(items)
	if from >= length {
		from = length - 1
	}

	for i := from; i >= 0; i-- {
		item := items[i]
		if found := predicate(item); found {
			return item, true
		}
	}

	return nil, false
}

func FindLast(items DashSlice, predicate Predicate) (interface{}, bool) {
	return FindLastFrom(items, predicate, len(items) - 1)
}

func FlatMap(items DashSlice, iteratee Iteratee) DashSlice {
	mappedItems := Map(items, iteratee)
	mappedItems = Flatten(mappedItems)
	return mappedItems
}

func FlatMapDeep(items DashSlice, iteratee Iteratee) DashSlice {
	mappedItems := Map(items, iteratee)
	mappedItems = FlattenDeep(mappedItems)
	return mappedItems
}

func FlatMapDepth(items DashSlice, iteratee Iteratee, depth int) DashSlice {
	mappedItems := Map(items, iteratee)
	mappedItems = FlattenDepth(mappedItems, depth)
	return mappedItems
}

func Map(items DashSlice, iteratee Iteratee) DashSlice {
	mappedItems := DashSlice{}
	for _, item := range items {
		mappedItem := iteratee(item)
		mappedItems = append(mappedItems, mappedItem)
	}

	return mappedItems
}

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

func Includes(items DashSlice, value interface{}) bool {
	_, found := IndexOf(items, value)
	return found
}

func ReduceWithInitial(items DashSlice, reducer Reducer, initial interface{}) interface{} {
	result := initial
	for _, item := range items {
		result = reducer(result, item)
	}

	return result
}

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