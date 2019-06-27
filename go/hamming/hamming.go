package hamming

import (
	"errors"
)

//Distance - given two DNA strands of type string, hamming difference is returned. Strings must be of equal length.
func Distance(a, b string) (int, error) {

	if len([]rune(a)) != len([]rune(b)) {
		return 0, errors.New("DNA strands must be of equal length")
	}

	hammingDistance := 0

	for k, v := range a {
		if rune(b[k]) != v {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
