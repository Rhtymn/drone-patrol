package util

func Range(val, lower, higher int) bool {
	if val < lower || val > higher {
		return false
	}
	return true
}
