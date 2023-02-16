package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func getPhoneNumbers(country string, phoneNumber string) string {
	// Build the API URL with the provided country name
	url := "jsonmock.hackerrank.com:80"
	request := fmt.Sprintf("GET /api/countries?name=%s HTTP/1.1\r\nHost: %s\r\n\r\n", country, url)

	// Establish a TCP connection to the API server
	conn, err := net.Dial("tcp", url)
	if err != nil {
		panic(err)
	}

	// Send the HTTP request
	_, err = fmt.Fprintf(conn, request)
	if err != nil {
		panic(err)
	}

	// Read the response body
	response, err := io.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	// Extract the phone numbers from the response
	phoneNumbers := []string{}
	responseString := string(response)
	start := strings.Index(responseString, "<phone>")
	for start != -1 {
		end := strings.Index(responseString[start:], "</phone>")
		phone := responseString[start+7 : start+end]
		phoneNumbers = append(phoneNumbers, phone)
		start = strings.Index(responseString[start+end:], "<phone>")
		if start != -1 {
			start += end
		}
	}

	// Return the phone numbers as a string
	return strings.Join(phoneNumbers, ",")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	country := readLine(reader)

	phoneNumber := readLine(reader)

	result := getPhoneNumbers(country, phoneNumber)

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
