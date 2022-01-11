package godash

type Number interface {
	int | int16 | int32 | int64 | int8 | float32 | float64 | uint | uint16 | uint32 | uint64 | uint8
}

type DashAny interface {
	Number | string | interface{}
}

type DashSlice[E any] []E

func Map1[E, V any](slice *DashSlice[E], iteratee Iteratee[E, V]) *DashSlice[V] {
	result := DashSlice[V]{}
	for _, item := range *slice {
		result = append(result, iteratee(item))
	}

	return &result
}

func Map[E, V any](slice *[]E, iteratee func(E) V) *[]V {
	result := []V{}
	for _, item := range *slice {
		result = append(result, iteratee(item))
	}

	return &result
}

func (slice *DashSlice[E]) Filter(predicate Predicate[E]) *DashSlice[E] {
	result := DashSlice[E]{}
	for _, item := range *slice {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return &result
}

type Iteratee[E any, V any] func(E) V

type Predicate[E any] func(E) bool
