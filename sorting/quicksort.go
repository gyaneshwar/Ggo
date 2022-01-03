package sorting

//quick sort
// if len less than 2 return array
// set pivote and left to zero, right to arr.len - 1
// swap pivot to the right corner
// pile all the element left to the pivot
// place pivot after the last smaller element, i.e., the left index
// go down the rabit hole
// sort left array and right array
// return sorted array
//https://stackoverflow.com/questions/15802890/idiomatic-quicksort-in-go/15803401#15803401
func QuickSort(arr []int32) {
	// if len is less than 2, return the array
	if len(arr) < 2 {
		return
	}

	// identify left and right index
	pivot, left, right := 0, 0, len(arr)-1

	// assign pivot to the right most
	arr[pivot], arr[right] = arr[right], arr[pivot]

	// pile all the elements left to pivot
	for i := range arr {
		if arr[i] > arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	// place pivot element after the left smallest value (think when there are only 2 times in the array)
	arr[left], arr[right] = arr[right], arr[left]

	//go down the rabit hole
	// sort left array
	QuickSort(arr[:left])
	// sort right array
	QuickSort(arr[left+1:])

	//return sorted array
}
