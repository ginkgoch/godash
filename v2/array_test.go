package godash

import (
	"math"
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

func TestCompact(t *testing.T) {
	items := []interface{}{"a", "b", false, 0, 1, nil, ""}
	compacted := Compact(items)

	assert.DeepEqual(t, compacted, []interface{}{"a", "b", 1})
}

func TestConcat1(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	result := Concat(items, []string{"e", "f"})

	assert.DeepEqual(t, result, []string{"a", "b", "c", "d", "e", "f"})
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

func TestIndexOf(t *testing.T) {
	items := []string{"a", "b", "c", "d"}

	findTest := func(el string, expected int, found bool) {
		result, ok := IndexOf(items, el)

		assert.Equal(t, ok, found)
		assert.Equal(t, result, expected)
	}

	findTest("a", 0, true)
	findTest("d", 3, true)
	findTest("e", -1, false)
}

func TestDifferenceBy(t *testing.T) {
	items1 := []float64{1.2, 2.4, 5.9}
	items2 := []float64{1.3, 3.4, 5.1}

	result := DifferenceBy(items1, items2, func(el float64) float64 {
		return math.Floor(el)
	})

	assert.DeepEqual(t, result, []float64{2.4})
}
