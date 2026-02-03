package gomonad

// Result represents a value that is either a success (Ok) or a failure (error).
type Result[T any] struct {
	value   T
	err     error
	isError bool
}

// Ok wraps a successful value.
func Ok[T any](val T) Result[T] {
	return Result[T]{value: val, isError: false}
}

// Err wraps a standard Go error.
func Err[T any](err error) Result[T] {
	return Result[T]{err: err, isError: true}
}

// --- Methods ---

func (r Result[T]) IsOk() bool    { return !r.isError }
func (r Result[T]) IsErr() bool   { return r.isError }
func (r Result[T]) Get() T        { return r.value }
func (r Result[T]) GetErr() error { return r.err }

// --- Chaining Functions ---

// BindResult chains operations that return another Result.
func BindResult[T any, U any](r Result[T], f func(T) Result[U]) Result[U] {
	if r.isError {
		return Err[U](r.err)
	}
	return f(r.value)
}

// MapResult transforms the successful value if it exists.
func MapResult[T any, U any](r Result[T], f func(T) U) Result[U] {
	if r.isError {
		return Err[U](r.err)
	}
	return Ok[U](f(r.value))
}
