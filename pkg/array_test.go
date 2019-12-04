package pkg

import (
	"gotest.tools/assert"
	"math"
	"testing"
)

func TestChunk1(t *testing.T) {
	items := DashSlice{"a", "b", "c", "d"}
	chunked := Chunk(items, 2)

	assert.Equal(t, len(chunked), 2)
	assert.DeepEqual(t, chunked[0], DashSlice{"a", "b"})
	assert.DeepEqual(t, chunked[1], DashSlice{"c", "d"})
}

func TestChunk2(t *testing.T) {
	items := DashSlice{"a", "b", "c", "d"}
	chunked := Chunk(items, 3)

	assert.Equal(t, len(chunked), 2)
	assert.DeepEqual(t, chunked[0], DashSlice{"a", "b", "c"})
	assert.DeepEqual(t, chunked[1], DashSlice{"d"})
}

func TestCompact(t *testing.T) {
	items := DashSlice{"a", "b", false, 0, 1, nil, ""}
	compacted := Compact(items)

	assert.DeepEqual(t, compacted, DashSlice{"a", "b", 1})
}

func TestConcat1(t *testing.T) {
	items := DashSlice{"a", "b", "c", "d"}
	result := Concat(items, "e", "f")

	assert.DeepEqual(t, result, DashSlice{"a", "b", "c", "d", "e", "f"})
}

func TestConcat2(t *testing.T) {
	items := DashSlice{"a", "b", "c", "d"}
	result := Concat(items, []string{"e", "f"})

	assert.DeepEqual(t, result, DashSlice{"a", "b", "c", "d", "e", "f"})
}

func TestFindIndex(t *testing.T) {
	items := DashSlice{"a", "b", "c", "d"}

	findTest := func(el interface{}, expected int, found bool) {
		result, ok := FindIndex(items, el)

		assert.Equal(t, ok, found)
		assert.Equal(t, result, expected)
	}

	findTest("a", 0, true)
	findTest("d", 3, true)
	findTest("e", -1, false)
}

func TestFindIndexWith(t *testing.T) {
	comparison := func(el1 interface{}, el2 interface{})bool {
		if el1 == "e" && el2 == "c" {
			return true
		} else if el1 == el2 {
			return true
		} else {
			return false
		}
	}

	items := DashSlice{"a", "b", "c", "d"}
	findTest := func(el interface{}, expected int, found bool) {
		result, ok := FindIndexWith(items, el, comparison)

		assert.Equal(t, ok, found)
		assert.Equal(t, result, expected)
	}

	findTest("e", 2, true)
	findTest("a", 0, true)
	findTest("b", 1, true)
	findTest("f", -1, false)
}

func TestDifference1(t *testing.T) {
	items1 := DashSlice{"a", "b", "c", "d"}
	items2 := DashSlice{"a", "c", "e", "f"}

	result := Difference(items1, items2...)
	assert.DeepEqual(t, result, DashSlice{"b", "d"})
}

func TestDifference2(t *testing.T) {
	items1 := DashSlice{"z", "b", "q", "h"}
	items2 := DashSlice{"b", "h", "e", "f"}

	result := Difference(items1, items2...)
	assert.DeepEqual(t, result, DashSlice{"z", "q"})
}

func TestDifferenceBy(t *testing.T) {
	items1 := DashSlice{1.2, 2.4, 5.9}
	items2 := DashSlice{1.3, 3.4, 5.1}

	result := DifferenceBy(items1, items2, func(el interface{}) interface{} {
		return math.Floor(el.(float64))
	})

	assert.DeepEqual(t, result, DashSlice{2.4})
}

func TestDifferenceWith(t *testing.T) {
	items1 := DashSlice{1.2, 2.4, 5.9}
	items2 := DashSlice{1.3, 3.4, 5.1}

	result := DifferenceWith(items1, items2, func(el1 interface{}, el2 interface{}) bool {
		if el1 == 1.2 && el2 == 1.3 {
			return true
		} else if el1 == 2.4 && el2 == 3.4 {
			return true
		} else {
			return false
		}
	})

	assert.DeepEqual(t, result, DashSlice{5.9})
}
