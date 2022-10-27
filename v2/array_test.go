package godash

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"gotest.tools/assert"
)

func TestChunk1(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	chunked := Chunk(items, 2)

	assert.Equal(t, len(chunked), 2)
	assert.DeepEqual(t, chunked[0], []string{"a", "b"})
	assert.DeepEqual(t, chunked[1], []string{"c", "d"})
}

func TestChunk2(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	chunked := Chunk(items, 3)

	assert.Equal(t, len(chunked), 2)
	assert.DeepEqual(t, chunked[0], []string{"a", "b", "c"})
	assert.DeepEqual(t, chunked[1], []string{"d"})
}

func TestChunk3(t *testing.T) {
	items := []int{1, 2, 3, 4}
	chunked := Chunk(items, 2)

	assert.Equal(t, len(chunked), 2)
	assert.DeepEqual(t, chunked[0], []int{1, 2})
	assert.DeepEqual(t, chunked[1], []int{3, 4})
}

func ExampleChunk() {
	items := []string{"a", "b", "c", "d"}
	chunked := Chunk(items, 3)
	fmt.Println(chunked)
	// Output:
	// [[a b c] [d]]
}

func TestCompact(t *testing.T) {
	items := []interface{}{"a", "b", false, 0, 1, nil, ""}
	compacted := Compact(items)

	assert.DeepEqual(t, compacted, []interface{}{"a", "b", 1})
}

func ExampleCompact() {
	items := []interface{}{"a", "b", false, 0, 1, nil, ""}
	compacted := Compact(items)
	fmt.Println(compacted...)
	// Output:
	// a b 1
}

func TestConcat1(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	result := Concat(items, []string{"e", "f"})

	assert.DeepEqual(t, result, []string{"a", "b", "c", "d", "e", "f"})
}

func ExampleConcat() {
	items := []string{"a", "b", "c", "d"}
	result := Concat(items, []string{"e", "f"})
	fmt.Println(result)
	// Output:
	// [a b c d e f]
}

func TestDifference1(t *testing.T) {
	items1 := []string{"a", "b", "c", "d"}
	items2 := []string{"a", "c", "e", "f"}

	result := Difference(items1, items2)
	assert.DeepEqual(t, result, []string{"b", "d"})
}

func TestDifference2(t *testing.T) {
	items1 := []string{"z", "b", "q", "h"}
	items2 := []string{"b", "h", "e", "f"}

	result := Difference(items1, items2)
	assert.DeepEqual(t, result, []string{"z", "q"})
}

func ExampleDifference() {
	items1 := []string{"a", "b", "c", "d"}
	items2 := []string{"a", "c", "e", "f"}

	result := Difference(items1, items2)
	fmt.Println(result)
	// Output:
	// [b d]
}

func indexOfTemp(t *testing.T, find func([]string, string) (int, bool)) {
	items := []string{"a", "b", "c", "d"}

	findTest := func(el string, expected int, found bool) {
		result, ok := find(items, el)

		assert.Equal(t, ok, found)
		assert.Equal(t, result, expected)
	}

	findTest("a", 0, true)
	findTest("d", 3, true)
	findTest("e", -1, false)
}

func TestIndexOf(t *testing.T) {
	indexOfTemp(t, IndexOf[string])
}

func ExampleIndexOf() {
	items := []string{"a", "b", "c", "d"}
	result1, ok := IndexOf(items, "1")
	fmt.Println(ok, result1)

	result2, ok := IndexOf(items, "c")
	fmt.Println(ok, result2)
	// Output:
	// false -1
	// true 2
}

func TestLastIndexOf(t *testing.T) {
	indexOfTemp(t, LastIndexOf[string])
}

func ExampleLastIndexOf() {
	items := []string{"a", "b", "c", "b", "d"}
	result1, ok := LastIndexOf(items, "1")
	fmt.Println(ok, result1)

	result2, ok := LastIndexOf(items, "b")
	fmt.Println(ok, result2)
	// Output:
	// false -1
	// true 3
}

