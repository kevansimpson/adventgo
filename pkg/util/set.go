package util

type void struct{}

var member void

type Set[T comparable] map[T]void

func Add[T comparable](set Set[T], elem T) {
	set[elem] = member
}

func SetToSlice[T comparable](set Set[T]) []T {
	var list []T
	for key, _ := range set {
		list = append(list, key)
	}
	return list
}
