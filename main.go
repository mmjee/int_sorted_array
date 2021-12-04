package int_sorted_array

import (
	"math"
	"sort"
)

type SortedIntegerArray struct {
	Items *[]int
}

func NewSortedIntegerArray(s []int) SortedIntegerArray {
	sort.Ints(s)
	return SortedIntegerArray{Items: &s}
}

func (sia SortedIntegerArray) Len() int {
	return len(*sia.Items)
}

func (sia *SortedIntegerArray) Insert(i int) {
	sl := *sia.Items
	ln := len(sl)

	lo := 0
	hi := len(sl) - 1

	// Perform a binary search to find the position to insert i
	for lo < hi {
		// ?!?!?!?!?!?!?!!?!
		// mid := (lo + hi) / 2
		mid := lo + ((hi - lo) / 2)
		if sl[mid] < i {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if ln == lo { // nil or empty slice or after last element
		sl = append(sl, i)
		sia.Items = &sl
		return
	}

	sl = append(sl[:lo+1], sl[lo:]...) // index < len(a)
	sl[lo] = i
	sia.Items = &sl
	return
}

func (sia *SortedIntegerArray) Delete(idx int) {
	sl := *sia.Items
	newSlice := append(sl[:idx], sl[idx+1:]...)
	sia.Items = &newSlice
}

func (sia SortedIntegerArray) averageWithoutOverflow(a int, b int) int {
	return (a / 2) + (b / 2) + ((a%2 + b%2) / 2)
}

func (sia SortedIntegerArray) Median() int {
	sl := *sia.Items
	length := len(sl)
	avgNeeded := length%2 != 1
	if avgNeeded {
		mid := float64(length / 2)
		hiMid := int(math.Ceil(mid))
		lowMid := int(math.Floor(mid))

		hiVal := sl[hiMid]
		lowVal := sl[lowMid]

		return sia.averageWithoutOverflow(hiVal, lowVal)
	} else {
		return sl[length/2]
	}
}
