package godash

func Ternary(satisfy bool, truthyValue interface{}, falsyValue interface{}) interface{} {
	if satisfy {
		return truthyValue
	} else {
		return falsyValue
	}
}
