package xarray

func RemoveFirst[T comparable](arr []T, elem T) ([]T, bool) {
	for i, v := range arr {
		if v == elem {
			return append(arr[:i], arr[i+1:]...), true
		}
	}
	return arr, false
}
