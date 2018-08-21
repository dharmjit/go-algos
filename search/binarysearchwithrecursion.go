package search

func binarysearchwithrecursion(array []int, low int, high int, target int) (index int) {
	index = -1
	if high < low {
		return -1
	}
	mid := int((low + high) / 2)
	if target > array[mid] {
		index = binarysearchwithrecursion(array, mid+1, high, target)
	} else if target < array[mid] {
		index = binarysearchwithrecursion(array, low, mid, target)
	} else if target == array[mid] {
		index = mid
	} else {
		index = -1
	}
	return
}
