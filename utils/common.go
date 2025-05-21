package utils

func ContainsInSlice[T any](slice []T, elem T) bool {
	var comparator func(any, any) bool
	switch any(elem).(type) {
	default:
		comparator = comparatorString
	}
	for _, v := range slice {
		if comparator(v, elem) {
			return true
		}
	}
	return false
}

func comparatorString(s1, s2 any) bool {
	return s1 == s2
}
