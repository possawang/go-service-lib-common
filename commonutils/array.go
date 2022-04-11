package commonutils

func ArrayFilter[V any](array []V, qualifier func(v V) bool) []V {
	var result []V
	for _, i := range array {
		if qualifier(i) {
			result = append(result, i)
		}
	}
	return result
}

func ArrayMap[V any, M any](array []V, changer func(v V) M) []M {
	var result []M
	for _, i := range array {
		result = append(result, changer(i))
	}
	return result
}

func ArrayExists[V any](array []V, qualifier func(v V) bool) bool {
	for _, i := range array {
		if qualifier(i) {
			return true
		}
	}
	return false
}
