package bns

func Ternary[T any](condition bool, truth T, falsy T) T {
	if condition {
		return truth
	} else {
		return falsy
	}
}
