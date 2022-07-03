package main

import "testing"

var conversionTests = []struct {
	input    int
	expected string
}{
	{1, "1"},
	{2, "2"},
	{3, "Fizz"},
	{4, "4"},
	{5, "Buzz"},
	{6, "Fizz"},
	{7, "7"},
	{8, "8"},
	{9, "Fizz"},
	{10, "Buzz"},
	{11, "11"},
	{12, "Fizz"},
	{13, "13"},
	{14, "14"},
	{15, "FizzBuzz"},
}

func TestCorrectOutputString(t *testing.T) {
	for _, test := range conversionTests {
		result := buildOutputString(test.input)
		if result != test.expected {
			t.Errorf("Wanted %s for %d, got %s", test.expected, test.input, result)
		}
	}
}