func TestDifferenceBy(t *testing.T) {
	items1 := []float64{1.2, 2.4, 5.9}
	items2 := []float64{1.3, 3.4, 5.1}

	result := DifferenceBy(items1, items2, func(el float64) float64 {
		return math.Floor(el)
	})

	assert.DeepEqual(t, result, []float64{2.4})
}

func ExampleDifferenceBy() {
	items1 := []float64{1.2, 2.4, 5.9}
	items2 := []float64{1.3, 3.4, 5.1}

	result := DifferenceBy(items1, items2, func(el float64) float64 {
		return math.Floor(el)
	})
	fmt.Println(result)
	// Output:
	// [2.4]
}

func TestDifferenceWith(t *testing.T) {
	items1 := []float64{1.2, 2.4, 5.9}
	items2 := []float64{1.3, 3.4, 5.1}

	result := DifferenceWith(items1, items2, func(el1 float64, el2 float64) bool {
		if el1 == 1.2 && el2 == 1.3 {
			return true
		} else if el1 == 2.4 && el2 == 3.4 {
			return true
		} else {
			return false
		}
	})

	assert.DeepEqual(t, result, []float64{5.9})
}

func ExampleDifferenceWith() {
	items1 := []float64{1.2, 2.4, 5.9}
	items2 := []float64{1.3, 3.4, 5.1}

	result := DifferenceWith(items1, items2, func(el1 float64, el2 float64) bool {
		if el1 == 1.2 && el2 == 1.3 {
			return true
		} else if el1 == 2.4 && el2 == 3.4 {
			return true
		} else {
			return false
		}
	})
	fmt.Println(result)
	// Output:
	// [5.9]
}

func TestDrop(t *testing.T) {
	items := []string{"a", "b", "c", "d"}

	result := Drop(items, 0)
	assert.DeepEqual(t, result, []string{"a", "b", "c", "d"})

	result = Drop(items, 1)
	assert.DeepEqual(t, result, []string{"b", "c", "d"})

	result = Drop(items, 2)
	assert.DeepEqual(t, result, []string{"c", "d"})

	result = Drop(items, 4)
	assert.DeepEqual(t, result, []string{})
}

func ExampleDrop() {
	items := []string{"a", "b", "c", "d"}
	result := Drop(items, 1)
	fmt.Println(result)
	// Output:
	// [b c d]
}

func TestDropRight(t *testing.T) {
	items := []string{"a", "b", "c", "d"}

	result := DropRight(items, 0)
	assert.DeepEqual(t, result, []string{"a", "b", "c", "d"})

	result = DropRight(items, 1)
	assert.DeepEqual(t, result, []string{"a", "b", "c"})

	result = DropRight(items, 4)
	assert.DeepEqual(t, result, []string{})

	result = DropRight(items, 5)
	assert.DeepEqual(t, result, []string{})
}

func ExampleDropRight() {
	items := []string{"a", "b", "c", "d"}
	result := DropRight(items, 1)
	fmt.Println(result)
	// Output:
	// [a b c]
}

func TestDropWhile(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	result := DropWhile(items, func(el string) bool {
		return el == "a" || el == "b"
	})

	assert.DeepEqual(t, result, []string{"c", "d"})
}

func ExampleDropWhile() {
	items := []string{"a", "b", "c", "d"}
	result := DropWhile(items, func(el string) bool {
		return el == "a" || el == "b"
	})
	fmt.Println(result)
	// Output:
	// [c d]
}

