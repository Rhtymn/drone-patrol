package util

// Range check given val is in range from lower to higher (inclusive)
// It will return false if given val doesn't not in range from lower to higher
// It will return true if given val which in range from to lower to higher
func Range(val, lower, higher int) bool {
	if val < lower || val > higher {
		return false
	}
	return true
}
