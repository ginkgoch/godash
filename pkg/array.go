package pkg

import "reflect"

func Chunk(items DashSlice, size int) []DashSlice {
	dashSlices := []DashSlice{}

	for _, item := range items {
		sliceLength := len(dashSlices)
		if sliceLength == 0 || len(dashSlices[sliceLength-1]) == size {
			dashSlices = append(dashSlices, DashSlice{})
			sliceLength++
		}

		dashSlices[sliceLength-1] = append(dashSlices[sliceLength-1], item)
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

// Creates a new array concatenating array with any additional arrays and/or values.
func Concat(items DashSlice, newItems ...interface{}) DashSlice {
	result := append(DashSlice{}, items...)

	for _, newItem := range newItems {
		if reflect.TypeOf(newItem).Kind() == reflect.Slice {
			values := reflect.ValueOf(newItem)
			valuesLength := values.Len()
			for i := 0; i < valuesLength; i++ {
				result = append(result, values.Index(i).Interface())
			}
		} else {
			result = append(result, newItem)
		}
	}

	return result
}