func findIndexWithTemp(t *testing.T, find func([]string, string, Comparison[string]) (int, bool)) {
	comparison := func(el1 string, el2 string) bool {
		if el1 == "e" && el2 == "c" {
			return true
		} else if el1 == el2 {
			return true
		} else {
			return false
		}
	}

	items := []string{"a", "b", "c", "d"}
	findTest := func(el string, expected int, found bool) {
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
	findIndexWithTemp(t, FindIndexWith[string])
}

func ExampleFindIndexWith() {
	items := []float64{2.3, 3.4, 4.5, 5.6}
	comparison := func(a float64, b float64) bool {
		return math.Floor(a) == math.Floor(b)
	}
	result, ok := FindIndexWith(items, 4.1, comparison)
	fmt.Println(result, ok)
	// Output:
	// 2 true
}

func TestFillBy(t *testing.T) {
	items := []string{"a", "b", "c", "d"}

	FillInRange(items, "z", 0, 4)
	assert.DeepEqual(t, items, []string{"z", "z", "z", "z"})

	FillInRange(items, "y", 2, 4)
	assert.DeepEqual(t, items, []string{"z", "z", "y", "y"})
}

func ExampleFillInRange() {
	items := make([]string, 4)
	FillInRange(items, "z", 0, len(items))
	fmt.Println(items)
	// Output:
	// [z z z z]
}

func TestFill(t *testing.T) {
	items := []string{"a", "b", "c", "d"}

	Fill(items, "z")
	assert.DeepEqual(t, items, []string{"z", "z", "z", "z"})

	Fill(items, "y")
	assert.DeepEqual(t, items, []string{"y", "y", "y", "y"})
}

func ExampleFill() {
	items := []string{"a", "b", "c", "d"}
	Fill(items, "z")
	fmt.Println(items)
	// Output:
	// [z z z z]
}

func findIndexTemp(t *testing.T, find func([]string, Predicate[string]) (int, bool)) {
	items := []string{"a", "b", "c", "d"}
	i, ok := find(items, func(el string) bool {
		return el == "c"
	})

	assert.Equal(t, i, 2)
	assert.Equal(t, ok, true)

	j, ok := find(items, func(el string) bool {
		return el == "z"
	})

	assert.Equal(t, j, -1)
	assert.Equal(t, ok, false)
}

func TestFindIndex(t *testing.T) {
	findIndexTemp(t, FindIndex[string])
}

func ExampleFindIndex() {
	items := []string{"a", "b", "c", "d"}
	result, ok := FindIndex(items, func(el string) bool {
		return el == "c"
	})
	fmt.Println(result, ok)
	// Output:
	// 2 true
}

func TestFindLastIndexWith(t *testing.T) {
	findIndexWithTemp(t, FindLastIndexWith[string])
}

func ExampleFindLastIndexWith() {
	items := []string{"a", "b", "b", "d"}
	result, ok := FindIndex(items, func(el string) bool {
		return el == "b"
	})
	fmt.Println(result, ok)
	// Output:
	// 1 true
}

func ExampleFindLastIndex() {
	items := []string{"a", "b", "b", "d"}
	result, ok := FindLastIndex(items, func(el string) bool {
		return el == "b"
	})
	fmt.Println(result, ok)
	// Output:
	// 1 true
}

func TestFirst1(t *testing.T) {
	items := []int{1, 2, 4}
	i := First(items)

	assert.Equal(t, *i, 1)
}

func TestFirst2(t *testing.T) {
	items := []int{}
	i := First(items)

	assert.Equal(t, i == nil, true)
}

func TestFromPairsAny(t *testing.T) {
	pairs := [][]any{
		{1, 2},
		{"key1", "value1"},
		{"key2"},
		{},
	}

	result := FromPairsAny(pairs)
	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[1], 2)
	assert.Equal(t, result["key1"], "value1")
	assert.Equal(t, result["key2"], nil)
}

func TestFromPairs(t *testing.T) {
	a := KeyValuePair[string, any]{"key0", 2}
	b := KeyValuePair[string, any]{"key1", "value1"}
	c := KeyValuePair[string, any]{"key2", nil}
	pairs := []KeyValuePair[string, any]{
		a, b, c,
	}

	result := FromPairs(pairs)
	assert.Equal(t, len(result), 3)
	assert.Equal(t, result["key0"], 2)
	assert.Equal(t, result["key1"], "value1")
	assert.Equal(t, result["key2"], nil)
}

func TestInitial(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	result := Initial(items)

	assert.Equal(t, 4, len(result))
	assert.DeepEqual(t, []int{1, 2, 3, 4}, result)
}

