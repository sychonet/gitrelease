package util

import "errors"

// SliceIndex returns index of a given element in slice
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// GetOptionValue fetches value for a given option in command line arguments
func GetOptionValue(args []string, emsg string, ss string, sl string) (string, error) {
	var val string
	p := SliceIndex(len(args), func(i int) bool { return args[i] == ss })
	if p < 0 {
		p = SliceIndex(len(args), func(i int) bool { return args[i] == sl })
	}
	if p > 0 {
		if (p + 1) < len(args) {
			val = args[p+1]
			return val, nil
		}
	}
	err := errors.New(emsg)
	return val, err
}
