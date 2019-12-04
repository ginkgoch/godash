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
