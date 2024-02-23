package main

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + (high-low)/2
		guess := list[mid]
		if guess == item {
			return mid
		} else if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func bubleShort(list []string) {
	n := len(list) - 1
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
func binarysearchString(list []string, item string) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := low + (high-low)/2
		guess := list[mid]
		if guess == item {
			return mid
		} else if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func binarysearchString2(list []string, item string) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := low + (high-low)/2
		guess := list[mid]
		if guess == item {
			return mid
		} else if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
func main() {
	list := []int{1, 3, 5, 7, 9}

	result := binarySearch(list, 7)
	if result == -1 {
		println("none")
	} else {
		println(result)
	}
	//nameList := []string{"sakura", "sasuke", "naruto", "hinata", "sai"}
	//bubleShort(nameList)
	//resultStr := binarysearchString2(nameList, "sasuke")
	//if resultStr == -1 {
	//	println("none")
	//} else {
	//	println(resultStr)
	//}
}
