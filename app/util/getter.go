package util

func ValueOrDefault[T any](test bool, then T, otherwise T) T {
	if test {
		return then
	}
	return otherwise
}
