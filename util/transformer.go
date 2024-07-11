package util

import (
	"strconv"
	"strings"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"
)

// IntSlice transform separated white space numeric string to slice of int
// If given string with separator other than white space, it will return an internal error
// If given non numeric string, it will also return an internal error
func IntSlice(val string) ([]int, error) {
	// create array of string from given string (separated by white space)
	arrString := strings.Fields(val)
	intSlice := []int{}

	for i := 0; i < len(arrString); i++ {
		// convert string to int
		i, err := strconv.Atoi(arrString[i])
		if err != nil {
			return nil, apperror.Wrap(err)
		}
		intSlice = append(intSlice, i)
	}
	return intSlice, nil
}
