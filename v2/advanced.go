package godash

type Number interface {
	int | int16 | int32 | int64 | int8 | float32 | float64 | uint | uint16 | uint32 | uint64 | uint8
}

type DashAny interface {
	Number | string | interface{}
}

type Iteratee[E any, V any] func(E) V

type Predicate[E any] func(E) bool

type Comparison[E any] func(E, E) bool
