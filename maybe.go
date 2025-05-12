package maybe

import "errors"

var (
	// ErrFailed is returned when a function returning a boolean value wrapped in Maybe fails.
	ErrFailed = errors.New("maybe: operation failed")
)

// TryFn wraps a function returning either (T, error) or (T, bool)
// into a standard interface returning (T, error).
type TryFn[T any] interface {
	Try() (T, error)
}

// Result is a constraint for a function returning either (T, error) or (T, bool).
type Result[T any] interface {
	~func() (T, error) | ~func() (T, bool)
}

type tryFromError[T any] struct{ fn func() (T, error) }

func (t tryFromError[T]) Try() (T, error) {
	return t.fn()
}

type tryFromBool[T any] struct{ fn func() (T, bool) }

func (t tryFromBool[T]) Try() (T, error) {
	v, ok := t.fn()
	if !ok {
		return v, ErrFailed
	}
	return v, nil
}

// Wrap turns a function of (T, error) or (T, bool) into a TryFn[T] interface returning (T, error).
func Wrap[T any, F Result[T]](fn F) TryFn[T] {
	switch any(fn).(type) {
	case func() (T, error):
		return tryFromError[T]{fn: any(fn).(func() (T, error))}
	case func() (T, bool):
		return tryFromBool[T]{fn: any(fn).(func() (T, bool))}
	default:
		panic("maybe.Wrap: unsupported function type")
	}
}

// WrapFn turns a function of (T, error) or (T, bool) into a TryFn[T] interface returning (T, error).
func WrapFn[T any, F Result[T]](fn F) func() (T, error) {
	return Wrap[T](fn).Try
}

// Must panics if the wrapped call fails, using a default message.
func Must[T any, F Result[T]](fn F) T {
	v, err := WrapFn[T](fn)()
	if err != nil {
		panic(err)
	}
	return v
}

// MustFn is a convenience function that wraps a function returning either (T, error) or (T, bool)
// into a standard interface returning T and panics if the call fails.
func MustFn[T any, F Result[T]](fn F) func() T {
	return func() T {
		return Must[T](fn)
	}
}
