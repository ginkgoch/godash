package pkg

import (
	"gotest.tools/assert"
	"math"
	"testing"
)

func TestCountBy(t *testing.T) {
	result := CountBy(DashSlice{1, 1, 2, 2, 2, 3}, func(i interface{}) interface{} {
		return i
	})

	assert.DeepEqual(t, result, map[interface{}]int{1: 2, 2: 3, 3: 1})

	result = CountBy(DashSlice{1.2, 1.4, 2.2, 2.1, 2.4, 3.7}, func(i interface{}) interface{} {
		return int(math.Floor(i.(float64)))
	})

	assert.DeepEqual(t, result, map[interface{}]int{1: 2, 2: 3, 3: 1})
}

func TestForEach(t *testing.T) {
	slice := DashSlice{1, 2, 3, 4}

	times := 0
	ForEach(slice, func(v interface{}, key int) {
		times++

		switch key {
		case 0: assert.Equal(t, slice[key], 1)
		case 1: assert.Equal(t, slice[key], 2)
		case 2: assert.Equal(t, slice[key], 3)
		case 3: assert.Equal(t, slice[key], 4)
		}
	})

	assert.Equal(t, times, len(slice))
}

func TestForEachRight(t *testing.T) {
	slice := DashSlice{1, 2, 3, 4}

	times := 0
	ForEachRight(slice, func(v interface{}, key int) {
		times++

		switch times {
		case 1: assert.Equal(t, slice[key], 4)
		case 2: assert.Equal(t, slice[key], 3)
		case 3: assert.Equal(t, slice[key], 2)
		case 4: assert.Equal(t, slice[key], 1)
		}
	})

	assert.Equal(t, times, len(slice))
}

func TestEvery(t *testing.T) {
	result := Every(DashSlice{2, 4, 6, 8}, func(i interface{}) bool {
		return i.(int) % 2 == 0
	})

	assert.Equal(t, result, true)
}

func TestFilter(t *testing.T) {
	result := Filter(DashSlice{1, 2, 3, 4, 5, 6, 7, 8}, func(i interface{}) bool {
		return i.(int) % 2 == 0
	})

	assert.DeepEqual(t, result, DashSlice{2, 4, 6, 8})

	result = Filter(DashSlice{1, 2, 3, 4, 5, 6, 7, 8}, func(i interface{}) bool {
		return i.(int) % 2 != 0
	})

	assert.DeepEqual(t, result, DashSlice{1, 3, 5, 7})
}

func TestFind(t *testing.T) {
	result, ok := Find(DashSlice{1, 3, 5, 6, 8, 3}, func(i interface{}) bool {
		return i.(int) % 2 == 0
	})

	assert.Equal(t, ok, true)
	assert.Equal(t, result, 6)

	result, ok = Find(DashSlice{1, 3, 5, 6, 8, 3}, func(i interface{}) bool {
		return i.(int) == 10
	})

	assert.Equal(t, ok, false)
	assert.Equal(t, result, nil)
}

func TestFindFrom(t *testing.T) {
	predicate := func(i interface{}) bool {
		return i.(int) > 4
	}

	testSrc := DashSlice{1, 3, 5, 7, 8}

	result, ok := FindFrom(testSrc, predicate, 0)
	assert.Equal(t, result, 5)
	assert.Equal(t, ok, true)

	result, ok = FindFrom(testSrc, predicate, 3)
	assert.Equal(t, result, 7)
	assert.Equal(t, ok, true)

	result, ok = FindFrom(testSrc, predicate, 5)
	assert.Equal(t, result, nil)
	assert.Equal(t, ok, false)
}

func TestFindLast(t *testing.T) {
	predicate := func(i interface{}) bool {
		return i.(int) % 3 == 0
	}

	testSrc := DashSlice{1, 2, 3, 4, 5, 6, 7, 8}

	result, ok := FindLast(testSrc, predicate)
	assert.Equal(t, result, 6)
	assert.Equal(t, ok, true)

	result, ok = FindLastFrom(testSrc, predicate, 20)
	assert.Equal(t, result, 6)
	assert.Equal(t, ok, true)

	result, ok = FindLastFrom(testSrc, predicate, 4)
	assert.Equal(t, result, 3)
	assert.Equal(t, ok, true)

	result, ok = FindLastFrom(testSrc, predicate, 1)
	assert.Equal(t, result, nil)
	assert.Equal(t, ok, false)
}

func TestMap(t *testing.T) {
	result := Map(DashSlice{1, 2, 3}, func(i interface{}) interface{} {
		return i.(int) * 2
	})

	assert.DeepEqual(t, result, DashSlice{2, 4, 6})
}

func TestFlatMap(t *testing.T) {
	result := FlatMap(DashSlice{1, 2, 3}, func(i interface{}) interface{} {
		return []int{i.(int), i.(int)}
	})

	assert.DeepEqual(t, result, DashSlice{1, 1, 2, 2, 3, 3})
}

func TestFlatMapDeep(t *testing.T) {
	result := FlatMapDeep(DashSlice{1, 2, 3}, func(i interface{}) interface{} {
		ii := i.(int)
		return [][]int{{ii, ii}}
	})

	assert.DeepEqual(t, result, DashSlice{1, 1, 2, 2, 3, 3})
}

func TestFlatMapDepth(t *testing.T) {
	result := FlatMapDepth(DashSlice{1, 2, 3}, func(i interface{}) interface{} {
		ii := i.(int)
		return [][]int{{ii, ii}}
	}, 2)

	assert.DeepEqual(t, result, DashSlice{1, 1, 2, 2, 3, 3})
}

func TestGroupBy(t *testing.T) {
	items := DashSlice{6.1, 4.2, 6.3}
	result := GroupBy(items, func(i interface{}) interface{} {
		return math.Floor(i.(float64))
	})

	assert.Equal(t, len(result), 2)

	value1, found := result[6.0]
	assert.Equal(t, found, true)
	assert.Equal(t, len(value1), 2)
	assert.DeepEqual(t, value1, DashSlice{6.1, 6.3})

	value2, found := result[4.0]
	assert.Equal(t, found, true)
	assert.Equal(t, len(value2), 1)
	assert.DeepEqual(t, value2, DashSlice{4.2})

	_, found = result[5.0]
	assert.Equal(t, found, false)
}

func TestIncludes(t *testing.T) {
	items := DashSlice{"a", 1, 0, "c"}

	found := Includes(items, "a")
	assert.Equal(t, found, true)

	found = Includes(items, "c")
	assert.Equal(t, found, true)

	found = Includes(items, 1)
	assert.Equal(t, found, true)

	found = Includes(items, 0)
	assert.Equal(t, found, true)

	found = Includes(items, -1)
	assert.Equal(t, found, false)
}

func TestReduce(t *testing.T) {
	items := DashSlice{1, 2, 3, 4}
	result := Reduce(items, func(i1 interface{}, i2 interface{}) interface{} {
		return i1.(int) + i2.(int)
	})

	assert.Equal(t, result, 10)

	result = ReduceWithInitial(items, func(i1 interface{}, i2 interface{}) interface{} {
		return i1.(int) + i2.(int)
	}, 10)

	assert.Equal(t, result, 20)
}
