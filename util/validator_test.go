package util_test

import (
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/util"
)

func TestRange(t *testing.T) {
	t.Run("should return false when given value given not in range of lower to higher (inclusive)", func(t *testing.T) {
		val, l, h := 10, 11, 20

		res := util.Range(val, l, h)

		if res {
			t.Fatalf("should return false")
		}
	})

	t.Run("should return true when given value in range of lower to higher (inclusive)", func(t *testing.T) {
		val, l, h := 1, 0, 1

		res := util.Range(val, l, h)

		if !res {
			t.Fatalf("should return true")
		}
	})
}
