package cpal_test

import (
	"fmt"
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