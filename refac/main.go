package main

import "strings"

func main() {
	s := "I, want, to, be, softwarem engineer"

	strarr := strings.Split(s, ",")
	result := findString(strarr)
	println(result)
}

func findString(strarr []string) string {
	i := strarr[0]
	t := strarr[1]

	if len(t) == 1 {
		return i
	}

	return findString(t)
}
