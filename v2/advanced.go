package godash

type Number interface {
	int | int16 | int32 | int64 | int8 | float32 | float64 | uint | uint16 | uint32 | uint64 | uint8
}

type DashAny interface {
	Number | string | interface{}
}

// type DashComparable interface {
// 	Number | string
// }

type DashComparable comparable

type KeyValuePair[K DashComparable, V any] struct {
	key   K
	value V
}

func NewKeyValuePair[K DashComparable, V any](key K, value V) *KeyValuePair[K, V] {
	return &KeyValuePair[K, V]{key: key, value: value}
}

type Iteratee[E any, V any] func(E) V

type Predicate[E any] func(E) bool

type Comparison[E any] func(E, E) bool
