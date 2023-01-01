package main

import "fmt"

func main() {
	user0 := []string{"/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"}
	user1 := []string{"/start", "/pink", "/register", "/orange", "/red", "a"}
	user2 := []string{"a", "/one", "/two"}
	user3 := []string{"/pink", "/orange", "/yellow", "/plum", "/blue", "/tan", "/red", "/amber", "/CornflowerBlue", "/LightGoldenRodYellow"}
	user4 := []string{"/pink", "/orange", "/amber", "/plum", "/blue", "/tan", "/red", "/lavender", "/CornflowerBlue", "/LightGoldenRodYellow"}
	user5 := []string{"a"}

	fmt.Println(findLongestHistory(user0, user1)) // ["/pink", "/register", "/orange"]
	fmt.Println(findLongestHistory(user0, user2)) // [] (empty)
	fmt.Println(findLongestHistory(user0, user0)) // ["/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"]
	fmt.Println(findLongestHistory(user5, user2)) // ["a"]
	fmt.Println(findLongestHistory(user3, user4)) // ["/plum", "/blue", "/tan", "/red"]
	fmt.Println(findLongestHistory(user4, user3)) // ["/plum", "/blue", "/tan",

}

func findLongestHistory(user1, user2 []string) []string {
	// result will hold the longest contiguous sequence of URLs that appears in both lists
	var result []string

	// Iterate through both lists of URLs
	for i := 0; i < len(user1); i++ {
		for j := 0; j < len(user2); j++ {
			// If the URLs at index i in user1 and index j in user2 match,
			// check if the subsequent URLs in both lists also match
			if user1[i] == user2[j] {
				// temp will hold the current contiguous sequence of URLs
				var temp []string
				temp = append(temp, user1[i])

				// Check subsequent URLs in both lists
				k := i + 1
				l := j + 1
				for k < len(user1) && l < len(user2) && user1[k] == user2[l] {
					temp = append(temp, user1[k])
					k++
					l++
				}

				// If temp is longer than result, update result
				if len(temp) > len(result) {
					result = temp
				}
			}
		}
	}

	return result
}
