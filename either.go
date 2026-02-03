package gomonad

// Either represents a value of one of two types.
// A is usually the "Left" type, B is usually the "Right" type.
type Either[A any, B any] struct {
	left   A
	right  B
	isLeft bool
}

// Left creates an Either containing the left value.
func Left[A any, B any](val A) Either[A, B] {
	return Either[A, B]{left: val, isLeft: true}
}

// Right creates an Either containing the right value.
func Right[A any, B any](val B) Either[A, B] {
	return Either[A, B]{right: val, isLeft: false}
}

// --- Methods ---

func (e Either[A, B]) IsLeft() bool  { return e.isLeft }
func (e Either[A, B]) IsRight() bool { return !e.isLeft }

func (e Either[A, B]) Left() A  { return e.left }
func (e Either[A, B]) Right() B { return e.right }

// --- Transformation Functions ---

// MapRight transforms the Right value, leaving Left untouched.
// This is the standard "monadic map" for Either.
func MapRight[A any, B any, C any](e Either[A, B], f func(B) C) Either[A, C] {
	if e.isLeft {
		return Left[A, C](e.left)
	}
	return Right[A, C](f(e.right))
}

// Fold takes two functions and executes the one corresponding to the side present.
// This is the most common way to "exit" an Either.
func Fold[A any, B any, T any](e Either[A, B], leftFunc func(A) T, rightFunc func(B) T) T {
	if e.isLeft {
		return leftFunc(e.left)
	}
	return rightFunc(e.right)
}

// Swap exchanges the Left and Right values.
// If it was Left(a), it becomes Right(a). If it was Right(b), it becomes Left(b).
func Swap[A any, B any](e Either[A, B]) Either[B, A] {
	if e.isLeft {
		return Right[B, A](e.left)
	}
	return Left[B, A](e.right)
}

// MapLeft transforms the Left value, leaving Right untouched.
func MapLeft[A any, B any, C any](e Either[A, B], f func(A) C) Either[C, B] {
	if e.isLeft {
		return Left[C, B](f(e.left))
	}
	return Right[C, B](e.right)
}

// ToResult converts an Either to a Result.
// This assumes the Left side can be represented as a Go error.
func ToResult[T any](e Either[error, T]) Result[T] {
	if e.isLeft {
		return Err[T](e.left)
	}
	return Ok[T](e.right)
}
