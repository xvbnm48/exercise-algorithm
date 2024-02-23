package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumGroups' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY predators as parameter.
 */

func minimumGroups(predators []int32) int32 {
	// Write your code here
	depthSpecies := func(i int32) int32 {
		depth := int32(0)
		for predators[i] != -1 {
			i = predators[i]
			depth++
		}
		return depth + 1
	}
	maxDepth := int32(0)
	for i, _ := range predators {
		depth := depthSpecies(int32(i))
		if depth > maxDepth {
			maxDepth = depth
		}
	}

	return maxDepth

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	predatorsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var predators []int32

	for i := 0; i < int(predatorsCount); i++ {
		predatorsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		predatorsItem := int32(predatorsItemTemp)
		predators = append(predators, predatorsItem)
	}

	result := minimumGroups(predators)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
