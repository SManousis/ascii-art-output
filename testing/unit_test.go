package main

import (
	"bytes"
	"os"
	"testing"
)

// Helper function to capture the output of the main function
func captureOutput(f func()) string {
	old := os.Stdout

	// Create a pipe to capture the output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function f that writes to stdout
	f()

	// Close the writer and restore stdout
	w.Close()
	os.Stdout = old

	// Read the captured output from the reader
	var buf bytes.Buffer
	buf.ReadFrom(r)

	return buf.String()
}

func TestMainWithVariousStrings(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "Test with word 'HI'",
			args: []string{"cmd", "HI", "standard"},
		},
		{
			name: "Test with invalid argument",
			args: []string{"cmd", "HI", "mistake"},
		},
		{
			name: "Test with invalid number of arguments",
			args: []string{"cmd", "HI", "standard", "abc"},
		},
		{
			name: "Test with word 'HELLO'",
			args: []string{"cmd", "HELLO", "standard"},
		},
		{
			name: "Test with word 'A B'",
			args: []string{"cmd", "A B", "standard"},
		},
		{
			name: "Test with word 'HELLO' in shadow",
			args: []string{"cmd", "HELLO", "shadow"},
		},
		{
			name: "Test with word 'HELLO' in thinkertoy",
			args: []string{"cmd", "HELLO", "thinkertoy"},
		},
		{
			name: "Test with empty string",
			args: []string{"cmd", "", "standard"},
		},
		{
			name: "Test with newline in string",
			args: []string{"cmd", "HI\\nHELLO", "standard"},
		},
		{
			name: "Test with characters in string",
			args: []string{"cmd", "Hello\"skoul(i)kom@p$", "standard"},
		},
		{
			name: "Test with output file option",
			args: []string{"cmd", "--output=exampletest.txt", "Hello\"skoul(i)kom@p$", "standard"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Temporarily set os.Args to the test case arguments
			os.Args = tt.args

			// Capture the output of main()
			output := captureOutput(main)

			// Print the output to see the result
			t.Logf("Output for %s:\n%s", tt.name, output)
		})
	}
}
