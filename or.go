package manylogic

// Or returns true if at least one input is true
// otherwise returns false
// returns false when called with no input
func Or(inputs ...bool) bool {
	for _, input := range inputs {
		if input {
			return true
		}
	}
	return false
}

// AndCb returns true if the callback returns true with at least one input
// otherwise returns false
// returns false when called with no input
func OrCb[T any](inputs []T, callback func(T) bool) bool {
	for _, input := range inputs {
		if callback(input) {
			return true
		}
	}
	return false
}
