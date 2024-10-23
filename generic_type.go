package main

func getLast[T any](s []T) T {

	if len(s) == 0 {
		var zero T
		return zero
	}

	last := len(s) - 1
	return s[last]

}
