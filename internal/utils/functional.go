package utils

func Map[T any, U any](slc []T, fn func(int, T) U) []U {
	newSlc := make([]U, len(slc))
	for i, v := range slc {
		newSlc[i] = fn(i, v)
	}
	return newSlc
}

func Reduce[T any, U any](slc []T, fn func(int, U, T) U, dflt U) U {
	accumulator := dflt
	for i, v := range slc {
		accumulator = fn(i, accumulator, v)
	}
	return accumulator
}

func Filter[T any](slc []T, fn func(int, T) bool) []T {
	newSlc := []T{}
	for i, v := range slc {
		if fn(i, v) {
			newSlc = append(newSlc, v)
		}
	}
	return newSlc
}
