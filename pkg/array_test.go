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

func findIndexTemp(t *testing.T, find func(DashSlice, Predicate) (int, bool)) {
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

func TestIntersection(t *testing.T) {
	items1 := DashSlice{1, 2, 3, 4}
	items2 := DashSlice{3, 4, 5, 6}

	result := Intersection(items1, items2)
	assert.DeepEqual(t, result, DashSlice{3, 4})
}

func TestIntersectionBy(t *testing.T) {
	items1 := DashSlice{1.0, 2.0, 3.2, 4.0}
	items2 := DashSlice{3.4, 4.3, 5.0, 6.0}

	result := IntersectionBy(items1, items2, func(i interface{}) interface{} {
		return math.Floor(i.(float64))
	})

	assert.DeepEqual(t, result, DashSlice{3.2, 4.0})
}

func TestIntersectionWith(t *testing.T) {
	items1 := DashSlice{1.0, 2.0, 3.2, 4.0}
	items2 := DashSlice{3.4, 4.6, 5.0, 6.0}

	result := IntersectionWith(items1, items2, func(i1 interface{}, i2 interface{}) bool {
		return math.Round(i1.(float64)) == math.Round(i2.(float64))
	})
	assert.DeepEqual(t, result, DashSlice{3.2})
}

func TestJoin(t *testing.T) {
	items := DashSlice{1, 2, 3}

	result := Join(items, "~")
	assert.Equal(t, result, "1~2~3")
}

func TestLast(t *testing.T) {
	items := DashSlice{1, 2, 3}
	result := Last(items)
	assert.Equal(t, result, 3)

	items = DashSlice{}
	result = Last(items)
	assert.Equal(t, result, nil)
}

func TestNth(t *testing.T) {
	items := DashSlice{1, 2, 3}

	testNth := func(i int, r interface{}) {
		result := Nth(items, i)
		assert.Equal(t, result, r)
	}

	testNth(0, 1)
	testNth(1, 2)
	testNth(2, 3)
	testNth(3, nil)
	testNth(-1, 3)
	testNth(-2, 2)
	testNth(-3, 1)
	testNth(-4, nil)
}

func TestPull(t *testing.T) {
	items := DashSlice{"a", "b", "c", "a", "b", "c"}
	Pull(&items, "a", "c")
	assert.DeepEqual(t, items, DashSlice{"b", "b"})

	items = DashSlice{"a", "b", "c", "a", "b", "c"}
	PullAll(&items, DashSlice{"a", "c"})
	assert.DeepEqual(t, items, DashSlice{"b", "b"})

	items = DashSlice{"a", "b", "c", "a", "b", "c"}
	PullAllWith(&items, DashSlice{"a", "c"}, func(i1 interface{}, i2 interface{}) bool {
		if i1 == "b" && i2 == "a" {
			return true
		} else {
			return i1 == i2
		}
	})

	assert.DeepEqual(t, items, DashSlice{})
}

func TestPullAt(t *testing.T) {
	items := DashSlice{"a", "b", "c", "a", "b", "c"}
	pulled := PullAt(&items, 1, 3, 5)

	assert.DeepEqual(t, pulled, DashSlice{"b", "a", "c"})
	assert.DeepEqual(t, items, DashSlice{"a", "c", "b"})
}

func TestRemove(t *testing.T) {
	items := DashSlice{1, 2, 3, 4, 5, 6, 7}
	removed := Remove(&items, func(i interface{}) bool {
		return i.(int)%2 == 0
	})

	assert.DeepEqual(t, items, DashSlice{1, 3, 5, 7})
	assert.DeepEqual(t, removed, DashSlice{2, 4, 6})
}

func TestSlice(t *testing.T) {
	items := DashSlice{1, 2, 3, 4, 5, 6, 7}
	result := Slice(items, 2, 5)

	assert.DeepEqual(t, result, DashSlice{3, 4, 5})
}

func TestTail(t *testing.T) {
	items := DashSlice{1, 2, 3, 4}
	result := Tail(items)
	assert.DeepEqual(t, result, DashSlice{2, 3, 4})

	items = DashSlice{1}
	result = Tail(items)
	assert.DeepEqual(t, result, DashSlice{})

	items = DashSlice{}
	result = Tail(items)
	assert.DeepEqual(t, result, DashSlice{})
}

func TestTake(t *testing.T) {
	items := DashSlice{1, 2, 3, 4}
	result := Take(items, 2)
	assert.DeepEqual(t, result, DashSlice{1, 2})

	items = DashSlice{1, 2, 3, 4}
	result = Take(items, 9)
	assert.DeepEqual(t, result, DashSlice{1, 2, 3, 4})
}

func TestTakeWhile(t *testing.T) {
	items := DashSlice{1, 3, 4, 6, 8}
	result := TakeWhile(items, func(i interface{}) bool {
		return i.(int)%2 != 0
	})

	assert.DeepEqual(t, result, DashSlice{1, 3})
}

func TestTakeRight(t *testing.T) {
	items := DashSlice{1, 2, 3, 4}
	result := TakeRight(items, 2)
	assert.DeepEqual(t, result, DashSlice{3, 4})

	items = DashSlice{1, 2, 3, 4}
	result = Take(items, 9)
	assert.DeepEqual(t, result, DashSlice{1, 2, 3, 4})
}

func TestTakeRightWhile(t *testing.T) {
	items := DashSlice{1, 3, 4, 6, 8}
	result := TakeRightWhile(items, func(i interface{}) bool {
		return i.(int)%2 == 0
	})

	assert.DeepEqual(t, result, DashSlice{4, 6, 8})
}

func TestUnion(t *testing.T) {
	result := Union(DashSlice{2}, DashSlice{1, 2})
	assert.DeepEqual(t, result, DashSlice{2, 1})
}

func TestUnionBy(t *testing.T) {
	result := UnionBy(func(i interface{}) interface{} {
		return math.Floor(i.(float64))
	}, DashSlice{2.1}, DashSlice{1.2, 2.3})

	assert.DeepEqual(t, result, DashSlice{2.1, 1.2})
}

func TestUnionWith(t *testing.T) {
	result := UnionWith(func(i1 interface{}, i2 interface{}) bool {
		return i1.(int)%3 == i2.(int)%3
	}, DashSlice{1, 2, 3}, DashSlice{5, 8}, DashSlice{2, 6})

	assert.DeepEqual(t, result, DashSlice{1, 2, 3})
}

func TestUniq(t *testing.T) {
	items := DashSlice{1, 3, 8, 6, 8, 2, 3, 2, 3, 2}
	result := Uniq(items)
	assert.DeepEqual(t, result, DashSlice{1, 3, 8, 6, 2})
}

func TestUniqBy(t *testing.T) {
	items := DashSlice{2.1, 1.2, 2.4}
	result := UniqBy(items, func(i interface{}) interface{} {
		return math.Floor(i.(float64))
	})

	assert.DeepEqual(t, result, DashSlice{2.1, 1.2})
}

func TestUniqWith(t *testing.T) {
	items := DashSlice{1, 2, 3, 5, 8, 2, 6}
	result := UniqWith(items, func(i1 interface{}, i2 interface{}) bool {
		return i1.(int)%3 == i2.(int)%3
	})

	assert.DeepEqual(t, result, DashSlice{1, 2, 3})
}
