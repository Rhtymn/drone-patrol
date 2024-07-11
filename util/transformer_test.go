package util_test

import (
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/util"
)

func TestIntSlice(t *testing.T) {
	t.Run("should return error internal when not given numerical string", func(t *testing.T) {
		str := "a b c"

		_, err := util.IntSlice(str)

		if err == nil || !apperror.ErrorIs(err, apperror.Internal) {
			t.Fatalf("should return error internal")
		}
	})

	t.Run("should not return any error when given numerical string", func(t *testing.T) {
		str := "1 2 3"

		_, err := util.IntSlice(str)

		if err != nil {
			t.Fatalf("should not return any error")
		}
	})

	t.Run("should return error internal when given numerical string but invalid format", func(t *testing.T) {
		str := "1 2,3"

		_, err := util.IntSlice(str)

		if err == nil || !apperror.ErrorIs(err, apperror.Internal) {
			t.Fatalf("should return error internal")
		}
	})
}
