// This is lodash golang version. A modern Golang utility library delivering modularity, performance & extras.
package pkg

type DashSlice []interface{}

func NewDashSlice(items ...interface{}) DashSlice {
	dashSlice := append(DashSlice{}, items...)
	return dashSlice
}

func NewDashSliceFromIntArray(items ...int) DashSlice {
	dashSlice := DashSlice{}
	for _, i := range items {
		dashSlice = append(dashSlice, i)
	}

	return dashSlice
}

func (ds DashSlice) Map(iteratee func(interface{}) interface{}) DashSlice {
	result := DashSlice{}

	for _, item := range ds {
		result = append(result, iteratee(item))
	}

	return result
}

type Comparison func(interface{}, interface{}) bool

type Predicate func(interface{}) bool

type Iteratee func(interface{}) interface{}

type IterateeToInt func(interface{}) int

type IterateeToFloat func(interface{}) float64

type IterateeToString func(interface{}) string

type Action func(interface{}, int)

type Reducer func(interface{}, interface{}) interface{}

func Identity(i interface{}) interface{} {
	return i
}

func IdentityInt(i interface{}) int {
	return i.(int)
}

func IdentityFloat64(i interface{}) float64 {
	return i.(float64)
}

func IdentityString(i interface{}) string {
	return i.(string)
}


