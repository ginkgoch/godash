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

func ForEach(items DashSlice, action Action) {
	for key, item := range items {
		action(item, key)
	}
}

func ForEachRight(items DashSlice, action Action) {
	length := len(items)

	for i := length - 1; i >= 0; i-- {
		action(items[i], i)
	}
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