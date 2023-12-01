package manylogic

// Xor returns true if one and only one input is true
// otherwise returns false
// returns false when called with no input
func Xor(inputs ...bool) bool {
	previousTrue := false
	for _, input := range inputs {
		if !input {
			continue
		}
		if previousTrue {
			return false
		}
		previousTrue = true
	}
	return previousTrue
}

// XorCb returns true if the callback returns true with one and only one input
// otherwise returns false
// returns false when called with no input
func XorCb[T any](inputs []T, callback func(T) bool) bool {
	previousTrue := false
	for _, input := range inputs {
		if !callback(input) {
			continue
		}
		if previousTrue {
			return false
		}
		previousTrue = true
	}
	return previousTrue
}
