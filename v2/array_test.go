package godash

import (
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
