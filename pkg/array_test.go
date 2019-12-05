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

func TestDrop(t *testing.T) {
	items := DashSlice{"a", "b", "c", "d"}

	result := Drop(items, 0)
	assert.DeepEqual(t, result, DashSlice{"a", "b", "c", "d"})

	result = Drop(items, 1)
	assert.DeepEqual(t, result, DashSlice{"b", "c", "d"})

	result = Drop(items, 2)
	assert.DeepEqual(t, result, DashSlice{"c", "d"})

	result = Drop(items, 4)
	assert.DeepEqual(t, result, DashSlice{})
}

func TestDropRight(t *testing.T) {
	items := DashSlice{"a", "b", "c", "d"}

	result := DropRight(items, 0)
	assert.DeepEqual(t, result, DashSlice{"a", "b", "c", "d"})

	result = DropRight(items, 1)
	assert.DeepEqual(t, result, DashSlice{"a", "b", "c"})

	result = DropRight(items, 4)
	assert.DeepEqual(t, result, DashSlice{})

	result = DropRight(items, 5)
	assert.DeepEqual(t, result, DashSlice{})
}

func TestDropWhile(t *testing.T) {
	items := DashSlice{"a", "b", "c", "d"}
	result := DropWhile(items, func(el interface{}) bool {
		return el == "a" || el == "b"
	})

	assert.DeepEqual(t, result, DashSlice{"c", "d"})
}

func TestFillBy(t *testing.T) {
	items := DashSlice{"a", "b", "c", "d"}

	FillInRange(items, "z", 0, 4)
	assert.DeepEqual(t, items, DashSlice{"z", "z", "z", "z"})

	FillInRange(items, "y", 2, 4)
	assert.DeepEqual(t, items, DashSlice{"z", "z", "y", "y"})
}

func TestFill(t *testing.T) {
	items := DashSlice{"a", "b", "c", "d"}

	Fill(items, "z")
	assert.DeepEqual(t, items, DashSlice{"z", "z", "z", "z"})

	Fill(items, "y")
	assert.DeepEqual(t, items, DashSlice{"y", "y", "y", "y"})
}

func indexOfTemp(t *testing.T, find func(DashSlice, interface{}) (int, bool)) {
	items := DashSlice{"a", "b", "c", "d"}

	findTest := func(el interface{}, expected int, found bool) {
		result, ok := find(items, el)

		assert.Equal(t, ok, found)
		assert.Equal(t, result, expected)
	}

	findTest("a", 0, true)
	findTest("d", 3, true)
	findTest("e", -1, false)
}

func TestIndexOf(t *testing.T) {
	indexOfTemp(t, IndexOf)
}

func TestLastIndexOf(t *testing.T) {
	indexOfTemp(t, LastIndexOf)
}

func findIndexWithTemp(t *testing.T, find func(DashSlice, interface{}, Comparison) (int, bool)) {
	comparison := func(el1 interface{}, el2 interface{}) bool {
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
		result, ok := find(items, el, comparison)

		assert.Equal(t, ok, found)
		assert.Equal(t, result, expected)
	}

	findTest("e", 2, true)
	findTest("a", 0, true)
	findTest("b", 1, true)
	findTest("f", -1, false)
}

func TestFindIndexWith(t *testing.T) {
	findIndexWithTemp(t, FindIndexWith)
}

func TestFindLastIndexWith(t *testing.T) {
	findIndexWithTemp(t, FindLastIndexWith)
}

func findIndexTemp(t *testing.T, find func(DashSlice, Prediction) (int, bool)) {
	items := DashSlice{"a", "b", "c", "d"}
	i, ok := find(items, func(el interface{}) bool {
		return el == "c"
	})

	assert.Equal(t, i, 2)
	assert.Equal(t, ok, true)

	j, ok := find(items, func(el interface{}) bool {
		return el == "z"
	})

	assert.Equal(t, j, -1)
	assert.Equal(t, ok, false)
}

func TestFindIndex(t *testing.T) {
	findIndexTemp(t, FindIndex)
}

func TestFindLastIndex(t *testing.T) {
	findIndexTemp(t, FindLastIndex)
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

func TestReverse(t *testing.T) {
	items := DashSlice{"a", "b", "c", "d", "e"}
	Reverse(items)

	assert.DeepEqual(t, items, DashSlice{"e", "d", "c", "b", "a"})
}

func TestFirst1(t *testing.T) {
	items := DashSlice{1, 2, 4}
	i := First(items)

	assert.Equal(t, i, 1)
}

func TestFirst2(t *testing.T) {
	items := DashSlice{}
	i := First(items)

	assert.Equal(t, i, nil)
}

func TestFlatten(t *testing.T) {
	items := DashSlice{1, 2, []int{4, 5}, 6, 7}
	result := Flatten(items)

	assert.DeepEqual(t, result, DashSlice{1, 2, 4, 5, 6, 7})
}

func TestFlattenDeep(t *testing.T) {
	items := DashSlice{1, 2, []interface{}{
		3, 4,
		[]int{5, 6},
		[]interface{}{
			7,
			[]interface{}{
				8,
				[]int{9, 10},
			},
		},
	}}

	result := FlattenDeep(items)
	assert.DeepEqual(t, result, DashSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func TestFlattenDepth(t *testing.T) {
	items := DashSlice{1, 2, []interface{}{
		3, 4,
		[]int{5, 6},
		[]interface{}{
			7,
			[]interface{}{
				8,
				[]int{9, 10},
			},
		},
	}}

	result := FlattenDepth(items, 3)
	assert.DeepEqual(t, result, DashSlice{1, 2, 3, 4, 5, 6, 7, 8,
		[]int{9, 10},
	})
}

func TestFromPairs(t *testing.T) {
	pairs := []DashSlice{
		{1, 2},
		{"key1", "value1"},
		{"key2"},
		{},
	}

	result := FromPairs(pairs)
	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[1], 2)
	assert.Equal(t, result["key1"], "value1")
	assert.Equal(t, result["key2"], nil)
}

func TestInitial(t *testing.T) {
	items := DashSlice{}
	result := Initial(items)
	assert.Equal(t, len(result), 0)

	items = DashSlice{1}
	result = Initial(items)
	assert.Equal(t, len(result), 0)

	items = DashSlice{1, 2, 3}
	result = Initial(items)
	assert.Equal(t, len(result), 2)
	assert.DeepEqual(t, result, DashSlice{1, 2})
}
