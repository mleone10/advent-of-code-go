package maputil

// Keys returns a slice of all keys in a given map.
func Keys[K comparable, V any](m map[K]V) []K {
	var keys []K

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// Values returns a slice of all values in a given map.
func Values[K comparable, V any](m map[K]V) []V {
	var vals []V

	for _, v := range m {
		vals = append(vals, v)
	}

	return vals
}

// TODO: implement methods to simplify dealing with maps of key to slice
