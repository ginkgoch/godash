package godash

import (
	"gotest.tools/assert"
	"math"
	"testing"
)

func TestDashSlice_Map(t *testing.T) {
	items := DashSlice{1.2, 2.3, 3.4, 5.6}

	result := items.Map(func(el interface{}) interface{} {
		return math.Floor(el.(float64))
	})

	assert.DeepEqual(t, result, DashSlice{1.0, 2.0, 3.0, 5.0})
}
