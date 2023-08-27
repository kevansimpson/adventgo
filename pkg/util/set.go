package util

type void struct{}

var member void

type Set[T comparable] map[T]void

func Add[T comparable](set Set[T], elem T) {
	set[elem] = member
}

func AddAll[T comparable](set Set[T], elem []T) {
	for _, e := range elem {
		set[e] = member
	}
}

func AddSet[T comparable](set Set[T], elem Set[T]) {
	for e := range elem {
		Add(set, e)
	}
}

func SetToSlice[T comparable](set Set[T]) []T {
	var list []T
	for key := range set {
		list = append(list, key)
	}
	return list
}
