package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Determine the length of the two arrays
	m, n := len(nums1), len(nums2)

	// Make sure that nums1 is always the shorter array
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}

	// Start the binary search at the median of the shorter array
	low, high := 0, m
	halfLen := (m + n + 1) / 2

	// Perform the binary search until the desired value is found
	for low <= high {
		// Calculate the midpoint of the current search range
		mid := (low + high) / 2

		// Find the corresponding midpoint in the longer array
		j := halfLen - mid
		if mid < m && nums2[j-1] > nums1[mid] {
			// If the midpoint of the shorter array is too small, increase it
			low = mid + 1
		} else if mid > 0 && nums1[mid-1] > nums2[j] {
			// If the midpoint of the shorter array is too large, decrease it
			high = mid - 1
		} else {
			// If the midpoint of the shorter array is just right, calculate the median
			var maxLeft float64
			if mid == 0 {
				// If the midpoint of the shorter array is at the beginning, the max left value is the max of the leftmost elements of the two arrays
				maxLeft = float64(max(nums1[0], nums2[0]))
			} else if j == 0 {
				// If the midpoint of the longer array is at the beginning, the max left value is the max of the leftmost elements of the two arrays
				maxLeft = float64(max(nums1[mid-1], nums2[j-1]))
			} else {
				// If neither array has reached the beginning, the max left value is the maximum of the two elements to the left of the midpoints
				maxLeft = float64(max(nums1[mid-1], nums2[j-1]))
			}

			// If the length of the combined array is odd, the median is the max left value
			if (m+n)%2 == 1 {
				return maxLeft
			}

			// If the length of the combined array is even, the median is the average of the max left value and the min right value
			var minRight float64
			if mid == m {
				// If the midpoint of the shorter array is at the end, the min right value is the min of the rightmost elements of the two arrays
				minRight = float64(min(nums1[mid-1], nums2[j]))
			} else if j == n {
				// If the midpoint of the longer array is at the end, the min right value is the min of the rightmost elements of the two arrays
				minRight = float64(min(nums1[mid], nums2[j-1]))
			} else {
				// If neither array has reached the end, the min right value is the minimum of the two elements to the right of the midpoints
				minRight = float64(min(nums1[mid], nums2[j]))
			}

			// Return the average of the max left value and the min right value
			return (maxLeft + minRight) / 2
		}
	}

	// If the desired value is not found, return 0
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	println(findMedianSortedArrays(nums1, nums2)) // Output: 2
}
