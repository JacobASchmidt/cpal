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

func Reduce[Slice ~[]T, T any, R any](s Slice, init R, f func(R, T) R) R {
	for _, val := range(s) {
		init = f(init, val)
	}
	return init
}

func ZipReduce[SliceA ~[]A, A any, SliceB ~[]B, B any, R any](a SliceA, b SliceB, init R, f func(R, A, B) R) R {
	n := Min(len(a), len(b))
	for i := 0; i < n; i++ {
		init = f(init, a[i], b[i])
	}
	return init
}


func MapReduce[Slice ~[]T, T any, R any](s Slice, init R, f func(R, T) R) R {
	for _, val := range(s) {
		init = f(init, val)
	}
	return init
}

func ZipMapReduce[SliceA ~[]A, A any, SliceB ~[]B, B any, R any, M any](a SliceA, b SliceB, map_ func(A, B) M, init R, reduce_ func(R, M) R) R {
	n := Min(len(a), len(b))
	for i := 0; i < n; i++ {
		init = reduce_(init, map_(a[i], b[i]))
	}
	return init
}


func MapZipReduce[SliceA ~[]A, A any, SliceB ~[]B, B any, R any, AM any, BM any](a SliceA, b SliceB, mapA func(A) AM, mapB func(B) BM, init R, reduce_ func(R, AM, BM) R) R {
	n := Min(len(a), len(b))
	for i := 0; i < n; i++ {
		init = reduce_(init, mapA(a[i]), mapB(b[i]))
	}
	return init
}

func FindIf[Slice ~[]T, T any](arr Slice, f func(T) bool) int {
	for i, val := range(arr) {
		if f(val) {
			return i
		}
	}
	return -1
}


func Find[Slice ~[]T, T comparable](arr Slice, val T) int {
	for i, val2 := range(arr) {
		if val == val2 {
			return i
		}
	}
	return -1
}

func AdjecentFindIf[Slice ~[]T, T any](arr Slice, f func(T, T) bool) int {
	for i := 1; i < len(arr); i++ {
		if f(arr[i-1], arr[i]) {
			return i-1
		}
	}
	return -1
}

func MapCopy[InSlice ~[]T, OutSlice ~[]R, T any, R any](in InSlice, out OutSlice, f func(T) R) OutSlice {
	if out == nil || len(out) < len(in) {
		out = make([]R, len(in))
	} 
	for i, val := range(in) {
		out[i] = f(val)
	}
	return out	
}