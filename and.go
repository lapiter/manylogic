package manylogic

// And returns true if all inputs are true
// otherwise returns false
// returns true when called with no input
func And(inputs ...bool) bool {
	for _, input := range inputs {
		if !input {
			return false
		}
	}
	return true
}

// AndCb returns true if the callback returns true with all inputs
// otherwise returns false
// returns true when called with no input
func AndCb[T any](inputs []T, callback func(T) bool) bool {
	for _, input := range inputs {
		if !callback(input) {
			return false
		}
	}
	return true
}
