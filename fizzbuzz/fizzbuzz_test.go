package main

import "testing"

func TestCorrectOutputString(t *testing.T) {
	result := buildOutputString(15)
	if result != "FizzBuzz" {
		t.Errorf("Wanted FizzBuzz, got %s", result)
	}
}
