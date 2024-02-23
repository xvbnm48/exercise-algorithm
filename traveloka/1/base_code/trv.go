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
 * Complete the 'slowestKey' function below.
 *
 * The function is expected to return a CHARACTER.
 * The function accepts 2D_INTEGER_ARRAY keyTimes as parameter.
 */

func slowestKey(keyTimes [][]int32) string {
	// Write your code here
	if len(keyTimes) == 0 {
		return ""
	}
	maximumDuration := keyTimes[0][1]
	slowestKey := keyTimes[0][0]

	for i := 1; i < len(keyTimes); i++ {
		duration := keyTimes[i][1] - keyTimes[i-1][1]
		if duration > maximumDuration {
			maximumDuration = duration
			slowestKey = keyTimes[i][0]
		}
	}
	return string('a' + slowestKey)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	keyTimesRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	keyTimesColumns, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var keyTimes [][]int32
	for i := 0; i < int(keyTimesRows); i++ {
		keyTimesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var keyTimesRow []int32
		for _, keyTimesRowItem := range keyTimesRowTemp {
			keyTimesItemTemp, err := strconv.ParseInt(keyTimesRowItem, 10, 64)
			checkError(err)
			keyTimesItem := int32(keyTimesItemTemp)
			keyTimesRow = append(keyTimesRow, keyTimesItem)
		}

		if len(keyTimesRow) != int(keyTimesColumns) {
			panic("Bad input")
		}

		keyTimes = append(keyTimes, keyTimesRow)
	}

	result := slowestKey(keyTimes)

	fmt.Fprintf(writer, "%s\n", result)

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
