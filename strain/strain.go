package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	if i == nil {
		return nil
	}

	result := make(Ints, 0)

	for _, value := range i {
		if filter(value) {
			result = append(result, value)
		}
	}

	return result
}

func (i Ints) Discard(filter func(int) bool) Ints {
	if i == nil {
		return nil
	}

	result := make(Ints, 0)

	for _, value := range i {
		if !filter(value) {
			result = append(result, value)
		}
	}

	return result
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	if l == nil {
		return nil
	}

	result := make(Lists, 0)

	for _, value := range l {
		if filter(value) {
			result = append(result, value)
		}
	}

	return result
}

func (s Strings) Keep(filter func(string) bool) Strings {
	if s == nil {
		return nil
	}

	result := make(Strings, 0)

	for _, value := range s {
		if filter(value) {
			result = append(result, value)
		}
	}

	return result
}
