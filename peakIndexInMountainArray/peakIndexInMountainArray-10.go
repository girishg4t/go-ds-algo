package peakindexinmountainarray

// [24,69,100,99,79,78,67,36,26,19]
func PeakIndexInMountainArray(arr []int) int {
	low := 0
	hig := len(arr) - 1
	for low <= hig {
		mid := low + (hig-low)/2
		if mid == 0 {
			mid = mid + 1
		}
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid] < arr[mid-1] {
			hig = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
