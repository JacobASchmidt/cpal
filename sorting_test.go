package cpal_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/JacobASchmidt/cpal"
)


func FuzzReverse(f *testing.F) {
	for _, arr := range([][]byte{{}, {0, 4, 2}, {1, 7, 3, 2, 5}}) {
		f.Add(arr)
	}

	f.Fuzz(func(t *testing.T, s []byte) {
		orig := cpal.Clone(s)
		cpal.Reverse(s)
		cpal.Reverse(s)

		if !cpal.SliceEqual(orig, s) {
			t.Fatalf("two reverses should return same array, orig=%v, s=%v", orig, s)
		}
	})
}

func FuzzPartition(f *testing.F) {
	for _, arr := range([][]byte{{1}, {0, 4, 2}, {1, 7, 3, 2, 5}}) {
		f.Add(arr, arr[len(arr)/2])
	}
	f.Fuzz(func (t *testing.T, s []byte, i byte) {
		fmt.Println(s, i)
		cmp := cpal.LessThan(i)
		point := cpal.Partition(s, cmp)
		if !cpal.IsPartitioned(s, cmp) {
			t.Fatalf("arr should be partitioned! partition by less than %v with s=%v and point=%v", i, s, point)
		}
		partitionPoint := cpal.PartitionPoint(s, cmp)
		if point != partitionPoint {
			t.Fatalf("partition and partitionPoint return different values! partition by less than %v with s=%v, point=%v, and partitionPoint=%v", i, s, point, partitionPoint)
		}
	})
}

func TestParition(t *testing.T) {
	t.Run("fuzzing", func(t *testing.T) {
		const n = 1_000
		const arr_len = 100
		const max_val = n
		r := rand.New(rand.NewSource(1))
	
		for len_ := 0; len_ < arr_len; len_++ {
			for i := 0; i < n; i++ {
				s := make([]int, len_)
				cpal.Fill(s, func() int {
					return r.Intn(max_val)
				})
				cmp := cpal.LessThan(i)
				point := cpal.Partition(s, cmp)
				if !cpal.IsPartitioned(s, cmp) {
					t.Fatalf("arr should be partitioned! partition by less than %v with s=%v and point=%v", i, s, point)
				}
				partitionPoint := cpal.PartitionPoint(s, cmp)
				if point != partitionPoint {
					t.Fatalf("partition and partitionPoint return different values! partition by less than %v with s=%v, point=%v, and partitionPoint=%v", i, s, point, partitionPoint)
				}
			}
		}
	})

	t.Run("same val array", func(t *testing.T) {
		const len_ = 1000
		const n = 42
		s := make([]int, len_)
		cpal.Fill(s, cpal.Value(n))

		i := cpal.Partition(s, cpal.EqualTo(n))
		if i != len_ {
			t.Fatalf("partition of all true should be len of slice, slice=%v, partition point=%v", s, i)
		}
		i = cpal.Partition(s, cpal.NotEqualTo(n))
		if i != 0 {
			t.Fatalf("partition of all false should be 0, slice=%v, partition point=%v", s, i)
		}
	})
}