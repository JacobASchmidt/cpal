package cpal


func Swap[T any](a *T, b *T) {
        temp := *a
        *a = *b
        *b = temp
}

func Reverse[Slice ~[]T, T any](arr Slice) {
        l, r := 0, len(arr) - 1
        for l < r {
                Swap(&arr[l], &arr[r])
                l++
                r--
        }
}

func Rotate[Slice ~[]T, T any](arr Slice, idx int) int {
        s := []T(arr)
        Reverse(s[:idx])
        Reverse(s[idx:])
        Reverse(s)
        return len(s) - idx
}

func Partition[Slice ~[]T, T any](arr Slice, f func(T) bool) int {
        i := 0
        for j, val := range(arr) {
                if f(val) {
                        Swap(&arr[i], &arr[j])
                        i++
                }
        }
        return i
}

func StablePartition[Slice ~[]T, T any](arr Slice, f func(T) bool) int {
        s := []T(arr)
        if len(s) <= 1 {
                if len(s) == 0 {
                        return 0
                }
                if f(s[0]) {
                        return 1
                } else {
                        return 0
                }
        }

        mid := len(s) / 2
        left := StablePartition(s[:mid], f)
        right := StablePartition(s[mid:], f)

        return left + Rotate(s[left:right], mid - left)
}


//cmp must NOT be reflexive!
func Sort[Slice ~[]T, T any](arr Slice, cmp func(T, T) bool) {
       s := []T(arr)

       if len(s) <= 1 {
               return
       }

       mid := len(s) / 2
       last := len(s) - 1

       Swap(&s[mid], &s[last])

       point := Partition(s[:last], func(el T) bool {
               return cmp(s[last], el)
       })

       Swap(&s[point], &s[last])

       Sort(s[:point], cmp)
       Sort(s[point+1:], cmp)
}
//cmp must not be reflexive! note also already sorted array will have quadratic time complexity
func StableSort[Slice ~[]T, T any](arr Slice, cmp func(T, T) bool) {
        s := []T(arr)
        if len(s) <= 1 {
                return
        }
        
        point := StablePartition(s, func(el T) bool {
                return cmp(el, s[0])
        })
        if point == 0 {
                point++
        }

        StableSort(s[:point], cmp)
        StableSort(s[point:], cmp)        
}

func PlaceNthElement[Slice ~[]T, T any](arr Slice, i int, cmp func(T, T) bool) {
        
       s := []T(arr)

       if len(s) <= 1 {
               return
       }

       mid := len(s) / 2
       last := len(s) - 1

       Swap(&s[mid], &s[last])

       point := Partition(s[:last], func(el T) bool {
               return cmp(s[last], el)
       })

       Swap(&s[point], &s[last])
       if i == point {
               return
       } else if i < point {
               PlaceNthElement(s[:point], i, cmp)
       } else {
                PlaceNthElement(s[point+1:], i - (point + 1), cmp)
       }
}