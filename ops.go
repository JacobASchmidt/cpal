package cpal

import "constraints"


func Equal[T comparable](a T, b T) bool {
        return a == b
}

func NotEqual[T comparable](a T, b T) bool {
        return a != b
}
func Less[T constraints.Ordered](a T, b T) bool {
        return a < b
}
func LessEq[T constraints.Ordered](a T, b T) bool {
        return a <= b
}
func Greater[T constraints.Ordered](a T, b T) bool {
        return a > b
}
func GreaterEq[T constraints.Ordered](a T, b T) bool {
        return a >= b
}


func EqualTo[T comparable](a T) func(T) bool {
        return func(b T) bool {return a == b}
}

func NotEqualTo[T comparable](a T) func(T) bool  {
        return func(b T) bool {
		return a != b
	}
}
func LessThan[T constraints.Ordered](b T) func(T) bool  {
        return func(a T) bool {
		return a < b
	}
}
func LessEqTo[T constraints.Ordered](b T) func(T) bool  {
        return func(a T) bool {
		return a <= b
	}
}
func GreaterThan[T constraints.Ordered](b T) func(T) bool  {
        return func(a T) bool {
		return a > b
	}
}
func GreaterEqTo[T constraints.Ordered](b T) func(T) bool  {
        return func(a T) bool {
		return a >= b
	}
}
func Min[T constraints.Ordered](a T, b T) T {
	if a <= b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a T, b T) T {
	if a >= b {
		return a
	}
	return b
}


func ValueFunc[T any](t T) func() T {
	return func() T {
		return t
	}
}
