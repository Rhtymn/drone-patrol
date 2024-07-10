package apperror_test

import (
	"errors"
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"
)

func TestNewAppError(t *testing.T) {
	t.Run("should return app error with correct code", func(t *testing.T) {
		msg, code := "message", apperror.InvalidArguments

		err := apperror.NewAppError(msg, code)

		if !apperror.ErrorIs(err, apperror.InvalidArguments) {
			t.Fatalf("invalid code")
		}
	})
}

func TestNewInvalidArguments(t *testing.T) {
	t.Run("should return app error with code Invalid Arguments", func(t *testing.T) {
		msg := "message"

		err := apperror.NewInvalidArguments(msg)

		if !apperror.ErrorIs(err, apperror.InvalidArguments) {
			t.Fatalf("invalid code")
		}
	})
}

func TestNewInternal(t *testing.T) {
	t.Run("should return app error with code Internal", func(t *testing.T) {
		msg := "message"

		err := apperror.NewInternal(msg)

		if !apperror.ErrorIs(err, apperror.Internal) {
			t.Fatalf("invalid code")
		}
	})
}

func TestNewOutOfRangePosition(t *testing.T) {
	t.Run("should return app error with code OutOfRangePosition", func(t *testing.T) {
		msg := "message"

		err := apperror.NewOutOfRangePosition(msg)

		if !apperror.ErrorIs(err, apperror.OutOfRangePosition) {
			t.Fatalf("invalid code")
		}
	})
}

func TestErrorIs(t *testing.T) {
	t.Run("should return false when given error not apperror", func(t *testing.T) {
		e := errors.New("message")

		res := apperror.ErrorIs(e, apperror.Internal)

		if res {
			t.Fatalf("method should return false")
		}
	})

	t.Run("should return true when given error and valid code", func(t *testing.T) {
		e, code := apperror.NewInvalidArguments("message"), apperror.InvalidArguments

		res := apperror.ErrorIs(e, code)

		if !res {
			t.Fatalf("method should return true")
		}
	})
}
