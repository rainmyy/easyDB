package file

import (
	"sort"
)

func InIntSliceSortedFunc(stack []int) func(int) bool {
	sort.Ints(stack)
	return func(needle int) bool {
		index := sort.SearchInts(stack, needle)
		return index < len(stack) && stack[index] == needle
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
