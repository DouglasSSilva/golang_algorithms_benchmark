package sort

// BubbleSort sort array with bubblesort algorithm
func BubbleSort(arr []int) []int {
	arrLen := len(arr)
	for i := range arr {
		for k := i; k < arrLen; k++ {
			if arr[i] > arr[k] {
				arr[i], arr[k] = arr[k], arr[i]
			}
		}
	}
	return arr
}
