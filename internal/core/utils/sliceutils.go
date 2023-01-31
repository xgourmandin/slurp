package utils

func Contains[T comparable](slice []T, searched T) bool {
	for _, v := range slice {
		if v == searched {
			return true
		}
	}
	return false
}
