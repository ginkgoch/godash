package pkg

import (
	"gotest.tools/assert"
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
