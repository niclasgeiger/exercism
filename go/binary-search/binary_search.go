package binarysearch

import "fmt"

const testVersion = 1

/*
Example:

SEARCH 3 in (1, 3, 5, 8, 13, 21, 34, 55, 89, 144)
	FIRST L=0, R=10 m=5
	[1, 3, 5, 8, 13, (21), 34, 55, 89, 144]
	SECOND L=0, R=5 m=2
	[ 1, 3, (5), 8, 13, 21 ], 34, 55, 89, 144
	THIRD L=0, R=2 m=1
	[ 1, (3), 5 ], 8, 13, 21, 34, 55, 89, 144
	FOUND 3 AT m=1
*/
func SearchInts(a []int, x int) int {
	L, R := 0, len(a)
	for L < R {
		m := L + (R-L)/2
		if a[m] < x {
			L = m + 1
		} else {
			R = m
		}
	}
	return L
}

//	MESSAGES
//   k found at beginning of slice
//   k found at end of slice
//   k found at index fx
//   k < all values
//   k > all n values
//   k > lv at lx, < gv at gx
//   slice has no values
func Message(slice []int, wanted int) string {
	if len(slice) == 0 {
		return "slice has no values"
	}
	index := SearchInts(slice, wanted)
	if index == len(slice) {
		return fmt.Sprintf("%d > all %d values", wanted, len(slice))
	}
	if slice[index] != wanted {
		if index == 0 {
			return fmt.Sprintf("%d < all values", wanted)
		}
		return fmt.Sprintf("%d > %d at index %d, < %d at index %d", wanted, slice[index-1], index-1, slice[index], index)
	}
	if index == 0 {
		return fmt.Sprintf("%d found at beginning of slice", wanted)
	}
	if index == len(slice)-1 {
		return fmt.Sprintf("%d found at end of slice", wanted)
	}
	return fmt.Sprintf("%d found at index %d", wanted, index)
}
