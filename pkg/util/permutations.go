package util

// Returns all combinations with given length
func Combinations[T any](input []T, length int) [][]T {
	if length == 0 {
		return [][]T{make([]T, 0)}
	}

	var result [][]T
	max := len(input) - length
	for i := 0; i <= max; i++ {
		sub := Combinations(input[i+1:], length-1)
		for _, perm := range sub {
			sz := len(perm)
			add := make([]T, sz+1)
			add[0] = input[i]
			for a := 0; a < sz; a++ {
				add[a+1] = perm[a]
			}
			result = append(result, add)
		}
	}

	return result
}

// Returns all permutations with same length as input
func Permutations[T any](input []T) [][]T {
	perms := [][]T{}
	fact := factorial(len(input))
	for ix := int64(0); ix < fact; ix += 1 {
		perms = append(perms, permStart(ix, input))
	}

	return perms
}

func factorial(num int) int64 {
	f, max := int64(1), int64(num)
	for i := int64(2); i <= max; i += 1 {
		f *= i
	}
	return f
}

func permStart[T any](count int64, input []T) []T {
	copy := append([]T(nil), input...)
	return perm(count, copy, []T{})
}

func perm[T any](count int64, input []T, output []T) []T {
	sz := len(input)
	if sz == 0 {
		return output
	}

	fact := factorial(sz - 1)
	ix := count / fact
	value := input[ix]
	nextIn := append(input[:ix], input[ix+1:]...)
	nextOut := append(output, value)

	return perm(count%fact, nextIn, nextOut)
}
