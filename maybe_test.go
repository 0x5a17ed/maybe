package maybe_test

import (
	"errors"
	"testing"

	"github.com/0x5a17ed/maybe"
)

func TestMust_SuccessError(t *testing.T) {
	v := maybe.Must[int](func() (int, error) {
		return 99, nil
	})

	if v != 99 {
		t.Errorf("Must() = %v, want 99", v)
	}
}

func TestMust_PanicError(t *testing.T) {
	ErrExpected := errors.New("oops")

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic, got none")
		} else if r != ErrExpected {
			t.Fatalf("expected panic with 'oops', got %v", r)
		}
	}()

	maybe.Must[string](func() (string, error) {
		return "", ErrExpected
	})
}

func TestMust_SuccessBool(t *testing.T) {
	fn := maybe.MustFn[int](func() (int, bool) {
		return 99, true
	})

	if v := fn(); v != 99 {
		t.Errorf("Must() = %v, want 99", v)
	}
}

func TestMust_PanicBool(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic, got none")
		} else if r != maybe.ErrFailed {
			t.Fatalf("expected panic with ErrFailed, got %v", r)
		}
	}()

	maybe.Must[string](func() (string, bool) {
		return "", false
	})
}
