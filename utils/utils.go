package utils

func StringInSlice(s string, list []string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	return false
}

func IntInSlice(i int, list []int) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func IndexOfSlice(s string, list []string) int {
	for i, b := range list {
		if b == s {
			return i
		}
	}
	return -1
}

