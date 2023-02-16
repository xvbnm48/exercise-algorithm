package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'calculateScore' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING text
 *  2. STRING prefixString
 *  3. STRING suffixString
 */

func calculateScore(text string, prefixString string, suffixString string) string {
	// Write your code here
	var prefix, suffix string
	var prefixScore, suffixScore int
	var prefixIndex, suffixIndex int
	var prefixFound, suffixFound bool

	for i := 0; i < len(text); i++ {
		if text[i] == prefixString[0] {
			prefixIndex = i
			prefixFound = true
			break
		}
	}

	if prefixFound {
		prefix = text[prefixIndex : prefixIndex+len(prefixString)]
	}

	for i := len(text) - 1; i >= 0; i-- {
		if text[i] == suffixString[0] {
			suffixIndex = i
			suffixFound = true
			break
		}
	}

	if suffixFound {
		suffix = text[suffixIndex : suffixIndex+len(suffixString)]
	}

	if prefix == prefixString {
		prefixScore = 1
	}

	if suffix == suffixString {
		suffixScore = 1
	}

	if prefixScore == 1 && suffixScore == 1 {
		return "BOTH"
	}

	if prefixScore == 1 {
		return "PREFIX"
	}

	if suffixScore == 1 {
		return "SUFFIX"
	}

	return "nothing"

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	text := readLine(reader)

	prefixString := readLine(reader)

	suffixString := readLine(reader)

	result := calculateScore(text, prefixString, suffixString)

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