func TestMap(t *testing.T) {
	items := []string{"1", "2", "3"}
	results := Map(items, func(s string) int {
		if r, err := strconv.Atoi(s); err != nil {
			panic(err)
		} else {
			return r
		}
	})

	assert.DeepEqual(t, results, []int{1, 2, 3})
}

func TestFilter(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6}
	results := Filter(items, func(i int) bool {
		return i%2 == 0
	})

	assert.DeepEqual(t, results, []int{2, 4, 6})
}

func TestIntersection(t *testing.T) {
	items1 := []int{1, 2, 3, 4}
	items2 := []int{3, 2, 5, 9}
	result := Intersection(items1, items2)

	assert.DeepEqual(t, result, []int{2, 3})
}

type Person struct {
	Name string
}

func TestIntersectionBy(t *testing.T) {
	items1 := []Person{{Name: "A"}, {Name: "B"}}
	items2 := []Person{{Name: "B"}, {Name: "C"}}

	results := IntersectionBy(items1, items2, func(item Person) string {
		return item.Name
	})

	assert.Equal(t, len(results), 1)
	assert.Equal(t, results[0].Name, "B")
}

func TestIntersectionWith(t *testing.T) {
	items1 := []int{1, 2, 3}
	items2 := []int{3, 4, 5}

	results := IntersectionWith(items1, items2, func(i1 int, i2 int) bool {
		return i1 == i2
	})

	assert.Equal(t, len(results), 1)
	assert.Equal(t, results[0], 3)
}

func TestJoin(t *testing.T) {
	items1 := []string{"1", "2", "3"}

	result := Join(items1, ", ")

	assert.Equal(t, result, "1, 2, 3")

	items2 := []int{4, 5, 6}

	result = Join(items2, ", ")

	assert.Equal(t, result, "4, 5, 6")
}

func TestLast(t *testing.T) {
	items := []string{"1", "2", "3"}
	exists, lastItem := Last(items)
	assert.Equal(t, true, exists)
	assert.Equal(t, "3", lastItem)

	items = []string{}
	exists, _ = Last(items)
	assert.Equal(t, false, exists)
}

func TestNth(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6}

	exists, i1 := Nth(items, 2)
	assert.Equal(t, exists, true)
	assert.Equal(t, i1, 3)

	exists, i2 := Nth(items, -1)
	assert.Equal(t, exists, true)
	assert.Equal(t, i2, 6)

	exists, i3 := Nth(items, -6)
	assert.Equal(t, exists, true)
	assert.Equal(t, i3, 1)

	exists, _ = Nth(items, 6)
	assert.Equal(t, exists, false)

	exists, _ = Nth(items, -7)
	assert.Equal(t, exists, false)
}

func TestPull(t *testing.T) {
	items := []int{1, 2, 3, 4}
	results := Pull(&items, 1, 2)

	assert.Equal(t, len(results), 2)
	assert.DeepEqual(t, results, []int{3, 4})
	assert.DeepEqual(t, items, []int{3, 4})
}

func TestPullAll(t *testing.T) {
	items := []int{1, 2, 3, 4}
	itemsToPull := []int{3, 4}
	results := PullAll(&items, itemsToPull)

	assert.Equal(t, len(results), 2)
	assert.DeepEqual(t, results, []int{1, 2})
	assert.DeepEqual(t, items, []int{1, 2})
}

func TestPullWith(t *testing.T) {
	items := &[]int{1, 2, 3, 4}
	itemsToPull := []int{3, 4}
	results := PullAllWith(items, itemsToPull, func(i1 int, i2 int) bool {
		return i1 == i2
	})

	assert.Equal(t, len(results), 2)
	assert.DeepEqual(t, results, []int{1, 2})
	assert.DeepEqual(t, *items, []int{1, 2})
}

func TestPullAt(t *testing.T) {
	items := []int{1, 2, 3, 4}
	results := PullAt(&items, 0, 2)

	assert.DeepEqual(t, items, []int{2, 4})
	assert.DeepEqual(t, results, []int{1, 3})
}
