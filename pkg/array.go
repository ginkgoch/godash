package pkg

func Chunk(items DashSlice, size int) []DashSlice {
	dashSlices := []DashSlice{}

	for _, item := range items {
		sliceLength := len(dashSlices)
		if sliceLength == 0 || len(dashSlices[sliceLength - 1]) == size {
			dashSlices = append(dashSlices, DashSlice{})
			sliceLength++
		}

		dashSlices[sliceLength - 1] = append(dashSlices[sliceLength - 1], item)
	}

	return dashSlices
}

// Creates an array with all falsey values removed. The values false, 0, "", nil are falsey.
func Compact(items DashSlice) DashSlice {
	dashSlice := DashSlice{}

	for _, item := range items {
		if item != nil && item != false && item != 0 && item != "" {
			dashSlice = append(dashSlice, item)
		}
	}

	return dashSlice
}