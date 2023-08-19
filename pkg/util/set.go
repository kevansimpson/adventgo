package util

type void struct{}

var member void

type Set[T comparable] map[T]void

func Add[T comparable](set Set[T], elem T) {
	set[elem] = member
}
