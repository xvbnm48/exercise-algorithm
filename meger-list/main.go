package main

// mergeSortedLists merges two sorted lists of order numbers into a single,
import "fmt" // sorted list.
func mergeSortedLists(list1 []int, list2 []int) []int {
	// Create a new slice to hold the merged list.
	mergedList := make([]int, 0)

	// Keep track of the current index in each list.
	index1, index2 := 0, 0

	// Iterate over both lists until one of them is fully traversed.
	for index1 < len(list1) && index2 < len(list2) {
		// If the current element in the first list is smaller than the current
		// element in the second list, append it to the merged list and increment
		// the index in the first list. Otherwise, do the same for the second list.
		if list1[index1] < list2[index2] {
			mergedList = append(mergedList, list1[index1])
			index1++
		} else {
			mergedList = append(mergedList, list2[index2])
			index2++
		}
	}

	// Append the remaining elements from the non-fully traversed list to the
	// merged list.
	mergedList = append(mergedList, list1[index1:]...)
	mergedList = append(mergedList, list2[index2:]...)

	return mergedList
}

func main() {
	list1 := []int{3, 4, 6, 10, 11, 15}
	list2 := []int{1, 5, 8, 12, 14, 19}
	fmt.Println(mergeSortedLists(list1, list2))
}
