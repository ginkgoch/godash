package godash

type DashSlice []interface{}

func (ds DashSlice) AsSliceString() []string {
	 sliceString := []string{}
	 for _, str := range ds {
	 	sliceString = append(sliceString, str.(string))
	 }

	 return sliceString
}


