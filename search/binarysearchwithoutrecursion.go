package search

func binarysearchwithoutrecursion(array []int, v int) int {
	low := 0
	high := len(array) - 1
	var mid int
	for low <= high {
		mid = int((low + high) / 2)
		if array[mid] == v {
			return mid
		}
		if array[mid] < v {
			low = mid + 1
		}
		if array[mid] > v {
			high = mid - 1
		}
	}
	return -1
}
