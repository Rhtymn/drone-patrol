package util

import (
	"strconv"
	"strings"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"
)

func IntSlice(val string) ([]int, error) {
	arrString := strings.Fields(val)
	intSlice := []int{}

	for i := 0; i < len(arrString); i++ {
		i, err := strconv.Atoi(arrString[i])
		if err != nil {
			return nil, apperror.Wrap(err)
		}
		intSlice = append(intSlice, i)
	}
	return intSlice, nil
}
