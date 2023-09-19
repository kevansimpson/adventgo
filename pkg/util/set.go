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

func Intersection[T comparable](set1 Set[T], set2 Set[T]) Set[T] {
	set3 := make(Set[T])
	for pt := range set1 {
		_, hasPt := set2[pt]
		if hasPt {
			Add(set3, pt)
		}
	}

	return set3
}

func SetToSlice[T comparable](set Set[T]) []T {
	var list []T
	for key := range set {
		list = append(list, key)
	}
	return list
}

func SliceToSet[T comparable](slice []T) Set[T] {
	set := make(Set[T])
	AddAll(set, slice)
	return set
}
