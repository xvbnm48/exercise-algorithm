package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	var cases = []struct {
		input  string
		output string
	}{
		{"1\n100\n", "Anda membeli product1\nJumlah uang yang dibayarkan: 100\nKembalian: 50\n"},
		{"2\n200\n", "Anda membeli product2\nJumlah uang yang dibayarkan: 200\nKembalian: 100\n"},
		// Add more test cases here
	}

	for _, test := range cases {
		in := bytes.Buffer{}
		out := bytes.Buffer{}

		in.Write([]byte(test.input))

		// Temporarily replace the standard input/output with our buffers
		stdin := os.Stdin
		stdout := os.Stdout

		defer func() {
			os.Stdin = stdin
			os.Stdout = stdout
		}()

		os.Stdin = &in
		os.Stdout = &out

		main()

		result := out.String()
		if !strings.Contains(result, test.output) {
			t.Errorf("Expected '%s', got '%s'", test.output, result)
		}
	}
}
