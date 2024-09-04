package main

import (
	"fmt"
	"testing"
)

func TestGenMD5(t *testing.T) {
	input := "will"
	output := genMD5(input)
	expected := 40075

	if output != expected {
		fmt.Printf("Wrong output:\n output: %d\n expected: %d\n", output, expected)
		fmt.Println()
		return 
	}
}
