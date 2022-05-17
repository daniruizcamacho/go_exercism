package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Not same length")
	}

	count := 0
	for k := range a {
		if a[k] != b[k] {
			count += 1
		}
	}

	return count, nil
}
