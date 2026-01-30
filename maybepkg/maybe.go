package maybepkg

// Maybe represents a value that might or might not exist.
type Maybe[T any] struct {
	value  T
	isJust bool
}

// Just wraps a concrete value.
func Just[T any](val T) Maybe[T] {
	return Maybe[T]{value: val, isJust: true}
}

// Nothing represents the absence of a value.
func Nothing[T any]() Maybe[T] {
	return Maybe[T]{isJust: false}
}

// Bind (or FlatMap) allows chaining operations that also return a Maybe.
func Bind[T any, U any](m Maybe[T], f func(T) Maybe[U]) Maybe[U] {
	if !m.isJust {
		return Nothing[U]()
	}
	return f(m.value)
}

// Map transforms the value inside if it exists.
func Map[T any, U any](m Maybe[T], f func(T) U) Maybe[U] {
	if !m.isJust {
		return Nothing[U]()
	}
	return Just(f(m.value))
}

// Unpack returns the value or a provided default.
func (m Maybe[T]) Unpack(defaultValue T) T {
	if !m.isJust {
		return defaultValue
	}
	return m.value
}

// IsJust returns true if the Maybe contains a value.
func (m Maybe[T]) IsJust() bool {
	return m.isJust
}

// IsNothing returns true if the Maybe is empty.
func (m Maybe[T]) IsNothing() bool {
	return !m.isJust
}

// Get returns the underlying value.
// Note: Use this only after checking IsJust() to avoid unexpected zero-values.
func (m Maybe[T]) Get() T {
	return m.value
}
