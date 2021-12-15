package cpal


func Clone[Slice ~[]T, T any](arr Slice) Slice {
	ret := make(Slice, len(arr))
	for i := range(ret) {
		ret[i] = arr[i]
	}
	return ret
}

func SliceEqual[Slice ~[]T, T comparable](a Slice, b Slice) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range(a) {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func AllOf[Slice ~[]T, T any](arr Slice, f func(T) bool) bool {
	for _, val := range(arr) {
		if !f(val) {
			return false
		}
	}
	return true
}

func AnyOf[Slice ~[]T, T any](arr Slice, f func(T) bool) bool {
	for _, val := range(arr) {
		if f(val) {
			return true
		}
	}
	return false
}

func NoneOf[Slice ~[]T, T any](arr Slice, f func(T) bool) bool {
	for _, val := range(arr) {
		if f(val) {
			return false
		}
	}
	return true
}

func AllOfZip[SliceA ~[]A, A any, SliceB ~[]B, B any](a SliceA, b SliceB, f func(A, B) bool) bool {
	n := Min(len(a), len(b))
	for i := 0; i < n; i++ {
		if !f(a[i], b[i]) {
			return false
		}
	}
	return true
}

func AnyOfZip[SliceA ~[]A, A any, SliceB ~[]B, B any](a SliceA, b SliceB, f func(A, B) bool) bool {
	n := Min(len(a), len(b))
	for i := 0; i < n; i++ {
		if f(a[i], b[i]) {
			return true
		}
	}
	return false
}

func NoneOfZip[SliceA ~[]A, A any, SliceB ~[]B, B any](a SliceA, b SliceB, f func(A, B) bool) bool {
	n := Min(len(a), len(b))
	for i := 0; i < n; i++ {
		if f(a[i], b[i]) {
			return false
		}
	}
	return true
}

