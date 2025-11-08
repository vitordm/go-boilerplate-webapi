package utils

func Map[T, R any](slice []T, f func(T, int) R) []R {
	var result []R
	for i, item := range slice {
		result = append(result, f(item, i))
	}
	return result
}

func FirstOrDefault[T any](slice []T, df T) T {
	if len(slice) == 0 {
		return df
	}

	return slice[0]
}
